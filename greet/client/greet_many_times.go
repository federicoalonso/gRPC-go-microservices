package main

import (
	"context"
	"io"
	"log"

	pb "github.com/federicoalonso/gRPC-go-microservices/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Printf("doGreetManyTimes function was invoked\n")

	req := &pb.GreetRequest{
		FirstName: "Federico",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("GreetManyTimes: %s", msg.Result)
	}
}
