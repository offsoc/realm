package c2

//go:generate protoc --go_out=./c2pb --go_opt=paths=source_relative --go-grpc_out=./c2pb --go-grpc_opt=paths=source_relative c2.proto
