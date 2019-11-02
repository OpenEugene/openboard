package main

import (
	"context"
	"fmt"
	"github.com/OpenEugene/openboard/back/internal/pb"
	"github.com/codemodus/relay"
	"google.golang.org/grpc"
)

func main() {
	userServiceTests()
	//postServiceTests()
}

func userServiceTests() {
	check, ff := relay.New().Fns()
	defer func() { ff(recover()) }()

	cc, err := grpc.Dial(":4242", grpc.WithInsecure())
	check(err)
	defer cc.Close()

	c := pb.NewUserSvcClient(cc)

	r1, err := c.AddRole(
		context.Background(),
		&pb.AddRoleReq{
			Name: "testRole",
		},
	)
	check(err)
	fmt.Printf("Response from user service add role: %s\n", r1)

	r2, err := c.AddUser(
		context.Background(),
		&pb.AddUserReq{
			Username:    "test user name a",
			Email:       "test user email a",
			EmailHold:   false,
			Altmail:     "",
			AltmailHold: false,
			FullName:    "test user full name a",
			Avatar:      "test user avatar a",
			Password:    "test user password a",
		},
	)
	check(err)
	fmt.Printf("Response from user service add user: %s\n", r2)
}

func postServiceTests() {
	check, ff := relay.New().Fns()
	defer func() { ff(recover()) }()

	cc, err := grpc.Dial(":4242", grpc.WithInsecure())
	check(err)
	defer cc.Close()

	c := pb.NewPostClient(cc)

	r, err := c.AddPost(
		context.Background(),
		&pb.AddPostReq{
			Title:  "test title",
			Body:   "test body",
			TypeId: 4,
		},
	)
	check(err)

	fmt.Printf("Response from post service: %s\n", r)
}
