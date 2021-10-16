package main

import (
	pb "go_programming.git/tutorial/grpc_test/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strings"
)

const (
	port = ":50001"
)

type server struct{
	pb.UnsafeToUpperServer	// 为啥一定要加这个？？
}

func (s *server) Upper(ctx context.Context, in *pb.UpperRequest) (*pb.UpperReply, error) {
	log.Printf("received: %s ", in.Name)
	return &pb.UpperReply{Message: strings.ToUpper(in.Name)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterToUpperServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


