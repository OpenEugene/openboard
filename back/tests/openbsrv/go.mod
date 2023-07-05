module github.com/OpenEugene/openboard/back/tests/openbsrv

go 1.12

replace github.com/OpenEugene/openboard/back/internal => ../../internal

require (
	github.com/OpenEugene/openboard/back/internal v0.0.0-20191127022437-d710cfcd0696
	github.com/golang/protobuf v1.5.2
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
)
