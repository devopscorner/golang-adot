package driver

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/devopscorner/golang-adot/src/config"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	DYN *dynamodb.DynamoDB
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
