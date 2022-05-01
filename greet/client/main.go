package main

import (
	"log"

	pb "github.com/federicoalonso/gRPC-go-microservices/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:5051"

func main() {
	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Failed loading certificates: %v", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	// doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	doGreetEveryone(c)
}
