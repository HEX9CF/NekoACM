package grpc

import (
	"google.golang.org/grpc"
	"neko-acm/internal/interfaces/grpc/handler"
	"neko-acm/pkg/pb"
	"net"
)

func InitServer() error {
	listen, err := net.Listen("tcp", ":14516")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	// 注册服务
	pb.RegisterProblemServiceServer(grpcServer, &handler.ProblemServer{})

	err = grpcServer.Serve(listen)
	if err != nil {
		return err
	}

	return nil
}
