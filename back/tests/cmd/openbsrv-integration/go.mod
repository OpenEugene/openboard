module github.com/OpenEugene/openboard/back/tests/cmd/openbsrv-integration

go 1.12

replace github.com/OpenEugene/openboard/back/internal => ../../../internal

require (
	github.com/OpenEugene/openboard/back/internal v0.0.0-20191127022437-d710cfcd0696
	golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135 // indirect
	google.golang.org/grpc v1.26.0
	honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc // indirect
)
