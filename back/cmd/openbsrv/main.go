package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/codemodus/sigmon/v2"

	"github.com/OpenEugene/openboard/back/internal/dbg"
	"github.com/OpenEugene/openboard/back/internal/log"
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
		dbdrvr    = "mysql"
		dbname    = "openeug_openb_dev"
		dbuser    = "openeug_openbdev"
		dbpass    = ""
		dbaddr    = "127.0.0.1"
		dbport    = ":3306"
		migrate   bool
		rollback  bool
		skipsrv   bool
		debug     bool
		frontDir  = "../../../front/public"
		migTblPfx = "mig_"
	)

	flag.StringVar(&dbname, "dbname", dbname, "database name")
	flag.StringVar(&dbuser, "dbuser", dbuser, "database user")
	flag.StringVar(&dbpass, "dbpass", dbpass, "database pass")
	flag.StringVar(&dbaddr, "dbaddr", dbaddr, "database addr")
	flag.StringVar(&dbport, "dbport", dbport, "database port")
	flag.BoolVar(&migrate, "migrate", migrate, "migrate up")
	flag.BoolVar(&rollback, "rollback", rollback, "migrate dn")
	flag.BoolVar(&skipsrv, "skipsrv", skipsrv, "skip server run")
	flag.BoolVar(&debug, "debug", debug, "debug true or false")
	flag.StringVar(&frontDir, "frontdir", frontDir, "front public assets directory")
	flag.Parse()

	sm := sigmon.New(nil)
	sm.Start()
	defer sm.Stop()

	logCfg := log.Config{
		Err: log.Output{
			Out:    os.Stdout,
			Prefix: "[ERROR] ",
		},
		Inf: log.Output{
			Out:    os.Stdout,
			Prefix: "[INFO] ",
		},
	}
	log := log.New(logCfg)

	if debug {
		dbg.SetDebugOut(os.Stdout)
	}

	dbg.Logf("set up SQL database at %s:%s.", dbaddr, dbport)
	db, err := newSQLDB(dbdrvr, dbCreds(dbname, dbuser, dbpass, dbaddr, dbport))
	if err != nil {
		return err
	}

	mig, err := newDBMig(db, dbdrvr, migTblPfx)
	if err != nil {
		return err
	}

	gsrv, err := newGRPCSrv(":4242", db, dbdrvr)
	if err != nil {
		return err
	}

	mig.addMigrators(gsrv.services()...)
	if mres, migType := mig.run(migrate, rollback); len(migType) > 0 {
		if mres.HasError() {
			return mres.ErrsErr()
		}
		log.Info("%s: %s", migType, mres)
	}

	if skipsrv {
		fmt.Println("servers will not be run; exiting")
		return nil
	}

	hsrv, err := newHTTPSrv(":4242", ":4243", nil)
	if err != nil {
		return err
	}

	fsrv, err := newFrontSrv(":4244", frontDir, nil)
	if err != nil {
		return err
	}

	m := newServerMgmt(gsrv, hsrv, fsrv)

	sm.Set(func(s *sigmon.State) {
		if err := m.stop(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	})

	log.Info("to gracefully stop the application, send signal like TERM (CTRL-C) or HUP")

	return m.serve()
}
