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
			Altmail:     "test alt email a",
			AltmailHold: false,
			FullName:    "test user full name a",
			Avatar:      "test user avatar a",
			Password:    "test user password a",
		},
	)
	check(err)
	fmt.Printf("Response from user service add user A: %s\n", r2)

	r3, err := c.AddUser(
		context.Background(),
		&pb.AddUserReq{
			Username:    "test user name b",
			Email:       "test user email b",
			EmailHold:   false,
			Altmail:     "test alt email b",
			AltmailHold: false,
			FullName:    "test user full name b",
			Avatar:      "test user avatar b",
			Password:    "test user password b",
		},
	)
	check(err)
	fmt.Printf("Response from user service add user B: %s\n", r3)

	r4, err := c.FndUsers(
		context.Background(),
		&pb.FndUsersReq{
			Email:       "test user email b",
			EmailHold:   false,
			Altmail:     "test alt email b",
			AltmailHold: false,
		},
	)
	check(err)
	fmt.Printf("Response from user service find user B: %s\n", r4)

	// Delete the user_id for user B
	r5, err := c.RmvUser(
		context.Background(),
		&pb.RmvUserReq{
			Id: r4.Items[0].Id,
		},
	)
	check(err)
	fmt.Printf("Response from user service remove user B: %s\n", r5)

	r6, err := c.FndUsers(
		context.Background(),
		&pb.FndUsersReq{
			Email:       "test user email a",
			EmailHold:   false,
			Altmail:     "",
			AltmailHold: false,
		},
	)
	check(err)
	fmt.Printf("Response from user service find user A: %s\n", r6)

	// Update existing user A username
	r7, err := c.OvrUser(
		context.Background(),
		&pb.OvrUserReq{
			Id: r6.Items[0].Id,
			Req: &pb.AddUserReq{
				Username:    "test username Aye",
				Email:       "test user email a",
				EmailHold:   false,
				Altmail:     "test user alt_mail a",
				AltmailHold: false,
				Avatar:      "test user a avatar",
			},
		},
	)
	check(err)
	fmt.Printf("Response from user service remove user B: %s\n", r7)
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
