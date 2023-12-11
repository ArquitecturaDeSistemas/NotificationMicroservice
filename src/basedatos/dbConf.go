package basedatos

import (
	"github.com/joho/godotenv"
)

var envar = func() map[string]string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	m, err := godotenv.Read()
	if err != nil {
		panic(err)
	}
	return m
}()

// var DSN = "user:password@tcp(127.0.0.1:3306)/" + DBNOMBRE
var DSN = envar["DB_USER"] + ":" + envar["DB_PASSWORD"] + "@" + envar["DB_PROT"] +
	"(" + envar["DB_SERVER"] + ")/" + envar["DB_NAME"]

type TipoConeccion interface {
	connect()
	getDB() any
	isok() (bool, error)
	close()
	ping()
}

func Close(t TipoConeccion) {
	t.close()
}

func Connect(t TipoConeccion) {
	t.connect()
}

func GetDB(t TipoConeccion) any {
	return t.getDB()
}

func IsOk(t TipoConeccion) (bool, error) {
	return t.isok()
}

func Ping(t TipoConeccion) {
	t.ping()
}
