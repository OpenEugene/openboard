package main

import (
	"context"
	"fmt"
	"github.com/OpenEugene/openboard/back/internal/pb"
	"google.golang.org/grpc"
)

func main() {
	var ctx context.Context

	cc, err := grpc.Dial(":4242", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer cc.Close()

	c := pb.NewUserSvcClient(cc)

	r, err := c.AddRole(ctx, &pb.AddRoleReq{
		Name: "testRole",
	})
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Printf("Response from server: %s", r)
}
