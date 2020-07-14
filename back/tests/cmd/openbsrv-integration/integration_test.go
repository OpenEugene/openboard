package integrationTests

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/OpenEugene/openboard/back/internal/pb"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
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
	t.Run("Add and find user role", userSvcAddAndFndRoleFn(ctx, conn, userClnt))
	t.Run("Add and find users", userSvcAddAndFndUsersFn(ctx, conn, userClnt))
	t.Run("Find and delete user", userSvcDelUserFn(ctx, conn, userClnt))

	// Post Service Tests
	postClnt := pb.NewPostClient(conn)
	t.Run("Add type", postSvcAddAndFndTypesFn(ctx, conn, postClnt))
	t.Run("Add and find posts", postSvcAddAndFndPostsFn(ctx, conn, postClnt))
	t.Run("Add, find and edit post", postSvcEdtPostFn(ctx, conn, postClnt))
	t.Run("Find and delete post", postSvcDelPostFn(ctx, conn, postClnt))
}

func userSvcAddAndFndRoleFn(ctx context.Context, conn *grpc.ClientConn, clnt pb.UserSvcClient) func(*testing.T) {
	return func(t *testing.T) {
		tests := []struct {
			want   string
			addReq *pb.AddRoleReq
			fndReq *pb.FndRolesReq
		}{
			{
				"user",
				&pb.AddRoleReq{Name: "user"},
				&pb.FndRolesReq{
					RoleIds:   []string{},
					RoleNames: []string{"user"},
					Limit:     100,
					Lapse:     0,
				},
			},
			{
				"admin",
				&pb.AddRoleReq{Name: "admin"},
				&pb.FndRolesReq{
					RoleIds:   []string{},
					RoleNames: []string{"admin"},
					Limit:     100,
					Lapse:     0,
				},
			},
		}

		for _, tt := range tests {
			_, err := clnt.AddRole(ctx, tt.addReq)
			if err != nil {
				t.Fatal(err)
			}

			r, err := clnt.FndRoles(ctx, tt.fndReq)

			if len(r.Items) != 1 {
				t.Fatalf("got: no items, want: %s", tt.want)
			}

			if got := r.Items[0].Name; got != tt.want {
				t.Fatalf("got: %v, want: %s", got, tt.want)
			}
		}

	}
}

func roleID(ctx context.Context, conn *grpc.ClientConn, clnt pb.UserSvcClient, req *pb.FndRolesReq) (string, error) {
	r, err := clnt.FndRoles(ctx, req)
	if err != nil {
		return "", err
	}

	if len(r.Items) == 0 {
		return "", nil
	}

	return r.Items[0].Id, nil
}

