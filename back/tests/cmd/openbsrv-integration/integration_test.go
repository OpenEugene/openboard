package integrationTests

import (
	"context"
	"fmt"
	"github.com/OpenEugene/openboard/back/internal/pb"
	"google.golang.org/grpc"
	"testing"
)

func TestClientServices(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.Dial(":4242", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	// User Service Tests
	userClnt := pb.NewUserSvcClient(conn)
	t.Run("Add and find user role", userSvcAddAndFndRoleFunc(ctx, conn, userClnt))
	t.Run("Add and find users", userSvcAddAndFndUsersFunc(ctx, conn, userClnt))
	t.Run("Find and delete user", userSvcDelUserFunc(ctx, conn, userClnt))

	// Post Service Tests
	postClnt := pb.NewPostClient(conn)
	t.Run("Add type", postSvcAddAndFndTypesFunc(ctx, conn, postClnt))
	t.Run("Add and find posts", postSvcAddAndFndPostsFunc(ctx, conn, postClnt))
	t.Run("Add, find and edit post", postSvcEdtPostFunc(ctx, conn, postClnt))
	t.Run("Find and delete post", postSvcDelPostFunc(ctx, conn, postClnt))
}

func userSvcAddAndFndRoleFunc(ctx context.Context, conn *grpc.ClientConn, clnt pb.UserSvcClient) func(*testing.T) {
	return func(t *testing.T) {
		want := "testRole"

		_, err := clnt.AddRole(ctx, &pb.AddRoleReq{Name: want})
		if err != nil {
			t.Fatal(err)
		}

		r, err := clnt.FndRoles(ctx, &pb.FndRolesReq{
			RoleIds:   []string{},
			RoleNames: []string{want},
			Limit:     100,
			Lapse:     0,
		})

		if len(r.Items) != 1 {
			t.Fatalf("got: no items, want: %s", want)
		}

		if got := r.Items[0].Name; got != want {
			t.Fatalf("got: %v, want: %s", got, want)
		}
	}
}

type user struct {
	Username    string
	Email       string
	EmailHold   bool
	Altmail     string
	AltmailHold bool
	FullName    string
	Avatar      string
}

func userSvcAddAndFndUsersFunc(ctx context.Context, conn *grpc.ClientConn, clnt pb.UserSvcClient) func(*testing.T) {
	return func(t *testing.T) {
		tests := []struct {
			addReq *pb.AddUserReq
			want   user
			fndReq pb.FndUsersReq
		}{
			{
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
				user{
					Username:    "test username A",
					Email:       "user_a@email.com",
					EmailHold:   false,
					Altmail:     "",
					AltmailHold: false,
					FullName:    "test user full name A",
					Avatar:      "test user avatar A",
				},
				pb.FndUsersReq{
					RoleIds:     []string{},
					Email:       "user_a@email.com",
					EmailHold:   false,
					Altmail:     "",
					AltmailHold: false,
					Limit:       100,
					Lapse:       0,
				},
			},
			{
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
				user{
					Username:    "test username B",
					Email:       "userB@email.com",
					EmailHold:   false,
					Altmail:     "",
					AltmailHold: false,
					FullName:    "test user full name B",
					Avatar:      "test user avatar B",
				},
				pb.FndUsersReq{
					RoleIds:     []string{},
					Email:       "userB@email.com",
					EmailHold:   false,
					Altmail:     "",
					AltmailHold: false,
					Limit:       100,
					Lapse:       0,
				},
			},
		}

		for _, tc := range tests {
			r, err := clnt.AddUser(ctx, tc.addReq)
			if err != nil {
				t.Fatal(err)
			}

			got := user{
				Username:    r.Item.Username,
				Email:       r.Item.Email,
				EmailHold:   r.Item.EmailHold,
				Altmail:     r.Item.Altmail,
				AltmailHold: r.Item.AltmailHold,
				FullName:    r.Item.FullName,
				Avatar:      r.Item.Avatar,
			}
			if got != tc.want {
				t.Fatalf("got: %v, want: %v", got, tc.want)
			}

			userID, err := userSvcFndUser(ctx, conn, clnt, tc.fndReq)
			if err != nil {
				t.Errorf("unable to find user: %v", err)
			}
			if r.Item.Id != userID {
				t.Fatalf("got: %s, want: %s", userID, r.Item.Id)
			}
		}
	}
}

func userSvcFndUser(ctx context.Context, conn *grpc.ClientConn, clnt pb.UserSvcClient, req pb.FndUsersReq) (string, error) {
	r, err := clnt.FndUsers(ctx, &req)
	if err != nil {
		return "", fmt.Errorf("unable to find user: %w", err)
	}
	if len(r.Items) > 0 {
		return r.Items[0].Id, nil
	}
	return "", nil
}

func userSvcDelUserFunc(ctx context.Context, conn *grpc.ClientConn, clnt pb.UserSvcClient) func(*testing.T) {
	return func(t *testing.T) {
		req := pb.FndUsersReq{
			RoleIds:     []string{},
			Email:       "user_a@email.com",
			EmailHold:   false,
			Altmail:     "",
			AltmailHold: false,
			Limit:       100,
			Lapse:       0,
		}

		userID, err := userSvcFndUser(ctx, conn, clnt, req)
		if err != nil {
			t.Error(err)
		}

		if userID == "" {
			t.Fatal("unable to find userID")
		}

		_, err = clnt.RmvUser(ctx, &pb.RmvUserReq{Id: userID})
		if err != nil {
			t.Error(err)
		}

		userID, err = userSvcFndUser(ctx, conn, clnt, req)
		if err != nil {
			t.Error(err)
		}
		if userID != "" {
			t.Fatalf("expected userID to be empty string, got: %s", userID)
			t.Fail()
		}
	}
}

func postSvcAddAndFndTypesFunc(ctx context.Context, conn *grpc.ClientConn, clnt pb.PostClient) func(*testing.T) {
	return func(t *testing.T) {
		ctx := context.Background()

		tests := []struct {
			wantType   string
			addTypeReq *pb.AddTypeReq
			fndTypeReq *pb.FndTypesReq
			wantCount  int
		}{
			{
				"testTypeA",
				&pb.AddTypeReq{Name: "testTypeA"},
				&pb.FndTypesReq{Limit: 100, Lapse: 0},
				1,
			},
			{
				"testTypeB",
				&pb.AddTypeReq{Name: "testTypeB"},
				&pb.FndTypesReq{Limit: 100, Lapse: 0},
				2,
			},
		}

		for _, tc := range tests {
			r1, err := clnt.AddType(ctx, tc.addTypeReq)
			if err != nil {
				t.Error(err)
			}
			if r1.Name != tc.wantType {
				t.Fatalf("want: %s, got: %s", tc.wantType, r1.Name)
			}

			r2, err := clnt.FndTypes(ctx, tc.fndTypeReq)
			if err != nil {
				t.Error(err)
			}
			if len(r2.Items) != tc.wantCount {
				t.Errorf("want %d items, got %d", tc.wantCount, len(r2.Items))
			}
			if r2.Items[tc.wantCount-1].Name != tc.wantType {
				t.Fatalf("got: %s, want: %s", r2.Items[tc.wantCount-1], tc.wantType)
			}
		}
	}
}

type post struct {
	title  string
	body   string
	typeId string
}

func postSvcAddAndFndPostsFunc(ctx context.Context, conn *grpc.ClientConn, clnt pb.PostClient) func(*testing.T) {
	return func(t *testing.T) {
		tests := []struct {
			addReq *pb.AddPostReq
			want   post
			fndReq *pb.FndPostsReq
		}{
			{
				&pb.AddPostReq{
					Title:  "test title postA first",
					Body:   "test body of first post",
					TypeId: "2",
				},
				post{
					"test title postA first",
					"test body of first post",
					"2",
				},
				&pb.FndPostsReq{Keywords: []string{
					"postA",
					"first",
					"multiple",
					"keywords",
					"not",
					"available",
					"yet",
				},
				},
			},
			{
				&pb.AddPostReq{
					Title:  "test title postB second",
					Body:   "test body of second postB",
					TypeId: "3",
				},
				post{
					"test title postB second",
					"test body of second postB",
					"3",
				},
				&pb.FndPostsReq{Keywords: []string{"postB"}},
			},
		}

		for _, tc := range tests {
			r, err := clnt.AddPost(ctx, tc.addReq)
			if err != nil {
				t.Fatal(err)
			}

			got := post{
				title:  r.Title,
				body:   r.Body,
				typeId: r.TypeId,
			}
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}

			postID, err := postSvcFndPost(ctx, conn, clnt, tc.fndReq)
			if err != nil {
				t.Error(err)
			}

			if r.Id != postID {
				t.Fatalf("got: %s, want: %s", postID, r.Id)
			}
		}
	}
}

