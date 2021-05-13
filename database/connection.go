package database

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func CreateDbConnection(dataSourceName string)  {
	db,err:=sql.Open("sql",dataSourceName)

	if err==nil {
		log.Fatal(err)
		panic("Error occured in CreateDbConnection :")
	}

	Db=db
}