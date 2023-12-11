package main

import (
	"fmt"
	"notificationmicroservice/basedatos"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := new(basedatos.GormMySQLConnection)
	basedatos.Connect(db)
	bd := basedatos.GetDB(db).(*gorm.DB)
	fmt.Println(bd)
	basedatos.EjecutarMigraciones(bd)

	/*
		db := new(basedatos.MySQLConnection)
		basedatos.Connect(db)
		bd := basedatos.GetDB(db).(*sql.DB) //retorna un "any", por lo que hay que castearlo
		basedatos.IsOk(db)
		fmt.Println(bd)
		basedatos.Ping(db)
		basedatos.Close(db)
	*/
}
