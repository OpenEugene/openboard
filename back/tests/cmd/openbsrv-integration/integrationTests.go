package integrationTests

import (
	"context"
	"fmt"
	"github.com/OpenEugene/openboard/back/internal/pb"
	"google.golang.org/grpc"
	"testing"
)

func TestClientServices(t *testing.T) {
	conn, err := grpc.Dial(":4242", grpc.WithInsecure())
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	defer conn.Close()

	t.Log("=====================Start User Service Tests=====================")
	userClnt := pb.NewUserSvcClient(conn)
	t.Run("Add user roles", userSvcAddRolesFunc(conn, userClnt))
	t.Run("Add users", userSvcAddUsersFunc(conn, userClnt))
	t.Run("Find users", userSvcFndUsersFunc(conn, userClnt))
	t.Run("Find and delete user", userSvcDelUserFunc(conn, userClnt))

	t.Log("=====================Start Post Service Tests=====================")
	postClnt := pb.NewPostClient(conn)
	t.Run("Add types", postSvcAddTypesFunc(conn, postClnt))
	t.Run("Find types", postSvcFndTypesFunc(conn, postClnt))
	t.Run("Add posts", postSvcAddPostsFunc(conn, postClnt))
	t.Run("Find posts", postSvcFndPostsFunc(conn, postClnt))
	t.Run("Find, edit, and find post", postSvcEdtPostFunc(conn, postClnt))
	t.Run("Find and post", postSvcDelPostFunc(conn, postClnt))

	// TODO
	// when finding users and posts, need to re-write include find parameters
	// to the find, so a specific thing can be found.
	//
	// test clean up
}

func userSvcAddRolesFunc(conn *grpc.ClientConn, clnt pb.UserSvcClient) func(*testing.T) {
	return func(t *testing.T) {
		r1, err := clnt.AddRole(
			context.Background(),
			&pb.AddRoleReq{
				Name: "testRole 1",
			},
		)
		if err != nil {
			t.Log(err)
			t.Fail()
		}

		r2, err := clnt.AddRole(
			context.Background(),
			&pb.AddRoleReq{
				Name: "testRole 2",
			},
		)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		t.Logf("Response from user service added roles:\n%s\n%s\n\n", r1, r2)
	}
}

func userSvcAddUsersFunc(conn *grpc.ClientConn, clnt pb.UserSvcClient) func(*testing.T) {
	return func(t *testing.T) {
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
			t.Log(err)
			t.Fail()
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
			t.Log(err)
			t.Fail()
		}

		t.Logf("Response from user service add user A and user B:\n%s\n%s\n\n", r1.Item, r2.Item)
	}
}

func userSvcFndUsersFunc(conn *grpc.ClientConn, clnt pb.UserSvcClient) func(*testing.T) {
	return func(t *testing.T) {
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
			t.Log(err)
			t.Fail()
		}

		t.Logf("Response from user service find one existing user b:\n%s\n\n", r1)

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
			t.Log(err)
			t.Fail()
		}

		t.Logf("Response from user service find no such user c:\n%s\n\n", r2)
	}
}

func userSvcFndUser(conn *grpc.ClientConn, clnt pb.UserSvcClient) (string, error) {
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
		return "", fmt.Errorf("unable to find user: %w", err)
	}

	if len(r.Items) > 0 {
		return r.Items[0].Id, nil
	}
	return "", nil
}

func userSvcDelUserFunc(conn *grpc.ClientConn, clnt pb.UserSvcClient) func(*testing.T) {
	return func(t *testing.T) {
		userID, err := userSvcFndUser(conn, clnt)
		if err != nil {
			t.Log(err)
			t.Fail()
		}

		r, err := clnt.RmvUser(
			context.Background(),
			&pb.RmvUserReq{
				Id: userID,
			},
		)
		if err != nil {
			t.Log(err)
			t.Fail()
		}

		t.Logf("Response from user service delete user A:\n%s\n\n", r)

		userID, err = userSvcFndUser(conn, clnt)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		if userID != "" {
			t.Logf("Expected userID to be empty string, got: %s", userID)
		}
	}
}

