package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/codemodus/sigmon/v2"
)

func main() {
	if err := run(); err != nil {
		cmd := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "%s: %s\n", cmd, err)
		os.Exit(1)
	}
}

func run() error {
	var (
		dbdrvr   = "mysql"
		dbname   = "openeug_openb_dev"
		dbuser   = "openeug_openbdev"
		dbpass   = ""
		dbaddr   = "127.0.0.1"
		dbport   = ":3306"
		migrate  bool
		rollback bool
		frontDir = "../../../front/public"
	)

	flag.StringVar(&dbname, "dbname", dbname, "database name")
	flag.StringVar(&dbuser, "dbuser", dbuser, "database user")
	flag.StringVar(&dbpass, "dbpass", dbpass, "database pass")
	flag.StringVar(&dbaddr, "dbaddr", dbaddr, "database addr")
	flag.StringVar(&dbport, "dbport", dbport, "database port")
	flag.BoolVar(&migrate, "migrate", migrate, "migrate up")
	flag.BoolVar(&rollback, "rollback", rollback, "migrate dn")
	flag.StringVar(&frontDir, "frontdir", frontDir, "front public assets directory")
	flag.Parse()

	sm := sigmon.New(nil)
	sm.Start()
	defer sm.Stop()

	db, err := newSQLDB(dbdrvr, dbCreds(dbname, dbuser, dbpass, dbaddr, dbport))
	if err != nil {
		return err
	}

	mig, err := newDBMig(db, dbdrvr)
	if err != nil {
		return err
	}

	gsrv, err := newGRPCSrv(":4242", db)
	if err != nil {
		return err
	}

	mig.addMigrators(gsrv.services()...)
	if mres, migType := mig.run(migrate, rollback); len(migType) > 0 {
		if mres.HasError() {
			return mres.ErrsErr()
		}
		fmt.Println(migType+":", mres)
	}

	hsrv, err := newHTTPSrv(nil, ":4242", ":4243")
	if err != nil {
		return err
	}

	fsrv, err := newFrontSrv(nil, frontDir, ":4244")
	if err != nil {
		return err
	}

	m := newServerMgmt(gsrv, hsrv, fsrv)

	sm.Set(func(s *sigmon.State) {
		if err := m.stop(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	})

	fmt.Println("to gracefully stop the application, send signal like TERM (CTRL-C) or HUP")

	if err := m.serve(); err != nil {
		return err
	}

	return fmt.Errorf("not a real error; demonstrating output")
}
