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

	// Do one test with user service
	uc := pb.NewUserSvcClient(cc)

	ur, err := uc.AddRole(
		context.Background(),
		&pb.AddRoleReq{
			Name: "testRole",
		},
	)
	check(err)

	fmt.Printf("Response from user service: %s\n", ur)

	// Do one test with post service
	pc := pb.NewPostClient(cc)

	pr, err := pc.AddPost(
		context.Background(),
		&pb.AddPostReq{
			Title:  "test title",
			Body:   "test body",
			TypeId: 4,
		},
	)
	check(err)

	fmt.Printf("Response from post service: %s\n", pr)
}
