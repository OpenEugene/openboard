package main

import (
	"context"
	"fmt"
	"github.com/OpenEugene/openboard/back/internal/pb"
	"google.golang.org/grpc"
)

func main() {
	userServiceTests()
}

func userServiceTests() {
	conn, err := grpc.Dial(":4242", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	clnt := pb.NewUserSvcClient(conn)

	r, err := clnt.AddRole(
		context.Background(),
		&pb.AddRoleReq{
			Name: "testRole",
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response from server: %s\n", r)
}
