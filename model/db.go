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
	"database/sql"
	"github.com/fale/drupal2hugo/util"
	"errors"
	"github.com/rickb777/gorp"
	"log"
	"os"
)

// see https://github.com/go-sql-driver/mysql#readme
const parameters = ""

type Database struct {
	Db     *sql.DB
	DbMap  *gorp.DbMap
	Prefix string
}

var traceLog = "-"

func Connect(driver, connection, prefix string, verbose bool) Database {
	db, err := sql.Open(driver, connection+parameters)
	util.CheckErrFatal(err, "Opening", connection)
	//db.SetMaxIdleConns(30)

	database := &gorp.DbMap{Db: db, Dialect: chooseDialect(driver)}
	controlTrace(verbose, database)
	return Database{db, database, prefix}
}

func chooseDialect(driver string) gorp.Dialect {
	switch driver {
	case "mysql": // already initialised
		return gorp.MySQLDialect{"InnoDB", "UTF8"}
	case "sqlite3":
		return gorp.SqliteDialect{}
	case "postgres":
		return gorp.PostgresDialect{}
	default:
		panic(errors.New(driver + ": unknown database driver."))
	}
	return nil
}

func controlTrace(trace bool, DbMap *gorp.DbMap) {
	if trace {
		dbTraceWriter := util.ConstructSomeLogWriter(traceLog, os.Stdout)
		DbMap.TraceOn("", "", log.New(dbTraceWriter, "gorptest: ", log.Lmicroseconds))
	} else {
		DbMap.TraceOff()
	}
}
