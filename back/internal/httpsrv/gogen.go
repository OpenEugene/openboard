package httpsrv

// alias protobuf/swagger
//go:generate -command pb-sw protoc -I ../../../msgs/proto -I ../../../msgs/proto/vendor-extra/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true,allow_merge=true:.

// generate swagger.json
//go:generate pb-sw ../../../msgs/proto/auth.proto ../../../msgs/proto/user.proto

// embed json
//go:generate go-bindata -nocompress -pkg=swagger -o ./internal/swagger/bindata.go ./apidocs.swagger.json

// remove json
//go:generate withdraw ./apidocs.swagger.json
