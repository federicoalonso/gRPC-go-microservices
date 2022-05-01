package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/federicoalonso/gRPC-go-microservices/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Printf("GreetEveryone function was invoked with a streaming request")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while calling GreetEveryone RPC: %v", err)
	}

	requests := []*pb.GreetRequest{
		{FirstName: "Federico"},
		{FirstName: "Alonso"},
		{FirstName: "Micaela"},
		{FirstName: "Rafael"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range requests {
			stream.Send(req)
			log.Printf("Sent request: %v\n", req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			log.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)
	}()

	<-waitc
}
