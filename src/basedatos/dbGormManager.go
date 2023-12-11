package basedatos

import (
	//"fmt"
	//"log"
	//"time"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	//"os"
)

var gormdb *gorm.DB

var DataBase = func() (db *gorm.DB) {
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		fmt.Println("error carga Variables de Entorno")
		panic(errorVariables)
	}

	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
