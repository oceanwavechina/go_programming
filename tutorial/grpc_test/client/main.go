package main

import (
	pb "go_programming.git/tutorial/grpc_test/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
)

const (
	address = "127.0.0.1:50001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewToUpperClient(conn)

	name := "hello world"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	r, err := c.Upper(context.Background(), &pb.UpperRequest{Name: name})
	if err != nil {
		log.Fatalf("cloud not greet: %v", err)
	}

	log.Printf("Response: %s", r.Message)
}