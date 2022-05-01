package main

import (
	"io"
	"log"

	pb "github.com/federicoalonso/gRPC-go-microservices/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone function was invoked with a streaming request")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		firstName := req.FirstName
		result := "Hello " + firstName
		log.Printf("Responding with: %v\n", result)

		err = stream.Send(&pb.GreetResponse{
			Result: result,
		})
		if err != nil {
			log.Fatalf("Error while sending data to client: %v", err)
			return err
		}
	}
}
