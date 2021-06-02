package httpsrv

// alias protobuf/swagger
//go:generate -command pb-sw protoc -I ../../../msgs/proto -I ../../tools/.include -I ../../../msgs/proto/include/googleapis --openapiv2_out=logtostderr=true,allow_merge=true:.

// generate swagger.json
//go:generate pb-sw ../../../msgs/proto/auth.proto ../../../msgs/proto/user.proto ../../../msgs/proto/post.proto

//go:embed ./apidocs.swagger.json