func postSvcAddTypesFunc(conn *grpc.ClientConn, clnt pb.PostClient) func(*testing.T) {
	return func(t *testing.T) {
		r1, err := clnt.AddType(
			context.Background(),
			&pb.AddTypeReq{
				Name: "testTypeA",
			},
		)
		if err != nil {
			t.Log(err)
			t.Fail()
		}

		t.Logf("Response from post service add typeA:\n%s\n\n", r1)

		r2, err := clnt.AddType(
			context.Background(),
			&pb.AddTypeReq{
				Name: "testTypeB",
			},
		)
		if err != nil {
			t.Log(err)
			t.Fail()
		}

		t.Logf("Response from post service add typeB:\n%s\n\n", r2)
	}
}

func postSvcFndTypesFunc(conn *grpc.ClientConn, clnt pb.PostClient) func(*testing.T) {
	return func(t *testing.T) {
		r, err := clnt.FndTypes(
			context.Background(),
			&pb.FndTypesReq{Limit: 100, Lapse: 0},
		)
		if err != nil {
			t.Log(err)
			t.Fail()
		}

		t.Logf("Response from post service find types:\n%s\n\n", r)
	}
}

func postSvcAddPostsFunc(conn *grpc.ClientConn, clnt pb.PostClient) func(*testing.T) {
	return func(t *testing.T) {
		r1, err := clnt.AddPost(
			context.Background(),
			&pb.AddPostReq{
				Title:  "test title post first",
				Body:   "test body of first post",
				TypeId: "2",
			},
		)
		if err != nil {
			t.Log(err)
			t.Fail()
		}

		t.Logf("Response from post service add post:\n%s\n\n", r1)

		r2, err := clnt.AddPost(
			context.Background(),
			&pb.AddPostReq{
				Title:  "test title postB second",
				Body:   "test body of second postB",
				TypeId: "3",
			},
		)
		if err != nil {
			t.Log(err)
			t.Fail()
		}

		t.Logf("Response from post service add postB:\n%s\n\n", r2)
	}
}

func postSvcFndPostsFunc(conn *grpc.ClientConn, clnt pb.PostClient) func(*testing.T) {
	return func(t *testing.T) {
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
			t.Log(err)
			t.Fail()
		}

		t.Logf("Response from post service find postB:\n%s\n\n", r1)

		r2, err := clnt.FndPosts(
			context.Background(),
			&pb.FndPostsReq{
				Keywords: []string{"PostC"},
			},
		)
		if err != nil {
			t.Log(err)
			t.Fail()
		}

		t.Logf("Response from post service find non-existant PostC:\n%s\n\n", r2)
	}
}

func postSvcFndPost(conn *grpc.ClientConn, clnt pb.PostClient) (string, error) {
	postID, err := clnt.FndPosts(
		context.Background(),
		&pb.FndPostsReq{Keywords: []string{"post"}})
	if err != nil {
		return "", fmt.Errorf("Unable to find post: %w", err)
	}

	if len(postID.Posts) > 0 {
		return postID.Posts[0].Id, nil
	}
	return "", nil
}

func postSvcEdtPostFunc(conn *grpc.ClientConn, clnt pb.PostClient) func(*testing.T) {
	return func(t *testing.T) {
		postID, err := postSvcFndPost(conn, clnt)
		if err != nil {
			t.Log(err)
			t.Fail()
		}

		addPostReq := &pb.AddPostReq{
			Title:  "Post Edited",
			Body:   "This post has been edited from the original.",
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
			t.Log(err)
			t.Fail()
		}

		postID, err = postSvcFndPost(conn, clnt)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		// TODO Check that the correct postID was found
		t.Logf("Response from post service edit post:\n%s\n\n", r)
	}
}

func postSvcDelPostFunc(conn *grpc.ClientConn, clnt pb.PostClient) func(*testing.T) {
	return func(t *testing.T) {
		postID, err := postSvcFndPost(conn, clnt)
		if err != nil {
			t.Log(err)
			t.Fail()
		}

		r, err := clnt.RmvPost(
			context.Background(),
			&pb.RmvPostReq{Id: postID},
		)
		if err != nil {
			t.Log(err)
			t.Fail()
		}

		t.Logf("Response from post service delete post:\n%s\n\n", r)

		postID, err = postSvcFndPost(conn, clnt)
		if err != nil {
			t.Log(err)
			t.Fail()
		}

		if postID != "" {
			t.Logf("Expected userID to be empty string, got: %s", postID)
		}
	}
}
