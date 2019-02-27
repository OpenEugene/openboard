package main

import (
	"fmt"
	"os"
	"path"

	"github.com/champagneabuelo/openboard/back/authsvc"
	"github.com/champagneabuelo/openboard/back/grpcsrv"
	"github.com/champagneabuelo/openboard/back/pb"
	"github.com/champagneabuelo/openboard/back/usersvc"
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

	auth, err := authsvc.New()
	if err != nil {
		return err
	}

	user, err := usersvc.New()
	if err != nil {
		return err
	}

	srv, err := grpcsrv.New()
	if err != nil {
		return err
	}

	if err = srv.RegisterServices(auth, user); err != nil {
		return err
	}

	sm.Set(func(s *sigmon.State) {
		srv.GracefulStop()
	})

	fmt.Println("to gracefully stop the server, send signal like TERM (CTRL-C) or HUP")
	if err = srv.Serve(":4242"); err != nil {
		return err
	}

	fmt.Println(pb.UserResp{})

	return fmt.Errorf("not a real error; demonstrating output")
}
