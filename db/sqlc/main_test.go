package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/techschool/simplebank/util"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("Cannot load configs: ", err)
	}
	testDB, err = sql.Open(config.DbDriver, config.DBSoruce)
	if err != nil {
		log.Fatal("Cannot connected to the database: ", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
