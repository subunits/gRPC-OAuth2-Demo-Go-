package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "grpc-oauth/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// ==== Server implementation ====

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello, " + req.Name + "!"}, nil
}

func startServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
	pb.RegisterGreeterServer(s, &greeterServer{})

	log.Println("Server running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// ==== Client ====

func startClient(token string) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	md := metadata.Pairs("authorization", fmt.Sprintf("Bearer %s", token))
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "OAuth2"})
	if err != nil {
		log.Fatalf("RPC failed: %v", err)
	}

	fmt.Println("Response:", resp.Message)
}

// ==== Main ====

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [server|client <token>]")
		return
	}

	switch os.Args[1] {
	case "server":
		startServer()
	case "client":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go client <token>")
			return
		}
		startClient(os.Args[2])
	default:
		fmt.Println("Unknown command")
	}
}
