package pb

// alias protobuf/grpc, protobuf/grpc-gateway
//go:generate -command gen-pb protoc -I../../../msgs/proto -I../../../msgs/proto/vendor-extra/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=. --go-grpc_out=require_unimplemented_servers=false:. --grpc-gateway_out=logtostderr=true:.

// generate grpc, grpc-gateway
//go:generate gen-pb ../../../msgs/proto/auth.proto ../../../msgs/proto/user.proto ../../../msgs/proto/post.proto
