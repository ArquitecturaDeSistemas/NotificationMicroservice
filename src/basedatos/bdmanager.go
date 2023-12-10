package basedatos

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const url = "user:password@tcp(localhost:3306)/cartdb"

var db *sql.DB

func Connect() {
	conexion, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion")
	db = conexion
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
	var tableName = "carrito"
	var result string
	err := db.QueryRow(query, tableName).Scan(&result)

	switch {
	case err == sql.ErrNoRows:
		fmt.Println("Tabla no existe")
		return false, nil // La tabla no existe
	case err != nil:
		fmt.Println("Error de Consulta")
		return false, err // Error al ejecutar la consulta
	default:
		fmt.Println("Tabla si existe")
		return true, nil // La tabla existe
	}
}