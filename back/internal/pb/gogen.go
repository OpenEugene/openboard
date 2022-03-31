package pb

// alias protobuf/grpc, protobuf/grpc-gateway
//go:generate -command gen-pb protoc -I ../../tools/.include -I ../../../msgs/proto -I ../../../msgs/proto/include/googleapis --go_out=. --go-grpc_out=require_unimplemented_servers=false:. --grpc-gateway_out=logtostderr=true:.

// generate grpc, grpc-gateway
//go:generate gen-pb ../../../msgs/proto/auth.proto ../../../msgs/proto/user.proto ../../../msgs/proto/post.proto
