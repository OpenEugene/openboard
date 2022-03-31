module github.com/OpenEugene/openboard/back/cmd/openbsrv

go 1.16

require (
	github.com/OpenEugene/openboard/back/internal v0.0.0
	github.com/codemodus/alfred v0.2.1
	github.com/codemodus/chain/v2 v2.1.2
	github.com/codemodus/hedrs v0.1.1
	github.com/codemodus/sigmon/v2 v2.0.1
	github.com/codemodus/sqlmig v0.2.3
	github.com/go-sql-driver/mysql v1.5.0
)

replace github.com/OpenEugene/openboard/back/internal => ../../internal
