package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if Port == "" {
		port = "50051"
	}
	listen, err := net.Listen("tcp", ":"+Port)

	if err != nil {
		log.Fatalf("Imposivel enviar dados por %v : %v", v, err)
	}
	log.Println("Conexão Estabelecida " + Port)

	//gRPC Server Instance
	grpcserver := grpc.NewServer()

	//Register CHatService
	cs := grpc.ChatServer{}
	chatserver.RegisterServicesServer(grpcserver, &cs)

	err = grpcserver.Serve()
	if err != nil {
		log.Fatalf("Error ao Iniciar gRPC server :%v", Port)
	}

}
