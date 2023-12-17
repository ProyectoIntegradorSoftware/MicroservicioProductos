package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	repository "github.com/ProyectoIntegradorSoftware/MicroservicioProductos/adapters"
	service "github.com/ProyectoIntegradorSoftware/MicroservicioProductos/aplicacion"
	"github.com/ProyectoIntegradorSoftware/MicroservicioProductos/database"
	pb "github.com/ProyectoIntegradorSoftware/MicroservicioProductos/proto"
	"google.golang.org/grpc"
)

func main() {

	db := database.Connect()
	database.EjecutarMigraciones(db.GetConn())
	userRepository := repository.NewProductRepository(db)
	userService := service.NewProductService(userRepository)
	// Configura el servidor gRPC
	//este servidor está escuchando en el puerto 50051
	//y se encarga de registrar el servicio de usuarios
	grpcServe := grpc.NewServer()
	// Registra el servicio de usuarios en el servidor gRPC
	pb.RegisterProductServiceServer(grpcServe, userService)

	// Define el puerto en el que se ejecutará el servidor gRPC
	port := "50051"
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
