package main

//go:generate go install -v google.golang.org/protobuf/cmd/protoc-gen-go
//go:generate go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc

//go:generate go run ./eprotogen/main.go -pwd .
