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
	r1, err := clnt.AddRole(
		context.Background(),
		&pb.AddRoleReq{
			Name: "testRole 1",
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	r2, err := clnt.AddRole(
		context.Background(),
		&pb.AddRoleReq{
			Name: "testRole 2",
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response from user service added roles:\n%s\n%s\n\n", r1, r2)
}

func userSvcAddUsers(conn *grpc.ClientConn, clnt pb.UserSvcClient) {
	r1, err := clnt.AddUser(
		context.Background(),
		&pb.AddUserReq{
			Username:    "test username A",
			Email:       "user_a@email.com",
			EmailHold:   false,
			Altmail:     "",
			AltmailHold: false,
			FullName:    "test user full name A",
			Avatar:      "test user avatar A",
			Password:    "test user password A",
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	r2, err := clnt.AddUser(
		context.Background(),
		&pb.AddUserReq{
			Username:    "test username B",
			Email:       "userB@email.com",
			EmailHold:   false,
			Altmail:     "",
			AltmailHold: false,
			FullName:    "test user full name B",
			Avatar:      "test user avatar B",
			Password:    "test user password B",
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response from user service add user A and user B:\n%s\n%s\n\n", r1.Item, r2.Item)
}

func userSvcFndUsers(conn *grpc.ClientConn, clnt pb.UserSvcClient) {
	r1, err := clnt.FndUsers(
		context.Background(),
		&pb.FndUsersReq{
			RoleIds:     []string{},
			Email:       "userB@email.com",
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

	r2, err := clnt.FndUsers(
		context.Background(),
		&pb.FndUsersReq{
			RoleIds:     []string{},
			Email:       "userC@email.com",
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

	fmt.Printf("Response from user service find one existing user and one not:\n%s\n%s\n\n", r1, r2)
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
