package model

import (
	"drupal2hugo/util"
	"fmt"
)

type TaxonomyIndex struct {
	Nid     int32
	Tid     int32
	Sticky  bool
	Created int64
}

type JoinedTaxonomyTerm struct {
	Nid     int32
	Name    string
	Vocab   string
}

type Vocabulary struct {
	Vid     int32
	Name    string
}

func (db Database) AllVocabularies() []*Vocabulary {
	sql := "select vid,name from " + db.Prefix + "taxonomy_vocabulary"
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

func (db Database) JoinedTaxonomyTerms() []*JoinedTaxonomyTerm {
	sql := `select
	    idx.Nid, t.Name, v.Name as Vocab
	    from %staxonomy_index idx
	    inner join %staxonomy_term_data t on idx.tid = t.tid
	    inner join %staxonomy_term_hierarchy v on t.vid = v.vid`
	s2 := fmt.Sprintf(sql, db.Prefix, db.Prefix, db.Prefix)
	list, err := db.DbMap.Select(JoinedTaxonomyTerm{}, s2)
	util.CheckErrFatal(err, sql)
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

