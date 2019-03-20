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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

type dbmig struct {
	m *sqlmig.SQLMig
}

func newDBMig(db *sql.DB, driver string) (*dbmig, error) {
	mig, err := sqlmig.New(db, driver)
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
		if qm, ok := u.(sqlmig.QueryingMigrator); ok {
			m.m.AddQueryingMigs(qm)
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
