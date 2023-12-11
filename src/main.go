package main

import (
	"fmt"
	"notificationmicroservice/basedatos"

	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//DB GORM start
	db := new(basedatos.GormMySQLConnection)
	basedatos.Connect(db)
	bd := basedatos.GetDB(db).(*gorm.DB)
	basedatos.EjecutarMigraciones(bd)

	//gRPC start
	puerto := "8089"
	_, err := net.Listen("tcp", ":"+puerto)
	if err != nil {
		log.Fatalf("No se pudo crear Listener: %s", err)
	}
	fmt.Printf("Servidor activo en puerto %s", puerto)
}
