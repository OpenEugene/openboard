package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/codemodus/sigmon/v2"

	"github.com/OpenEugene/openboard/back/internal/logsvc"
)

var dbgLog func(string, ...interface{})

func main() {
	handle := logsvc.Handle{
		Err: os.Stderr,
		Inf: os.Stdout,
	}
	srvLog := logsvc.NewServerLog(handle)

	if err := run(srvLog); err != nil {
		cmd := path.Base(os.Args[0])
		srvLog.Error("%s: %s", cmd, err)
		os.Exit(1)
	}
}

func run(srvLog logsvc.LineLogger) error {
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

	if debug {
		dLog := log.New(os.Stdout, "[debug]", log.Ldate|log.Ltime|log.Lshortfile)

		// dbgLog is a package-level variable.
		dbgLog = func(format string, as ...interface{}) {
			dLog.Printf(format+"\n", as...)
		}
	}

	if dbgLog != nil {
		dbgLog("set up SQL database at %s:%s.", dbaddr, dbport)
	}
	db, err := newSQLDB(dbdrvr, dbCreds(dbname, dbuser, dbpass, dbaddr, dbport))
	if err != nil {
		return err
	}

	mig, err := newDBMig(db, dbdrvr, migTblPfx)
	if err != nil {
		return err
	}

	gsrv, err := newGRPCSrv(srvLog, ":4242", db, dbdrvr)
	if err != nil {
		return err
	}

	mig.addMigrators(gsrv.services()...)
	if mres, migType := mig.run(migrate, rollback); len(migType) > 0 {
		if mres.HasError() {
			return mres.ErrsErr()
		}
		srvLog.Info("%s: %s", migType, mres)
	}

	if skipsrv {
		fmt.Println("servers will not be run; exiting")
		return nil
	}

	hsrv, err := newHTTPSrv(srvLog, ":4242", ":4243", nil)
	if err != nil {
		return err
	}

	fsrv, err := newFrontSrv(srvLog, ":4244", frontDir, nil)
	if err != nil {
		return err
	}

	m := newServerMgmt(srvLog, gsrv, hsrv, fsrv)

	sm.Set(func(s *sigmon.State) {
		if err := m.stop(); err != nil {
			srvLog.Error(err.Error())
		}
	})

	srvLog.Info("to gracefully stop the application, send signal like TERM (CTRL-C) or HUP")

	return m.serve()
}
