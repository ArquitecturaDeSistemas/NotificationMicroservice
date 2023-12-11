package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	repository "github.com/ArquitecturaDeSistemas/notificationmicroservice/adapters"
	service "github.com/ArquitecturaDeSistemas/notificationmicroservice/aplicacion"
	"github.com/ArquitecturaDeSistemas/notificationmicroservice/database"
	pb "github.com/ArquitecturaDeSistemas/notificationmicroservice/proto"
	"google.golang.org/grpc"
)

func main() {

	db := database.Connect()
	if db == nil {
		log.Fatalf("No se pudo conectar a la base de datos")
	}
	database.EjecutarMigraciones(db.GetConn())
	notificationRepository := repository.NewNotificacionRepository(db)
	notificationService := service.NewNotificacionService(notificationRepository)
	// Configura el servidor gRPC
	//este servidor está escuchando en el puerto 50051
	//y se encarga de registrar el servicio de usuarios
	grpcServe := grpc.NewServer()
	// Registra el servicio de usuarios en el servidor gRPC
	pb.RegisterNotificacionServiceServer(grpcServe, notificationService)

	// Define el puerto en el que se ejecutará el servidor gRPC
	port := "50052"
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Printf("Server listening on port %s...\n", port)

	// Inicia el servidor gRPC en segundo plano
	go func() {
		if err := grpcServe.Serve(listen); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Espera una señal para detener el servidor gRPC
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch

	fmt.Println("Shutting down the server...")

	// Detén el servidor gRPC de manera segura
	grpcServe.GracefulStop()
}
