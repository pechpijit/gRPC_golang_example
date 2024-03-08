package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "gRPC/calculator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type CalculatorServer struct {
	pb.UnimplementedCalculatorServer
}

func (s *CalculatorServer) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	result := in.A + in.B
	return &pb.AddResponse{Result: result}, nil
}

func (s *CalculatorServer) Subtract(ctx context.Context, in *pb.SubtractRequest) (*pb.SubtractResponse, error) {
	result := in.A - in.B
	return &pb.SubtractResponse{Result: result}, nil
}

func (s *CalculatorServer) Multiply(ctx context.Context, in *pb.MultiplyRequest) (*pb.MultiplyResponse, error) {
	result := in.A * in.B
	return &pb.MultiplyResponse{Result: result}, nil
}

func (s *CalculatorServer) Divide(ctx context.Context, in *pb.DivideRequest) (*pb.DivideResponse, error) {
	if in.B == 0 {
		return nil, fmt.Errorf("division by zero")
	}
	result := float32(in.A) / float32(in.B)
	return &pb.DivideResponse{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterCalculatorServer(s, &CalculatorServer{})
	fmt.Println("Server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
