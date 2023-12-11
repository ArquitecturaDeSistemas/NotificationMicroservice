package basedatos

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//var db *sql.DB

type MySQLConnection struct {
	connection *sql.DB
}

func (db *MySQLConnection) connect() {
	fmt.Println("intentando coneccion MySQL")
	conexion, err := sql.Open("mysql", DSN)
	if err != nil {
		panic(err)
	}
	db.connection = conexion
	fmt.Println(db.connection)
}

func (db *MySQLConnection) getDB() any {
	fmt.Println("intentando recobrar MySQL")
	return db.connection
}

func (db *MySQLConnection) isok() (bool, error) {
	query := "SELECT table_name FROM information_schema.tables WHERE table_name = ?"
	var tableName string = DBTABLE
	var existe string
	err := db.connection.QueryRow(query, tableName).Scan(&existe)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("Base de Datos no existe")
		return false, nil // La tabla no existe
	case err != nil:
		fmt.Println("Error en ejecucion de consulta")
		return false, err // Error al ejecutar la consulta
	case tableName != existe:
		fmt.Println("Tabla no existe")
		return false, nil // La tabla no existe
	default:
		fmt.Println("Tabla si existe")
		return true, nil // La tabla existe
	}
}

/*
const dbnombre, tabla string = "cartdb", "carrito"
const url = "user:password@tcp(127.0.0.1:3306)/" + dbnombre
*/
/*
func Connect() {
	conexion, err := sql.Open("mysql", DSN)
	if err != nil {
		panic(err)
	}
	db = conexion
	Ping()
}

func Close() {
	fmt.Println("Conexion  terminada")
	db.Close()
}

func Ping() {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
}

func IsTable() (bool, error) {
	query := "SELECT table_name FROM information_schema.tables WHERE table_name = ?"
	var tableName string = DBTABLE
	var existe string
	err := db.QueryRow(query, tableName).Scan(&existe)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("Base de Datos no existe")
		return false, nil // La tabla no existe
	case err != nil:
		fmt.Println("Error en ejecucion de consulta")
		return false, err // Error al ejecutar la consulta
	case tableName != existe:
		fmt.Println("Tabla no existe")
		return false, nil // La tabla no existe
	default:
		fmt.Println("Tabla si existe")
		return true, nil // La tabla existe
	}
}
*/
