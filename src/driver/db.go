package driver

import (
	"github.com/devopscorner/golang-adot/src/config"
	"github.com/guregu/dynamo"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	DYN *dynamo.DB
)

func ConnectDatabase() {
	if config.DbConnection() == "mysql" {
		ConnectMySQL()
		DB = DB_MySQL
	} else if config.DbConnection() == "postgres" {
		ConnectPSQL()
		DB = DB_PSQL
	} else if config.DbConnection() == "dynamo" {
		ConnectDynamo()
		DYN = DB_Dynamo
	} else {
		ConnectSQLite()
		DB = DB_SQLite
	}
}
