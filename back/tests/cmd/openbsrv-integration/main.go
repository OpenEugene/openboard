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
	userAID := userSvcFndUserA(conn, userClnt)
	userSvcDelUserA(conn, userClnt, userAID)
	_ = userSvcFndUserA(conn, userClnt)

	fmt.Println("=====================Start Post Service Tests=====================")
	postClnt := pb.NewPostClient(conn)
	postSvcAddTypes(conn, postClnt)
	postSvcFndTypes(conn, postClnt)
	postSvcAddPosts(conn, postClnt)
	postSvcFndPosts(conn, postClnt)
	postID := postSvcFndPostA(conn, postClnt)
	postSvcEdtPostA(conn, postClnt, postID)
	_ = postSvcFndPostA(conn, postClnt)
	postSvcDelPostA(conn, postClnt, postID)
	_ = postSvcFndPostA(conn, postClnt)
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

	fmt.Printf("Response from user service find one existing user b:\n%s\n\n", r1)

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

	fmt.Printf("Response from user service find no such user c:\n%s\n\n", r2)
}

func userSvcFndUserA(conn *grpc.ClientConn, clnt pb.UserSvcClient) string {
	r, err := clnt.FndUsers(
		context.Background(),
		&pb.FndUsersReq{
			RoleIds:     []string{},
			Email:       "user_a@email.com",
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

	fmt.Printf("Response from user service find user A:\n%s\n\n", r)

	if len(r.Items) > 0 {
		return r.Items[0].Id
	}
	return ""
}

func userSvcDelUserA(conn *grpc.ClientConn, clnt pb.UserSvcClient, userID string) {
	r, err := clnt.RmvUser(
		context.Background(),
		&pb.RmvUserReq{
			Id: userID,
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response from user service delete user A:\n%s\n\n", r)
}

func postSvcAddTypes(conn *grpc.ClientConn, clnt pb.PostClient) {
	r1, err := clnt.AddType(
		context.Background(),
		&pb.AddTypeReq{
			Name: "testTypeA",
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response from post service add typeA:\n%s\n\n", r1)

	r2, err := clnt.AddType(
		context.Background(),
		&pb.AddTypeReq{
			Name: "testTypeB",
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response from post service add typeB:\n%s\n\n", r2)
}

func postSvcFndTypes(conn *grpc.ClientConn, clnt pb.PostClient) {
	r, err := clnt.FndTypes(
		context.Background(),
		&pb.FndTypesReq{Limit: 100, Lapse: 0},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response from post service find types:\n%s\n\n", r)
}

func postSvcAddPosts(conn *grpc.ClientConn, clnt pb.PostClient) {
	r1, err := clnt.AddPost(
		context.Background(),
		&pb.AddPostReq{
			Title:  "test title postA first",
			Body:   "test body of first postA",
			TypeId: "2",
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response from post service add postA:\n%s\n\n", r1)

	r2, err := clnt.AddPost(
		context.Background(),
		&pb.AddPostReq{
			Title:  "test title postB second",
			Body:   "test body of second postB",
			TypeId: "3",
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response from post service add postB:\n%s\n\n", r2)
}

func postSvcFndPosts(conn *grpc.ClientConn, clnt pb.PostClient) {
	r1, err := clnt.FndPosts(context.Background(),
		&pb.FndPostsReq{Keywords: []string{
			"postB",
			"second",
			"multiple",
			"keywords",
			"not",
			"available",
			"yet",
		},
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response from post service find postB:\n%s\n\n", r1)

	r2, err := clnt.FndPosts(
		context.Background(),
		&pb.FndPostsReq{
			Keywords: []string{"PostC"},
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response from post service find non-existant PostC:\n%s\n\n", r2)
}

func postSvcFndPostA(conn *grpc.ClientConn, clnt pb.PostClient) string {
	postAID, err := clnt.FndPosts(
		context.Background(),
		&pb.FndPostsReq{Keywords: []string{"postA"}})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response from post service find postA:\n%s\n\n", postAID)
	if len(postAID.Posts) > 0 {
		return postAID.Posts[0].Id
	}
	return ""
}

func postSvcEdtPostA(conn *grpc.ClientConn, clnt pb.PostClient, postID string) {
	addPostReq := &pb.AddPostReq{
		Title:  "Post A Edited",
		Body:   "This first postA has been edited from the original.",
		TypeId: "2",
	}

	r, err := clnt.OvrPost(
		context.Background(),
		&pb.OvrPostReq{
			Id:  postID,
			Req: addPostReq,
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response from post service edit postA:\n%s\n\n", r)
}

func postSvcDelPostA(conn *grpc.ClientConn, clnt pb.PostClient, postID string) {
	r, err := clnt.RmvPost(
		context.Background(),
		&pb.RmvPostReq{Id: postID},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response from post service delete postA:\n%s\n\n", r)
}
