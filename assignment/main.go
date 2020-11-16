package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
	"protoimport/assignment"
	"syscall"
)

func main() {
	// Starting listener
	listener, _ := net.Listen("tcp", ":50051")

	// Create new gRPC implementation
	impl := serverImplementation{}

	// Create default gRPC server and register implementation with it
	server := grpc.NewServer()
	assignment.RegisterApikeyServiceServer(server, &impl)

	// Serve the gRPC server using the previously generated listener
	go func() {
		_ = server.Serve(listener)
	}()
	log.Println("Serving @:50051")

	// Termination handling
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGTERM)
	signal.Notify(sigc, syscall.SIGINT)
	select {
	case <-sigc:
		log.Println("Terminating gracefully...")
		server.GracefulStop()
		log.Println("bye.")
	}
}

// serverImplementation implements the ApikeyServiceServer interface
type serverImplementation struct {
}

func (s serverImplementation) AddApikey(ctx context.Context, request *assignment.AddApikeyRequest) (*assignment.AddApikeyResponse, error) {
	return nil, status.Error(codes.Unimplemented, "implement me")
}

func (s serverImplementation) ListApikeys(ctx context.Context, request *assignment.ListApikeysRequest) (*assignment.ListApikeysResponse, error) {
	return nil, status.Error(codes.Unimplemented, "implement me")
}

func (s serverImplementation) GetApikey(ctx context.Context, request *assignment.GetApikeyRequest) (*assignment.GetApikeyResponse, error) {
	return nil, status.Error(codes.Unimplemented, "implement me")
}