func postSvcFndPost(ctx context.Context, conn *grpc.ClientConn, clnt pb.PostClient, req *pb.FndPostsReq) (string, error) {
	postID, err := clnt.FndPosts(ctx, req)
	if err != nil {
		return "", fmt.Errorf("Unable to find post: %w", err)
	}

	if len(postID.Posts) > 0 {
		return postID.Posts[0].Id, nil
	}
	return "", nil
}

func postSvcEdtPostFunc(ctx context.Context, conn *grpc.ClientConn, clnt pb.PostClient) func(*testing.T) {
	return func(t *testing.T) {
		addReq := &pb.AddPostReq{
			Title:  "Post D",
			Body:   "This is post D.",
			TypeId: "2",
		}
		_, err := clnt.AddPost(ctx, addReq)
		if err != nil {
			t.Fatal(err)
		}

		fndReq := &pb.FndPostsReq{Keywords: []string{"post D"}}
		postID, err := postSvcFndPost(ctx, conn, clnt, fndReq)
		if err != nil {
			t.Error(err)
		}

		editReq := &pb.AddPostReq{
			Title:  "Post D (edited)",
			Body:   "This is post D after edits.",
			TypeId: "2",
		}

		ovrReq := &pb.OvrPostReq{Id: postID, Req: editReq}
		r, err := clnt.OvrPost(ctx, ovrReq)
		if err != nil {
			t.Error(err)
		}

		want := post{
			editReq.Title,
			editReq.Body,
			editReq.TypeId,
		}

		got := post{
			title:  r.Title,
			body:   r.Body,
			typeId: r.TypeId,
		}
		if got != want {
			t.Fatalf("got: %v, want: %v", got, want)
		}
	}
}

func postSvcDelPostFunc(ctx context.Context, conn *grpc.ClientConn, clnt pb.PostClient) func(*testing.T) {
	return func(t *testing.T) {
		req := &pb.FndPostsReq{Keywords: []string{"postC"}}
		postID, err := postSvcFndPost(ctx, conn, clnt, req)
		if err != nil {
			t.Fatal(err)
		}

		_, err = clnt.RmvPost(ctx, &pb.RmvPostReq{Id: postID})
		if err != nil {
			t.Error(err)
		}

		postID, err = postSvcFndPost(ctx, conn, clnt, req)
		if err != nil {
			t.Error(err)
		}

		if postID != "" {
			t.Fatalf("Expected userID to be empty string, got: %s", postID)
		}
	}
}
