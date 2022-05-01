package main

import (
	"context"
	"log"
	"time"

	pb "github.com/federicoalonso/gRPC-go-microservices/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("Starting to do a long greet RPC...")

	reqs := []*pb.GreetRequest{
		{FirstName: "Federico"},
		{FirstName: "Micaela"},
		{FirstName: "Lorenzo"},
		{FirstName: "Donatelo"},
		{FirstName: "Rafael"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		log.Printf("Sending: %v\n", req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}

	log.Printf("LongGreet Response: %v\n", res.Result)
}
