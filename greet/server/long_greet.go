package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/federicoalonso/gRPC-go-microservices/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked with a streaming request")

	result := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			// we have finished reading the client stream
			return stream.SendAndClose(&pb.GreetResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		log.Printf("Received: %v\n", req)
		result += fmt.Sprintf("Hello %s! \n", req.FirstName)
	}
}
