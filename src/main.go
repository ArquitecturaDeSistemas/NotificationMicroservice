package main

import (
	"bufio"
	"fmt"
	"notificationmicroservice/basedatos"
	"os"

	"log"
	"net"

	pb "notificationmicroservice/proto"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Server struct {
	pb.NotificationServiceServer
}

func main() {

	//DB GORM start
	db := new(basedatos.GormMySQLConnection)
	basedatos.Connect(db)
	bd := basedatos.GetDB(db).(*gorm.DB)
	basedatos.EjecutarMigraciones(bd)

	//gRPC start
	puerto := "localhost:50051"
	listener, err := net.Listen("tcp", puerto)
	if err != nil {
		log.Fatalf("No se pudo crear Listener: %s", err)
	}

	fmt.Printf("Servidor activo en puerto %s", puerto)
	serverNotification := grpc.NewServer()

	if err = serverNotification.Serve(listener); err != nil {
		log.Fatalf("Error Servidor (?): %s", err)
	}

	fmt.Printf("Servidor ok %s", puerto)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
