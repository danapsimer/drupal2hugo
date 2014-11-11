package model

import (
	"github.com/rickb777/gorp"
	"drupal2hugo/util"
	"fmt"
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
	sql := "select * from " + prefix + "node"
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
	Type        string
	Name        string
	Base        string
	Module      string
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

func AllNodeTypes(dbMap *gorp.DbMap, prefix string) []*NodeType {
	sql := "select type, name, base, module from " + prefix + "node_type"
	list, err := dbMap.Select(NodeType{}, sql)
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
	Nid          int32
	Vid          int32
	Type         string
	Title        string
	Published    bool // column=status
	Created      int64
	Changed      int64
	Comment      int8
	Promote      bool
	Sticky       bool
	Bundle       string
	Deleted      bool
	RevisionId   int32
	Delta        int32
	BodyValue    string
	BodySummary  string
	BodyFormat   string
}

func JoinedNodeFields(dbMap *gorp.DbMap, prefix string) []*JoinedNodeDataBody {
	sql := `select
	    Nid, Vid, Type, Title, status as Published, Created, Changed, Comment,
	    Promote, Sticky, Bundle, Deleted, Revision_Id as RevisionId,
	    Delta, Body_Value as BodyValue, Body_Summary as BodySummary, Body_Format as BodyFormat
	    from %snode inner join %sfield_data_body on %snode.nid = %sfield_data_body.entity_id limit 10`
	s2 := fmt.Sprintf(sql, prefix, prefix, prefix, prefix)
	list, err := dbMap.Select(JoinedNodeDataBody{}, s2)
	util.CheckErrFatal(err, sql)
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
