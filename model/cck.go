package model

import (
	"github.com/fale/drupal2hugo/util"
	"fmt"
	php "github.com/wulijun/go-php-serialize/phpserialize"
)

type CCKFieldType struct {
	Name           string                      `db:"field_name"`
	Type           string                      `db:"type"`
	ContentType    string                      `db:"type_name"`
	Gs             string                      `db:"global_settings"`
	GlobalSettings map[interface{}]interface{} `db:"-"`
	Required       bool                        `db:"required"`
	Multiple       bool                        `db:"multiple"`
	DBStorage      bool                        `db:"db_storage"`
	Module         string                      `db:"module"`
	Dbc            string                      `db:"db_columns"`
	DBColumns      map[interface{}]interface{} `db:"-"`
	Locked         bool                        `db:"locked"`
}

func (db Database) CCKFields() (map[string][]*CCKFieldType, error) {
	var err error = nil
	_sql := `select cnf.field_name, type, global_settings, required, multiple, 
	  db_storage, module, db_columns, locked, cnfi.type_name as type_name
	  from %scontent_node_field cnf, %scontent_node_field_instance cnfi 
	  where cnf.field_name = cnfi.field_name and cnf.active = 1 and cnfi.widget_active = 1`
	query := fmt.Sprintf(_sql, db.Prefix, db.Prefix)
	var fields []*CCKFieldType
	_, err = db.DbMap.Select(&fields, query)
	if err != nil {
		util.LogError(err, "Error executing CCKFields select.")
		return nil, err
	}
	fieldsByType := make(map[string][]*CCKFieldType)
	for _, cft := range fields {
		decoded, err := php.Decode(cft.Gs)
		cft.GlobalSettings = decoded.(map[interface{}]interface{})
		if err != nil {
			util.LogError(err, "Unable to parse global settings for "+cft.Name)
		}
		decoded, err = php.Decode(cft.Dbc)
		cft.DBColumns = decoded.(map[interface{}]interface{})
		if err != nil {
			util.LogError(err, "Unable to parse db columns for "+cft.Name)
		}
		typeFields, ok := fieldsByType[cft.ContentType]
		if !ok {
			typeFields = make([]*CCKFieldType, 0, 100)
		}
		fieldsByType[cft.ContentType] = append(typeFields, cft)
		err = nil
	}
	return fieldsByType, err
}

type CCKField struct {
	Name, SubName, Type string
}

func (db Database) CCKDataForNode(node *JoinedNodeDataBody, fields []*CCKFieldType) (map[CCKField]interface{}, error) {
	results := make(map[CCKField]interface{})
	var err error
	selectFields := make([]CCKField, 0, 100)
	_sql := "SELECT "
	first := true
	for _, cft := range fields {
		fmt.Printf("looping through dbcolumns for %s\n", cft.Name)
		for k, v := range cft.DBColumns {
			fmt.Printf("%#v = %#v\n", k, v)
			key := k.(string)
			dbColumn := v.(map[interface{}]interface{})
			typ, ok := dbColumn["type"]
			if ok {
				cckField := CCKField{cft.Name, key, typ.(string)}
				if !first {
					_sql = _sql + ", "
				}
				_sql = _sql + cft.Name + "_" + key
				selectFields = append(selectFields, cckField)
				first = false
			}
		}
	}
	if first {
		return results, nil
	}
	_sql += " FROM %scontent_type_%s WHERE vid = %d and nid = %d"
	query := fmt.Sprintf(_sql, db.Prefix, node.Type, node.Vid, node.Nid)
	rows, err := db.Db.Query(query)
	util.CheckErrFatal(err, "Selecting CCK field values for node.", node.Nid)
	defer rows.Close()
	columnNames, err := rows.Columns()
	if err != nil {
		util.Fatal("Error getting column names: %v", err) // or whatever error handling is appropriate
	}
	if len(columnNames) != len(selectFields) {
		util.Fatal("column names length and select fields length do not match!")
	}
	columns := make([]interface{}, len(columnNames))
	columnPointers := make([]interface{}, len(columnNames))
	for i := 0; i < len(columnNames); i++ {
		columnPointers[i] = &columns[i]
	}
	if rows.Next() {
		if err := rows.Scan(columnPointers...); err != nil {
			util.Fatal("Error reading a row: %v", err)
		}
		for i, colName := range selectFields {
			switch colName.Type {
			case "text", "varchar":
				switch cv := columns[i].(type) {
				case []byte:
					results[colName] = string(cv)
				default:
					if cv != nil {
						results[colName] = cv.(string)
					}
				}
			default:
				results[colName] = columns[i]
			}
		}
	}
	return results, nil
}
