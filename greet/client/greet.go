package main

import (
	"context"
	"log"

	pb "github.com/federicoalonso/gRPC-go-microservices/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("doGreet function was invoked\n")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Federico",
	})

	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v\n", err)
	}

	log.Printf("Response from Greet: %s\n", res.Result)
}
