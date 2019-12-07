package main

import (
	"context"
	"fmt"
	"github.com/OpenEugene/openboard/back/internal/pb"
	"google.golang.org/grpc"
)

func main() {
	userServiceTests()
	postServiceTests()
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

	fmt.Printf("Response from user service add role: %s\n", r)

	r2, err := clnt.AddUser(
		context.Background(),
		&pb.AddUserReq{
			Username:    "test username a",
			Email:       "test user email a",
			EmailHold:   false,
			Altmail:     "",
			AltmailHold: false,
			FullName:    "test user full name a",
			Avatar:      "test user avatar a",
			Password:    "test user password a",
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response from user service add user: %s\n", r2.Item)

	r3, err := clnt.FndUsers(
		context.Background(),
		&pb.FndUsersReq{
			RoleIds:     []string{},
			Email:       "test user email a",
			EmailHold:   false,
			Altmail:     "",
			AltmailHold: false,
			Limit:       100,
			Lapse:       0,
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response from user service find user a: %s\n", r3)
}

func postServiceTests() {
	conn, err := grpc.Dial(":4242", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	clnt := pb.NewPostClient(conn)

	r, err := clnt.AddPost(
		context.Background(),
		&pb.AddPostReq{
			Title:  "test title",
			Body:   "test body",
			TypeId: "2",
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response from post service: %s\n", r)
}
