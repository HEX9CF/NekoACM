package rpc

import (
	"google.golang.org/grpc"
	"neko-acm/internal/interfaces/rpc/handler"
	"neko-acm/internal/interfaces/rpc/interceptors"
	"neko-acm/pkg/config"
	"neko-acm/pkg/pb"
	"net"
)

func InitServer() error {
	conf := config.Conf.Grpc

	listen, err := net.Listen("tcp", ":"+conf.Port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.TokenAuthInterceptor(),
		),
	)

	// 注册服务
	pb.RegisterProblemServiceServer(grpcServer, &handler.ProblemServer{})
	pb.RegisterTestcaseServiceServer(grpcServer, &handler.TestcaseServer{})
	pb.RegisterSolutionServiceServer(grpcServer, &handler.SolutionServer{})
	pb.RegisterAssistantServiceServer(grpcServer, &handler.AssistantServer{})
	pb.RegisterJudgeServiceServer(grpcServer, &handler.JudgeServer{})
	pb.RegisterMiscServiceServer(grpcServer, &handler.MiscServer{})

	err = grpcServer.Serve(listen)
	if err != nil {
		return err
	}

	return nil
}
