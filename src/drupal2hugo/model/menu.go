package model

import (
	"drupal2hugo/util"
	"github.com/go-sql-driver/mysql"
	"fmt"
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

type JoinedMenu struct {
	MenuName      string
	Title         string
	Mlid          int32
	Plid          int32
	LinkPath      string
	LinkTitle     string
	Module        string
	External      bool
	HasChildren   bool
	Expanded      bool
	Weight        int
	TreeDepth     int
	Customized    bool
}

func (db Database) JoinedMenus(path string) []*JoinedMenu {
	sql := `select c.menu_name as MenuName, c.title,
	               m.mlid, m.plid, m.link_path as LinkPath, m.link_title as LinkTitle,
	               m.module, m.external, m.has_children as HasChildren, m.expanded,
	               m.weight, m.depth as TreeDepth, m.customized
	    from %smenu_links as m
	    join %smenu_custom as c on m.menu_name = c.menu_name
	    where m.link_path = ? and hidden = 0`
	s2 := fmt.Sprintf(sql, db.Prefix, db.Prefix)
	list, err := db.DbMap.Select(JoinedMenu{}, s2, path)
	util.CheckErrFatal(err, s2)
	return copyOutJoinedMenu(list)
}

func copyOutJoinedMenu(rows []interface{}) []*JoinedMenu {
	size := len(rows)
	result := make([]*JoinedMenu, size)
	for i := 0; i < size; i++ {
		result[i] = rows[i].(*JoinedMenu)
	}
	return result
}

