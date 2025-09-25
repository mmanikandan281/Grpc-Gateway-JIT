package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	pb "jit.io"
)

type accessServer struct {
	pb.UnimplementedAccessServiceServer
}

func (s *accessServer) RequestAccess(ctx context.Context, req *pb.AccessRequest) (*pb.AccessResponse, error) {
	log.Printf("Received AccessRequest: userId=%s, role=%s, duration=%d, justification=%s",
		req.UserId, req.Role, req.DurationMinutes, req.Justification)

	return &pb.AccessResponse{
		RequestId: "req-123",
		Status:    "approved",
	}, nil
}

func main() {
	// Start gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on :50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAccessServiceServer(grpcServer, &accessServer{})

	// Run gRPC in background goroutine
	go func() {
		log.Println("gRPC server listening on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	// Start HTTP gateway
	ctx := context.Background()
	mux := runtime.NewServeMux()
	err = pb.RegisterAccessServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:50051",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatalf("failed to register HTTP gateway: %v", err)
	}

	log.Println("HTTP server listening on :5000")
	if err := http.ListenAndServe(":5000", mux); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
