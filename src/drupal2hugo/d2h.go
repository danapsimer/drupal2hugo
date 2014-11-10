package main

import (
	"flag"
	"github.com/rickb777/gorp"
	"database/sql"
	"drupal2hugo/util"
	"log"
	"os"
	"errors"
	"drupal2hugo/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// see https://github.com/go-sql-driver/mysql#readme
const parameters = ""

var db = flag.String("db", "", "Drupal database name")
var driver = flag.String("driver", "mysql", "SQL driver")
var prefix = flag.String("prefix", "drp_", "Drupal table prefix")
var user = flag.String("user", "", "Drupal user (defaults to be the same as the Drupal database name)")
var pass = flag.String("pass", "", "Drupal password (you will be prompted for the password if this is absent)")
//var dir = flag.String("dir", "", "Run in directory")
//var force = flag.Bool("f", false, "Force overwriting existing files")
var verbose = flag.Bool("v", false, "Verbose")

var traceLog = "-"

func controlTrace(trace bool, DbMap *gorp.DbMap) {
	if trace {
		dbTraceWriter := util.ConstructSomeLogWriter(traceLog, os.Stdout)
		DbMap.TraceOn("", log.New(dbTraceWriter, "gorptest: ", log.Lmicroseconds))
	} else {
		DbMap.TraceOff()
	}
}

func chooseDialect() gorp.Dialect {
	switch *driver {
	case "mysql": // already initialised
		return gorp.MySQLDialect{"InnoDB", "UTF8"}
	case "sqlite":
		return gorp.SqliteDialect{}
	case "postgres":
		return gorp.PostgresDialect{}
	default:
		panic(errors.New(*driver + ": unknown database driver."))
	}
	return nil
}

func connect(driver, connection string) (database *gorp.DbMap) {
	db, err := sql.Open(driver, connection+parameters)
	util.CheckErrFatal(err, "Opening", connection)
	//db.SetMaxIdleConns(30)

	database = &gorp.DbMap{Db: db, Dialect: chooseDialect()}
	controlTrace(*verbose, database)
	return
}

// http://blog.golang.org/profiling-go-programs
// use "go tool pprof" after program termination

func main() {
	flag.Parse()
	if *user == "" {
		*user = *db
	}

	// username:password@protocol(address)/dbname?param=value
	db := connect(*driver, *user+":"+*pass+"@/"+*db)

	for _, nt := range model.AllNodeTypes(db, *prefix) {
		fmt.Printf("%v\n", nt)
	}

	for _, node := range model.AllNodes(db, *prefix) {
		fmt.Printf("%v\n", node)
	}
}


