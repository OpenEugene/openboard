module github.com/OpenEugene/openboard/back/tests/cmd/openbsrv-integration

go 1.12

replace github.com/OpenEugene/openboard/back/internal => ../../../internal

require (
	github.com/OpenEugene/openboard/back/internal v0.0.0-20191127022437-d710cfcd0696
	google.golang.org/grpc v1.26.0
)
