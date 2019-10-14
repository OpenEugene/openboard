package main

import (
	"context"
	"fmt"
	"github.com/OpenEugene/openboard/back/internal/pb"
	"github.com/codemodus/relay"
	"google.golang.org/grpc"
)

func main() {
	check, ff := relay.New().Fns()
	defer func() { ff(recover()) }()

	cc, err := grpc.Dial(":4242", grpc.WithInsecure())
	check(err)

	defer cc.Close()

	c := pb.NewUserSvcClient(cc)

	r, err := c.AddRole(
		context.Background(),
		&pb.AddRoleReq{
			Name: "testRole",
		},
	)
	check(err)

	fmt.Printf("Response from server: %s", r)
}
