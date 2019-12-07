package main

import (
	"context"
	"fmt"
	"github.com/OpenEugene/openboard/back/internal/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":4242", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	userClnt := pb.NewUserSvcClient(conn)
	fmt.Println("=====================Start User Service Tests=====================")
	userSvcAddRoles(conn, userClnt)
	userSvcAddUsers(conn, userClnt)
	userSvcFndUsers(conn, userClnt)

	fmt.Println("=====================Start Post Service Tests=====================")
	postClnt := pb.NewPostClient(conn)
	postSvcAddPosts(conn, postClnt)
}

func userSvcAddRoles(conn *grpc.ClientConn, clnt pb.UserSvcClient) {
	r, err := clnt.AddRole(
		context.Background(),
		&pb.AddRoleReq{
			Name: "testRole",
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response from user service add role:\n%s\n\n", r)
}

func userSvcAddUsers(conn *grpc.ClientConn, clnt pb.UserSvcClient) {
	r, err := clnt.AddUser(
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

	fmt.Printf("Response from user service add user:\n%s\n\n", r.Item)
}

func userSvcFndUsers(conn *grpc.ClientConn, clnt pb.UserSvcClient) {
	r, err := clnt.FndUsers(
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

	fmt.Printf("Response from user service find user a:\n%s\n\n", r)
}

func postSvcAddPosts(conn *grpc.ClientConn, clnt pb.PostClient) {
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

	fmt.Printf("Response from post service:\n%s\n\n", r)
}
