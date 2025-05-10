package proto

//go:generate bash -c "protoc --go_out=../../pkg/pb --go_opt=paths=source_relative *.proto"
//go:generate bash -c "protoc --go-grpc_out=../../pkg/pb --go-grpc_opt=paths=source_relative *.proto"
