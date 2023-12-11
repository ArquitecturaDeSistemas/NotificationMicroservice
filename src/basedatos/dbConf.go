package basedatos

const DBNOMBRE, DBTABLE string = "cartdb", "carrito"
const DSN = "user:password@tcp(127.0.0.1:3306)/" + DBNOMBRE

type TipoConeccion interface {
	connect()
	getDB() any
	isok() (bool, error)
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
