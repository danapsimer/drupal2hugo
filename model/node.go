//-----------------------------------------------------------------------------
// The MIT License
//
// Copyright (c) 2012 Rick Beton <rick@bigbeeconsultants.co.uk>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
//-----------------------------------------------------------------------------

package model

import (
	"github.com/fale/drupal2hugo/util"
	"fmt"
	"github.com/rickb777/gorp"
)

//type NodeRevision struct {
//	Nid       int32
//	Vid       int32
//	Uid       int32
//	Title     string
//	Timestamp int64
//	Status    bool // whether published
//	Comment   bool // whether allowed
//	Promote   bool
//	Sticky    bool
//}
//
//type FieldConfigInstance struct {
//	Id         int32
//	FieldId    int32
//	FieldName  string // body, comment_body, field_tags, field_image, ...
//	EntityType string // comment, node
//	Bundle     string // page, article, blog, book, ...
//}
//
//type FieldDataBody struct {
//	EntityType  string
//	Bundle      string
//	Deleted     bool
//	EntityId    int32 // -> Node.Nid
//	RevisionId  int32
//	Delta       int32
//	Language    string
//	BodyValue   string
//	BodySummary string
//	BodyFormat  string
//}
//
//type FieldDataFieldTags struct {
//	EntityType   string
//	Bundle       string
//	Deleted      bool
//	EntityId     int32
//	RevisionId   int32
//	Delta        int32
//	Language     string
//	FieldTagsTid int32
//}

type Node struct {
	Nid       int32
	Vid       int32
	Type      string
	Language  string
	Title     string
	Uid       int32
	Status    bool
	Created   int64
	Changed   int64
	Comment   int8
	Promote   bool
	Sticky    bool
	Tnid      int32
	Translate int32
}

func AllNodes(dbMap *gorp.DbMap, prefix string) []*Node {
	sql := "select nid,vid,type,language,title,uid,status,created,changed,comment,promote,sticky,tnid,translate from " + prefix + "node"
	list, err := dbMap.Select(Node{}, sql)
	util.CheckErrFatal(err, sql)
	return copyOutNode(list)
}

func copyOutNode(rows []interface{}) []*Node {
	size := len(rows)
	result := make([]*Node, size)
	for i := 0; i < size; i++ {
		result[i] = rows[i].(*Node)
	}
	return result
}

type NodeType struct {
	Type   string
	Name   string
	Module string
	// Base        string
	//	Description string
	//	Help        string
	//	HasTitle    bool
	//	TitleLabel  string
	//	Custom      bool
	//	Modified    bool
	//	Locked      bool
	//	Disabled    bool
	//	OrigType    string
}

func (db Database) AllNodeTypes() []*NodeType {
	sql := "select type, name, module from " + db.Prefix + "node_type"
	list, err := db.DbMap.Select(NodeType{}, sql)
	util.CheckErrFatal(err, sql)
	return copyOutNodeType(list)
}

func copyOutNodeType(rows []interface{}) []*NodeType {
	size := len(rows)
	result := make([]*NodeType, size)
	for i := 0; i < size; i++ {
		result[i] = rows[i].(*NodeType)
	}
	return result
}

type JoinedNodeDataBody struct {
	Nid       int32
	Vid       int32
	Type      string
	Title     string
	Published bool // column=status
	Created   int64
	Changed   int64
	Comment   int8
	Promote   bool
	Sticky    bool
	//Bundle      string
	//Deleted     bool
	//RevisionId  int32
	//Delta       int32
	BodyValue   string
	BodySummary string
	BodyFormat  string
}

func (db Database) JoinedNodeFields(offset, count int) []*JoinedNodeDataBody {
	sql := `select
	    n.Nid, n.Vid, n.Type, n.Title, n.status as Published, n.Created, n.Changed, n.Comment,
	    n.Promote, n.Sticky, nr.Body as BodyValue, nr.Teaser as BodySummary, nr.Format as BodyFormat
	    from %snode n inner join %snode_revisions nr on n.nid = nr.nid 
	      and n.vid = nr.vid limit %d,%d`
	s2 := fmt.Sprintf(sql, db.Prefix, db.Prefix, offset, count)
	list, err := db.DbMap.Select(JoinedNodeDataBody{}, s2)
	util.CheckErrFatal(err, s2)
	return copyOutJoinedNodeDataBody(list)
}

func copyOutJoinedNodeDataBody(rows []interface{}) []*JoinedNodeDataBody {
	size := len(rows)
	result := make([]*JoinedNodeDataBody, size)
	for i := 0; i < size; i++ {
		result[i] = rows[i].(*JoinedNodeDataBody)
	}
	return result
}

func (node JoinedNodeDataBody) Filename() string {
	return ""
}

type UrlAlias struct {
	Pid      int32
	Source   string
	Alias    string
	Language string
}

func (db Database) GetUrlAlias(nid int32) string {
	sql := `select pid, src as Source, dst as Alias, language  from %surl_alias where src = ? ORDER BY pid DESC`
	s2 := fmt.Sprintf(sql, db.Prefix)
	source := fmt.Sprintf("node/%d", nid)
	list, err := db.DbMap.Select(UrlAlias{}, s2, source)
	util.CheckErrFatal(err, s2)
	if len(list) > 1 {
		util.Stderr("Expected only one alias for %s but got %d.\n%+v\n", source, len(list), list)
	}
	if len(list) == 1 {
		return list[0].(*UrlAlias).Alias
	}
	return source
}
