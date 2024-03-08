package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "gRPC/calculator"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)

	addResp, err := c.Add(context.Background(), &pb.AddRequest{A: 10, B: 5})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	fmt.Printf("Add result: %d\n", addResp.Result)
}
