package basedatos

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const dbnombre, tabla string = "microservicionotificacion", "notificacion"
const url = "user:password@tcp(127.0.0.1:3306)/" + dbnombre

var db *sql.DB

func Connect() {
	conexion, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	db = conexion
	//Ping()
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
	var tableName string = tabla
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
