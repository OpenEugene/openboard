module github.com/champagneabuelo/openboard/back/cmd/openbsrv

go 1.12

require (
	github.com/champagneabuelo/openboard/back/internal v0.0.0
	github.com/codemodus/alfred v0.2.1
	github.com/codemodus/chain/v2 v2.1.2
	github.com/codemodus/hedrs v0.1.1
	github.com/codemodus/sigmon/v2 v2.0.1
	github.com/codemodus/sqlmig v0.2.3
	github.com/go-sql-driver/mysql v1.4.1
)

replace github.com/champagneabuelo/openboard/back/internal => ../../internal
