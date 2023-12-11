package basedatos

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormMySQLConnection struct {
	connection *gorm.DB
}

func (db *GormMySQLConnection) connect() {
	fmt.Println("intentando coneccion MySQL")
	conexion, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.connection = conexion
	fmt.Println(err)
	fmt.Println(conexion)
}

func (db *GormMySQLConnection) getDB() any {
	fmt.Println("intentando recobrar MySQL")
	fmt.Println(db.connection)
	return db.connection
}

func (db *GormMySQLConnection) isok() (bool, error) {
	return false, db.connection.Error
}

func (db *GormMySQLConnection) close() {

}

func (db GormMySQLConnection) ping() {

}
