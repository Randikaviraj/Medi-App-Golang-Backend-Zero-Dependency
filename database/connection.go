package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func CreateDbConnection(dataSourceName string)  {

	db,err:=sql.Open("mysql",dataSourceName)
	
	if err!=nil {
		log.Fatal(err.Error())
		panic("Error occured in CreateDbConnection :")
	}

	db.SetConnMaxLifetime(time.Second * 5)
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(151)


	Db=db
}