func userSvcAddAndFndUsersFn(ctx context.Context, conn *grpc.ClientConn, clnt pb.UserSvcClient) func(*testing.T) {
	return func(t *testing.T) {
		fndAdminRoleIdReq := &pb.FndRolesReq{
			RoleIds:   []string{},
			RoleNames: []string{"admin"},
			Limit:     100,
			Lapse:     0,
		}

		fndUserRoleIdReq := &pb.FndRolesReq{
			RoleIds:   []string{},
			RoleNames: []string{"user"},
			Limit:     100,
			Lapse:     0,
		}

		adminRoleID, err := roleID(ctx, conn, clnt, fndAdminRoleIdReq)
		if err != nil {
			t.Fatalf("retrieving role ID from role name: %v", err)
		}

		userRoleID, err := roleID(ctx, conn, clnt, fndUserRoleIdReq)
		if err != nil {
			t.Fatalf("retrieving role ID from role name: %v", err)
		}

		tests := []struct {
			desc       string
			addUserReq *pb.AddUserReq
			want       *pb.User
			fndUserReq pb.FndUsersReq
		}{
			{
				"Add and find user A",
				&pb.AddUserReq{
					Username:    "test username A",
					Email:       "user_a@email.com",
					EmailHold:   false,
					Altmail:     "",
					AltmailHold: false,
					FullName:    "test user full name A",
					Avatar:      "test user avatar A",
					Password:    "test user password A",
					RoleIds:     []string{userRoleID},
				},
				&pb.User{
					Username:    "test username A",
					Email:       "user_a@email.com",
					EmailHold:   false,
					Altmail:     "",
					AltmailHold: false,
					FullName:    "test user full name A",
					Avatar:      "test user avatar A",
					Roles:       []*pb.RoleResp{{Name: "user"}},
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
				"Add and find user B",
				&pb.AddUserReq{
					Username:    "test username B",
					Email:       "userB@email.com",
					EmailHold:   false,
					Altmail:     "",
					AltmailHold: false,
					FullName:    "test user full name B",
					Avatar:      "test user avatar B",
					Password:    "test user password B",
					RoleIds:     []string{userRoleID, adminRoleID},
				},
				&pb.User{
					Username:    "test username B",
					Email:       "userB@email.com",
					EmailHold:   false,
					Altmail:     "",
					AltmailHold: false,
					FullName:    "test user full name B",
					Avatar:      "test user avatar B",
					Roles: []*pb.RoleResp{
						{Name: "user"},
						{Name: "admin"},
					},
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
			t.Log(tc.desc)
			r, err := clnt.AddUser(ctx, tc.addUserReq)
			if err != nil {
				t.Fatal(err)
			}

			got := r.Item
			wantUserID := got.Id

			// Unset fields that aren't being tested.
			got.Id = ""
			got.XXX_unrecognized = []byte{}
			got.XXX_sizecache = 0

			for i := 0; i < len(got.Roles); i++ {
				got.Roles[i].Id = ""
				got.Roles[i].XXX_unrecognized = []byte{}
				got.Roles[i].XXX_sizecache = 0
			}

			if reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got: %v, want: %v", got, tc.want)
			}

			gotUserID, err := userSvcFndUser(ctx, conn, clnt, tc.fndUserReq)
			if err != nil {
				t.Errorf("unable to find user: %v", err)
			}
			if gotUserID != wantUserID {
				t.Fatalf("got: %s, want: %s", gotUserID, wantUserID)
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

func userSvcDelUserFn(ctx context.Context, conn *grpc.ClientConn, clnt pb.UserSvcClient) func(*testing.T) {
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
			t.Fatalf("unable to find user %s", userID)
		}

		_, err = clnt.RmvUser(ctx, &pb.RmvUserReq{Id: userID})
		if err != nil {
			t.Error(err)
		}

		resp, err := clnt.FndUsers(ctx, &req)
		if err != nil {
			t.Error(err)
		}

		if resp.Items[0].Deleted == nil {
			t.Fatalf("expected user %s deleted_at to not be nil", resp.Items[0].Id)
			t.Fail()
		}
	}
}

func postSvcAddAndFndTypesFn(ctx context.Context, conn *grpc.ClientConn, clnt pb.PostClient) func(*testing.T) {
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

			gotType := r1.Name
			if gotType != tc.wantType {
				t.Fatalf("got: %s, want: %s", gotType, tc.wantType)
			}

			r2, err := clnt.FndTypes(ctx, tc.fndTypeReq)
			if err != nil {
				t.Error(err)
			}

			gotCount := len(r2.Items)
			if gotCount != tc.wantCount {
				t.Errorf("got %d items, want %d", gotCount, tc.wantCount)
			}
			gotType = r2.Items[tc.wantCount-1].Name
			if gotType != tc.wantType {
				t.Fatalf("got: %s, want: %s", gotType, tc.wantType)
			}
		}
	}
}

func postSvcAddAndFndPostsFn(ctx context.Context, conn *grpc.ClientConn, clnt pb.PostClient) func(*testing.T) {
	return func(t *testing.T) {
		tests := []struct {
			addReq *pb.AddPostReq
			want   *pb.PostResp
			fndReq *pb.FndPostsReq
		}{
			{
				&pb.AddPostReq{
					Title:  "test title postA first",
					Body:   "test body of first post",
					TypeId: "2",
				},
				&pb.PostResp{
					Title:  "test title postA first",
					Body:   "test body of first post",
					TypeId: "2",
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
				&pb.PostResp{
					Title:  "test title postB second",
					Body:   "test body of second postB",
					TypeId: "3",
				},
				&pb.FndPostsReq{Keywords: []string{"postB"}},
			},
		}

		for _, tc := range tests {
			r, err := clnt.AddPost(ctx, tc.addReq)
			if err != nil {
				t.Fatal(err)
			}

			got := r
			wantPostID := got.Id

			// Unset fields that aren't being tested.
			got.Id = ""
			got.Slug = ""
			got.Created = nil
			got.Updated = nil
			got.Deleted = nil
			got.Blocked = nil
			got.XXX_unrecognized = []byte{}
			got.XXX_sizecache = 0

			if reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}

			gotPostID, err := postSvcFndPost(ctx, conn, clnt, tc.fndReq)
			if err != nil {
				t.Error(err)
			}

			t.Logf("got: %s, want: %s", gotPostID, wantPostID)
			if gotPostID != wantPostID {
				t.Fatalf("got: %s, want: %s", gotPostID, wantPostID)
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

func postSvcEdtPostFn(ctx context.Context, conn *grpc.ClientConn, clnt pb.PostClient) func(*testing.T) {
	return func(t *testing.T) {
		addReq := &pb.AddPostReq{
			Title:  "Post C",
			Body:   "This is post C.",
			TypeId: "2",
		}
		_, err := clnt.AddPost(ctx, addReq)
		if err != nil {
			t.Fatal(err)
		}

		fndReq := &pb.FndPostsReq{Keywords: []string{"post C"}}
		postID, err := postSvcFndPost(ctx, conn, clnt, fndReq)
		if err != nil {
			t.Error(err)
		}

		editReq := &pb.AddPostReq{
			Title:  "Post C (edited)",
			Body:   "This is post C after edits.",
			TypeId: "2",
		}

		ovrReq := &pb.OvrPostReq{Id: postID, Req: editReq}
		r, err := clnt.OvrPost(ctx, ovrReq)
		if err != nil {
			t.Error(err)
		}

		want := &pb.PostResp{
			Title:  editReq.Title,
			Body:   editReq.Body,
			TypeId: editReq.TypeId,
		}

		got := r

		// Unset fields that aren't being tested.
		got.Id = ""
		got.Slug = ""
		got.Created = &timestamp.Timestamp{}
		got.Updated = &timestamp.Timestamp{}
		got.Deleted = &timestamp.Timestamp{}
		got.Blocked = &timestamp.Timestamp{}
		got.XXX_unrecognized = []byte{}
		got.XXX_sizecache = 0

		if reflect.DeepEqual(got, want) {
			t.Fatalf("got: %v, want: %v", got, want)
		}
	}
}

func postSvcDelPostFn(ctx context.Context, conn *grpc.ClientConn, clnt pb.PostClient) func(*testing.T) {
	return func(t *testing.T) {
		addReq := &pb.AddPostReq{
			Title:  "test title postD first",
			Body:   "test body of fourth post",
			TypeId: "4",
		}

		_, err := clnt.AddPost(ctx, addReq)
		if err != nil {
			t.Fatal(err)
		}

		fndReq := &pb.FndPostsReq{Keywords: []string{"postD"}}

		fndResp, err := clnt.FndPosts(ctx, fndReq)
		if err != nil {
			t.Error(err)
		}
		t.Log(fndResp)
		if fndResp.Posts[0].Deleted != nil {
			t.Fatalf(
				"expected post %s deleted_at to be nil, got %v",
				fndResp.Posts[0].Id,
				fndResp.Posts[0].Deleted,
			)
			t.Fail()
		}

		_, err = clnt.RmvPost(ctx, &pb.RmvPostReq{Id: fndResp.Posts[0].Id})
		if err != nil {
			t.Error(err)
		}

		fndResp, err = clnt.FndPosts(ctx, fndReq)
		if err != nil {
			t.Error(err)
		}

		if fndResp.Posts[0].Deleted == nil {
			t.Fatalf("expected post %s deleted_at to not be nil", fndResp.Posts[0].Id)
			t.Fail()
		}
	}
}
