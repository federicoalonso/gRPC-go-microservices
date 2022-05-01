package main

import (
	"context"
	"log"

	pb "github.com/federicoalonso/gRPC-go-microservices/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Printf("----createBlog was invoked----")

	blog := &pb.Blog{
		AuthorId: "Federico Alonso",
		Title:    "My first blog",
		Content:  "Content of the first blog is:\n\t- Hello world!",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Error while calling CreateBlog: %v", err)
	}

	log.Printf("Blog has been created: %v", res.Id)
	return res.Id
}
