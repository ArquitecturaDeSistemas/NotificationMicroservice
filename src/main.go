package main

import (
	"database/sql"
	"notificationmicroservice/basedatos"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := new(basedatos.MySQLConnection)
	basedatos.Connect(db)
	bd := basedatos.GetDB(db).(*sql.DB)
	basedatos.IsOk(db)

	/*
		basedatos.Connect()
		basedatos.IsTable()
		basedatos.Close()
	*/
}
