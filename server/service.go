package main

import (
	"context"

	pb "github.com/maaaato/sample-gRPC/proto"
)

type EchoService struct{}

func (s *EchoService) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{
		Message: req.GetMessage()}, nil
}
