package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	PROTOCOL = "tcp"
	IP       = "127.0.0.1"
	PORT     = "12345"
)

func main() {
	fmt.Println("Iniciando Servidor...")

	listen, err := net.Listen(PROTOCOL, ":"+PORT)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	//s := Calculator.Server{}
	servergRPC := grpc.NewServer()

	if err := servergRPC.Serve(listen); err != nil {
		log.Fatalf("Failed to Serve: %s", err)
	}
}
