module github.com/OpenEugene/openboard/back/tests/cmd/openbsrv-integration

go 1.12

require (
	github.com/OpenEugene/openboard/back/internal v0.0.0-20190926014330-eca0374adab8
	github.com/codemodus/relay v0.0.0-20190923033535-7e3d7567c3a7
	google.golang.org/grpc v1.24.0
)

replace github.com/OpenEugene/openboard/back/internal => ../../../internal
