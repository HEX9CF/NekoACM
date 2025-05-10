package proto

//go:generate protoc --go_out=../../pkg/pb --go_opt=paths=source_relative *.proto
//go:generate protoc --go-grpc_out=../../pkg/pb --go-grpc_opt=paths=source_relative *.proto
