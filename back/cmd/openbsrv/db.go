package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/codemodus/sqlmig"
	_ "github.com/go-sql-driver/mysql"
)

func newSQLDB(driver, creds string) (*sql.DB, error) {
	db, err := sql.Open(driver, creds)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(128)
	db.SetConnMaxLifetime(time.Hour)

	return db, patientPing(db)
}

func patientPing(db *sql.DB) error {
	limit := time.Second * 3
	pause := time.Millisecond * 500
	iters := 6 // 0 + .5 + 1 + 2 + 4 + 8
	var err error

	for i := 0; i < iters; i++ {
		if i > 0 {
			time.Sleep(pause)
			pause = pause * 2
		}

		func() {
			ctx, cancel := context.WithTimeout(context.Background(), limit)
			defer cancel()

			err = db.PingContext(ctx)
		}()

		if err == nil {
			return nil
		}
	}

	return err
}

type dbmig struct {
	m *sqlmig.SQLMig
}

func newDBMig(db *sql.DB, driver, tablePrefix string) (*dbmig, error) {
	mig, err := sqlmig.New(db, driver, tablePrefix)
	if err != nil {
		return nil, err
	}

	dbm := dbmig{
		m: mig,
	}

	return &dbm, nil
}

func (m *dbmig) addMigrators(us ...interface{}) {
	for _, u := range us {
		if p, ok := u.(sqlmig.DataProvider); ok {
			m.m.AddDataProviders(p)
		}

		if r, ok := u.(sqlmig.Regularizer); ok {
			m.m.AddRegularizers(r)
		}
	}
}

func (m *dbmig) run(migrate, rollback bool) (sqlmig.Results, string) {
	switch {
	case migrate && rollback:
		return sqlmig.Results{}, ""
	case migrate:
		return m.m.Migrate(), "migrated"
	case rollback:
		return m.m.RollBack(), "rolled back"
	default:
		return sqlmig.Results{}, ""
	}
}

func dbCreds(name, user, pass, addr, port string) string {
	return fmt.Sprintf("%s:%s@tcp(%s%s)/%s?parseTime=true", user, pass, addr, port, name)
}
