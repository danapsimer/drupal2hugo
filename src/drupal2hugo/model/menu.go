package model

import (
	"drupal2hugo/util"
	"github.com/go-sql-driver/mysql"
)

const (
	NoSuchTable uint16 = 1146 // http://dev.mysql.com/doc/refman/5.5/en/error-messages-server.html
)

type Book struct {
	Mlid     int32
	Nid      int32
	Bid      int32
}

type MenuCustom struct {
	MenuName    string  `db:"menu_name"`
	Title       string
	Description string
}

func (db Database) AllBooks() []*Book {
	sql := "select * from " + db.Prefix + "xx"
	list, err := db.DbMap.Select(Book{}, sql)
	if err != nil {
		e := err.(*mysql.MySQLError)
		if e.Number == NoSuchTable {
			return []*Book{}
		}
		util.CheckErrFatal(err, sql)
	}
	return copyOutBook(list)
}

func copyOutBook(rows []interface{}) []*Book {
	size := len(rows)
	result := make([]*Book, size)
	for i := 0; i < size; i++ {
		result[i] = rows[i].(*Book)
	}
	return result
}


func (db Database) AllMenus() []*MenuCustom {
	sql := "select * from " + db.Prefix + "menu_custom"
	list, err := db.DbMap.Select(MenuCustom{}, sql)
	util.CheckErrFatal(err, sql)
	return copyOutMenuCustoms(list)
}

func copyOutMenuCustoms(rows []interface{}) []*MenuCustom {
	size := len(rows)
	result := make([]*MenuCustom, size)
	for i := 0; i < size; i++ {
		result[i] = rows[i].(*MenuCustom)
	}
	return result
}

//func (db Database) JoinedTaxonomyTerms(nid int32) []*JoinedTaxonomyTerm {
//	sql := `select idx.Nid, t.Name, v.Name as Vocab
//	    from %staxonomy_index as idx
//	    inner join %staxonomy_term_data as t on idx.tid = t.tid
//	    join %staxonomy_vocabulary as v on t.vid = v.vid
//	    where idx.Nid = ?`
//	s2 := fmt.Sprintf(sql, db.Prefix, db.Prefix, db.Prefix)
//	list, err := db.DbMap.Select(JoinedTaxonomyTerm{}, s2, nid)
//	util.CheckErrFatal(err, s2)
//	return copyOutTaxonomyTerms(list)
//}
//
//func copyOutTaxonomyTerms(rows []interface{}) []*JoinedTaxonomyTerm {
//	size := len(rows)
//	result := make([]*JoinedTaxonomyTerm, size)
//	for i := 0; i < size; i++ {
//		result[i] = rows[i].(*JoinedTaxonomyTerm)
//	}
//	return result
//}

