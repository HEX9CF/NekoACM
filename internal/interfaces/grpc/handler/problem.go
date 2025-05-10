package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"neko-acm/pkg/pb"
)

// implement UnimplementedProblemServiceServer
type ProblemServer struct {
	pb.UnimplementedProblemServiceServer
}

func (s *ProblemServer) GenerateProblem(ctx context.Context, req *pb.ProblemInstructionRequest) (*pb.ProblemResponse, error) {
	return &pb.ProblemResponse{
		Title: "New" + req.Title,
	}, nil
}

func (s *ProblemServer) TranslateProblem(ctx context.Context, req *pb.TranslateInstructionRequest) (*pb.ProblemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TranslateProblem not implemented")
}

func (s *ProblemServer) ParseProblem(ctx context.Context, req *pb.ProblemDataRequest) (*pb.ProblemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseProblem not implemented")
}
