package model

import (
	"drupal2hugo/util"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
)

const (
	NoSuchTable uint16 = 1146 // http://dev.mysql.com/doc/refman/5.5/en/error-messages-server.html
)

type BookPage struct {
	Mlid int32
	Nid  int32
	Bid  int32
}

type Book struct {
	Bid   int32
	Title string
}

type MenuCustom struct {
	MenuName    string `db:"menu_name"`
	Title       string
	Description string
}

func (db Database) AllBookPagesAsMap() map[int32]*BookPage {
	result := make(map[int32]*BookPage)
	for _, book := range db.AllBookPages() {
		if x, has := result[book.Nid]; has {
			fmt.Fprintf(os.Stderr, "Warning: duplicate book entries %+v and %+v\n", book, x)
		}
		result[book.Nid] = book
	}
	return result
}

func (db Database) AllBookPages() []*BookPage {
	sql := "select * from " + db.Prefix + "book"
	list, err := db.DbMap.Select(BookPage{}, sql)
	if hasError(err, sql) {
		return []*BookPage{}
	}
	return copyOutBookPage(list)
}

func copyOutBookPage(rows []interface{}) []*BookPage {
	size := len(rows)
	result := make([]*BookPage, size)
	for i := 0; i < size; i++ {
		result[i] = rows[i].(*BookPage)
	}
	return result
}

func (db Database) AllBooksAsMap() map[int32]string {
	result := make(map[int32]string)
	for _, book := range db.AllBooks() {
		result[book.Bid] = book.Title
	}
	return result
}

func (db Database) AllBooks() []*Book {
	sql := `select distinct b.Bid, n.Title
	    from %sbook as b
	    join %snode as n on b.Bid = n.Nid`
	s2 := fmt.Sprintf(sql, db.Prefix, db.Prefix)
	list, err := db.DbMap.Select(Book{}, s2)
	if hasError(err, s2) {
		return []*Book{}
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

type Menu struct {
	MenuName    string
	Mlid        int32
	Plid        int32
	LinkPath    string
	LinkTitle   string
	Module      string
	External    bool
	HasChildren bool
	Expanded    bool
	Weight      int
	TreeDepth   int
	Customized  bool
}

type JoinedMenu struct {
	MenuName    string
	Title       string
	Mlid        int32
	Plid        int32
	LinkPath    string
	LinkTitle   string
	Module      string
	External    bool
	HasChildren bool
	Expanded    bool
	Weight      int
	TreeDepth   int
	Customized  bool
}

func (db Database) MenusForMlid(mlid int32) []*JoinedMenu {
	sql := `select menu_name as MenuName,
	               mlid, plid, link_path as LinkPath, link_title as LinkTitle,
	               module, external, has_children as HasChildren, expanded,
	               weight, depth as TreeDepth, customized
	    from %smenu_links
	    where hidden = 0 and mlid = ?`
	s2 := fmt.Sprintf(sql, db.Prefix)
	list, err := db.DbMap.Select(Menu{}, s2, mlid)
	util.CheckErrFatal(err, s2)
	return convertMenu(copyOutMenu(list))
}

func copyOutMenu(rows []interface{}) []*Menu {
	size := len(rows)
	result := make([]*Menu, size)
	for i := 0; i < size; i++ {
		result[i] = rows[i].(*Menu)
	}
	return result
}

func convertMenu(menus []*Menu) []*JoinedMenu {
	size := len(menus)
	result := make([]*JoinedMenu, size)
	for i, m := range menus {
		result[i] = &JoinedMenu{m.MenuName, "", m.Mlid, m.Plid, m.LinkPath, m.LinkTitle, m.Module, m.External,
			m.HasChildren, m.Expanded, m.Weight, m.TreeDepth, m.Customized}
	}
	return result
}

func (db Database) JoinedMenusForPath(path string) []*JoinedMenu {
	sql := `select c.menu_name as MenuName, c.title,
	               m.mlid, m.plid, m.link_path as LinkPath, m.link_title as LinkTitle,
	               m.module, m.external, m.has_children as HasChildren, m.expanded,
	               m.weight, m.depth as TreeDepth, m.customized
	    from %smenu_links as m
	    join %smenu_custom as c on m.menu_name = c.menu_name
	    where m.hidden = 0 and m.link_path = ?`
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

func hasError(err error, sql string) bool {
	if err != nil {
		switch err := err.(type) {
		case *mysql.MySQLError:
			if err.Number == NoSuchTable {
				return true
			}
			util.CheckErrFatal(err, sql)
		default:
			util.CheckErrFatal(err, sql)
		}
	}
	return false
}
