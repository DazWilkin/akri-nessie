package main

import (
	"context"
	"flag"
	"log"
	"net"

	pb "github.com/DazWilkin/akri-nessie/protos"
	"google.golang.org/grpc"
)

var _ pb.NessieServer = (*Server)(nil)

var (
	grpcEndpoint = flag.String("grpc_endpoint", "", "The endpoint of this gRPC server.")
)

// Server is a type that implements pb.NessieServer
type Server struct{}

// NewServer is a function that creates a new Server
func NewServer() *Server {
	return &Server{}
}

// GetNessieNow is a method that implements the pb.NessieServer interface
func (s *Server) GetNessieNow(ctx context.Context, rqst *pb.NotifyRequest) (*pb.NotifyResponse, error) {
	return &pb.NotifyResponse{
		Frame: []byte("Freddie"), //{0x46, 0x72, 0x65, 0x64, 0x64, 0x69, 0x65}, // Freddie
	}, nil
}
func main() {
	log.Println("[main] Starting gRPC server")

	flag.Parse()
	if *grpcEndpoint == "" {
		log.Fatal("[main] Unable to start server. Requires gRPC endpoint.")
	}

	serverOpts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(serverOpts...)
	pb.RegisterNessieServer(grpcServer, NewServer())

	listen, err := net.Listen("tcp", *grpcEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[main] Starting gRPC Listener [%s]\n", *grpcEndpoint)
	log.Fatal(grpcServer.Serve(listen))
}
