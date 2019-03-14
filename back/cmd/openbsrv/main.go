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
		frontDir = "../../../front/public"
	)

	flag.StringVar(&frontDir, "frontdir", frontDir, "front public assets directory")
	flag.Parse()

	sm := sigmon.New(nil)
	sm.Start()
	defer sm.Stop()

	gsrv, err := newGRPCSrv(":4242")
	if err != nil {
		return err
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
