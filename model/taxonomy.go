package model

import (
	"github.com/fale/drupal2hugo/util"
	"fmt"
)

type TaxonomyIndex struct {
	Nid     int32
	Tid     int32
	Sticky  bool
	Created int64
}

type JoinedTaxonomyTerm struct {
	Nid   int32
	Name  string
	Vocab string
}

type Vocabulary struct {
	Vid  int32
	Name string
}

func (db Database) AllVocabularies() []*Vocabulary {
	sql := "select vid,name from " + db.Prefix + "vocabulary"
	list, err := db.DbMap.Select(Vocabulary{}, sql)
	util.CheckErrFatal(err, sql)
	return copyOutVocabularies(list)
}

func copyOutVocabularies(rows []interface{}) []*Vocabulary {
	size := len(rows)
	result := make([]*Vocabulary, size)
	for i := 0; i < size; i++ {
		result[i] = rows[i].(*Vocabulary)
	}
	return result
}

func (db Database) JoinedTaxonomyTerms(nid int32) []*JoinedTaxonomyTerm {
	sql := `select idx.Nid, t.Name, v.Name as Vocab
	    from %sterm_node as idx
	    inner join %sterm_data as t on idx.tid = t.tid
	    join %svocabulary as v on t.vid = v.vid
	    where idx.Nid = ?`
	s2 := fmt.Sprintf(sql, db.Prefix, db.Prefix, db.Prefix)
	list, err := db.DbMap.Select(JoinedTaxonomyTerm{}, s2, nid)
	util.CheckErrFatal(err, s2)
	return copyOutTaxonomyTerms(list)
}

func copyOutTaxonomyTerms(rows []interface{}) []*JoinedTaxonomyTerm {
	size := len(rows)
	result := make([]*JoinedTaxonomyTerm, size)
	for i := 0; i < size; i++ {
		result[i] = rows[i].(*JoinedTaxonomyTerm)
	}
	return result
}
