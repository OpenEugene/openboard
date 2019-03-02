package main

import (
	"fmt"
	"os"
	"path"

	"github.com/champagneabuelo/openboard/back/pb"
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
	sm := sigmon.New(nil)
	sm.Start()
	defer sm.Stop()

	gsrv, err := newGRPCSrv(":4242")
	if err != nil {
		return err
	}

	/*fsrv, err := newFrontSrv(":4244")
	if err != nil {
		return err
	}*/

	m := newServerMgmt(gsrv)

	sm.Set(func(s *sigmon.State) {
		if err := m.stop(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	})

	fmt.Println(pb.UserResp{})
	fmt.Println("to gracefully stop the application, send signal like TERM (CTRL-C) or HUP")

	if err := m.serve(); err != nil {
		return err
	}

	return fmt.Errorf("not a real error; demonstrating output")
}
