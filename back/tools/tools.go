// +build tools

package tools

//go:generate go install github.com/codemodus/withdraw
//go:generate go install github.com/go-bindata/go-bindata/go-bindata
//go:generate go install github.com/golang/protobuf/protoc-gen-go
//go:generate go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
//go:generate go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
import (
	_ "github.com/codemodus/withdraw"
	_ "github.com/go-bindata/go-bindata/go-bindata"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
)

//go:generate go mod tidy
