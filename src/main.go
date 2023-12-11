package main

import (
	"fmt"
	"notificationmicroservice/basedatos"

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
	puerto := "0.0.0.0:50051"
	listener, err := net.Listen("tcp", puerto)
	if err != nil {
		log.Fatalf("No se pudo crear Listener: %s", err)
	}
	fmt.Printf("Servidor activo en puerto %s", puerto)

	serv := grpc.NewServer()

	pb.RegisterNotificationServiceServer(serv, &Server{})

	if err = serv.Serve(listener); err != nil {
		log.Fatalf("Fallo en servir: %s", err)
	}
}
