package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/codemodus/sqlmig"
	_ "github.com/go-sql-driver/mysql"
)

func dbCreds(name, user, pass, addr, port string) string {
	return fmt.Sprintf("%s:%s@tcp(%s%s)/%s?parseTime=true", user, pass, addr, port, name)
}

type dbmig struct {
	*sqlmig.SQLMig
}

func newDBMig(driver, creds string) (*dbmig, error) {
	db, err := sql.Open(driver, creds)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(128)
	db.SetConnMaxLifetime(time.Hour)

	mig, err := sqlmig.New(db, driver)
	if err != nil {
		return nil, err
	}

	dbm := dbmig{
		SQLMig: mig,
	}

	return &dbm, nil
}

func (m *dbmig) addMigrators(us ...interface{}) {
	for _, u := range us {
		if qm, ok := u.(sqlmig.QueryingMigrator); ok {
			m.AddQueryingMigs(qm)
		}

		if r, ok := u.(sqlmig.Regularizer); ok {
			m.AddRegularizers(r)
		}
	}
}
