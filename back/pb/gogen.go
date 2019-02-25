package pb

// alias protobuf/grpc, protobuf/grpc-gateway
//go:generate -command pb-go protoc -I../../msgs/proto -I../../msgs/proto/vendor-extra/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:.
//go:generate -command pb-gg protoc -I../../msgs/proto -I../../msgs/proto/vendor-extra/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:.

// generate grpc, grpc-gateway
//go:generate pb-go ../../msgs/proto/auth.proto ../../msgs/proto/user.proto
//go:generate pb-gg ../../msgs/proto/auth.proto ../../msgs/proto/user.proto
