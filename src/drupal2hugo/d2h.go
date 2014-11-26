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

package main

import (
	"flag"
	"drupal2hugo/util"
	"os"
	"drupal2hugo/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"path"
	"bufio"
	"io"
	"time"
	"strings"
	"github.com/davecgh/go-spew/spew"
)

var dbName = flag.String("db", "", "Drupal database name - required")
var driver = flag.String("driver", "mysql", "SQL driver")
var prefix = flag.String("prefix", "drp_", "Drupal table prefix")
var user = flag.String("user", "", "Drupal user (defaults to be the same as the Drupal database name)")
var pass = flag.String("pass", "", "Drupal password (you will be prompted for the password if this is absent)")

//var dir = flag.String("dir", "", "Run in directory")
//var force = flag.Bool("f", false, "Force overwriting existing files")
var verbose = flag.Bool("v", false, "Verbose")
var version = flag.Bool("V", false, "Version information")

// http://blog.golang.org/profiling-go-programs
// use "go tool pprof" after program termination

func main() {
	flag.Parse()
	if *version {
		fmt.Fprintf(os.Stderr, "Version %s\n", util.Version)
		fmt.Fprintf(os.Stderr, "Date %s rev %s branch %s\n", util.BuildDate, util.HgTip, util.HgBranch)
		fmt.Fprintf(os.Stderr, "From %s\n", util.HgPath)
		os.Exit(0)
	}

	if *dbName == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *user == "" {
		*user = *dbName
	}

	if *pass == "" {
		fmt.Printf("Password: ")
		os.Stdout.Sync()
		_, err := fmt.Scanln(pass)
		util.CheckErrFatal(err)
	}

	if !util.FileExists("content") {
		fmt.Fprintln(os.Stderr, "There is no content directory here. Did you mean to try somewhere else?")
		os.Exit(2)
	}

	// username:password@protocol(address)/dbname?param=value
	db := model.Connect(*driver, *user+":"+*pass+"@/"+*dbName, *prefix, *verbose)

	for _, nt := range db.AllNodeTypes() {
		fmt.Printf("%v\n", nt)
	}
	for _, b := range db.AllBooks() {
		fmt.Printf("%v\n", b)
	}
	for _, m := range db.AllMenus() {
		fmt.Printf("%v\n", m)
	}
	processVocabs(db)

	//	for _, node := range model.AllNodes(db, *prefix) {
	//		fmt.Printf("%v\n", node)
	//	}

	offset := 0
	nodes := db.JoinedNodeFields(offset, 10)
	for len(nodes) > 0 {
		for _, node := range nodes {
			alias := db.GetUrlAlias(node.Nid)
			terms := db.JoinedTaxonomyTerms(node.Nid)
			menus := db.JoinedMenus(fmt.Sprintf("node/%d", node.Nid))
			if len(menus) > 0 {
				spew.Printf("node/%d %s %s\n  %+v\n", node.Nid, alias, node.Title, menus)
			}
			processNode(node, alias, terms)
		}
		offset += len(nodes)
		nodes = db.JoinedNodeFields(offset, 10)
	}
	fmt.Printf("Total %d nodes.\n", offset)
}

func processVocabs(db model.Database) {
	vocabs := db.AllVocabularies()
	if len(vocabs) > 0 {
		fmt.Printf("Insert into config.yaml\n")
		fmt.Printf("-----------------------\n")
		fmt.Printf("Taxonomies:\n")
		for _, v := range vocabs {
			n := strings.ToLower(v.Name)
			fmt.Printf("  %s: \"%s\"\n", toSingular(n), n)
		}
	}
}

func processNode(node *model.JoinedNodeDataBody, alias string, terms []*model.JoinedTaxonomyTerm) {
	fileName := fmt.Sprintf("content/%s.md", alias)
	dir := path.Dir(fileName)
	if (*verbose) {
		fmt.Printf("%s %s '%s' pub=%v del=%v\n", node.Type, alias, node.Title, node.Published, node.Deleted)
		fmt.Printf("mkdir %s\n", dir)
		//		fmt.Printf("%+v\n", node)
	}

	err := os.MkdirAll(dir, os.FileMode(0755))
	util.CheckErrFatal(err, "mkdir", dir)

	tags := flattenTaxonomies(terms)
	writeFile(fileName, node, alias, tags)
}

func writeFile(fileName string, node *model.JoinedNodeDataBody, alias string, tags []string) {
	file, err := os.Create(fileName)
	util.CheckErrFatal(err, "create", fileName)

	w := bufio.NewWriter(file)
	writeFrontMatter(w, node, alias, tags)
	writeContent(w, node, alias)
	w.Flush()
	file.Close()
}

func writeFrontMatter(w io.Writer, node *model.JoinedNodeDataBody, alias string, tags []string) {
	created := time.Unix(node.Created, 0).Format("2006-01-02")
	changed := time.Unix(node.Changed, 0).Format("2006-01-02")
	fmt.Fprintln(w, "---")
	fmt.Fprintf(w, "title:       \"%s\"\n", node.Title)
	fmt.Fprintf(w, "description: \"%s\"\n", node.BodySummary)
	fmt.Fprintf(w, "type:        %s\n", node.Type)
	fmt.Fprintf(w, "date:        %s\n", created)
	if changed != created {
		fmt.Fprintf(w, "changed:     %s\n", changed)
	}
	fmt.Fprintf(w, "weight:      %d\n", node.Nid) // the node-id is normally ascending in Drupal and is always unique
	fmt.Fprintf(w, "draft:       %v\n", !node.Published)
	fmt.Fprintf(w, "promote:     %v\n", node.Promote)
	fmt.Fprintf(w, "sticky:      %v\n", node.Sticky)
	fmt.Fprintf(w, "deleted:     %v\n", node.Deleted)
	fmt.Fprintf(w, "url:         %s\n", alias)
	fmt.Fprintf(w, "aliases:     [ node/%d ]\n", node.Nid)
	for _, tag := range tags {
		fmt.Fprintf(w, "%s\n", tag)
	}
}

func writeContent(w io.Writer, node *model.JoinedNodeDataBody, alias string) {
	if node.BodySummary != "" {
		fmt.Fprintf(w, "\n# Summary:\n")
		for _, line := range strings.Split(node.BodySummary, "\n") {
			fmt.Fprintf(w, "# %s\n", line)
		}
	}

	fmt.Fprintln(w, "\n---")
	fmt.Fprintln(w, node.BodyValue)
}

func toSingular(plural string) string {
	if strings.HasSuffix(plural, "ies") {
		return string(plural[:len(plural) - 3]) + "y"
	}
	if strings.HasSuffix(plural, "s") {
		return string(plural[:len(plural) - 1])
	}
	return plural
}

func flattenTaxonomies(terms []*model.JoinedTaxonomyTerm) (result []string) {
	table := make(map[string][]string)
	for _, t := range terms {
		v := strings.ToLower(t.Vocab)
		table[v] = append(table[v], strings.ToLower(t.Name))
	}
	//fmt.Printf("taxonomies %+v\n", table)

	for t, list := range table {
		result = append(result, fmt.Sprintf("%-12s [ \"%s\" ]", t+":", strings.Join(list, "\", \"")))
	}
	return
}
