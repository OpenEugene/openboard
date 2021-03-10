package main_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/OpenEugene/openboard/back/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
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
}

func userSvcAddAndFndRoleFn(ctx context.Context, conn *grpc.ClientConn, clnt pb.UserSvcClient) func(*testing.T) {
	return func(t *testing.T) {
		tests := []struct {
			desc   string
			want   string
			addReq *pb.AddRoleReq
			fndReq *pb.FndRolesReq
		}{
			{
				"add and find user role",
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
				"add and find admin role",
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
				t.Fatalf("%s: adding role %v: %v", tt.desc, tt.want, err)
			}

			r, err := clnt.FndRoles(ctx, tt.fndReq)
			if err != nil {
				t.Fatalf("%s: finding role %v: %v", tt.desc, tt.want, err)
			}

			if len(r.Items) == 0 {
				t.Fatalf("%s: got: no items, want: %s", tt.desc, tt.want)
			}

			if len(r.Items) > 1 {
				t.Fatalf("%s: got: %+v, want: %s", tt.desc, r, tt.want)
			}

			if got := r.Items[0].Name; got != tt.want {
				t.Fatalf("%s: got: %v, want: %s", tt.desc, got, tt.want)
			}
		}
	}
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
			fndUserReq *pb.FndUsersReq
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
				&pb.FndUsersReq{
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
				&pb.FndUsersReq{
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
			r, err := clnt.AddUser(ctx, tc.addUserReq)
			if err != nil {
				t.Fatal(err)
			}

			got := r.Item
			wantUserID := got.Id

			unsetUntestedFields(got)

			for i := 0; i < len(got.Roles); i++ {
				unsetUntestedFields(got.Roles[i])
			}

			if !proto.Equal(got, tc.want) {
				t.Fatalf("%s: got: %#v, want: %#v", tc.desc, got, tc.want)
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

func userSvcDelUserFn(ctx context.Context, conn *grpc.ClientConn, clnt pb.UserSvcClient) func(*testing.T) {
	return func(t *testing.T) {
		req := &pb.FndUsersReq{
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

		resp, err := clnt.FndUsers(ctx, req)
		if err != nil {
			t.Error(err)
		}

		if resp.Items[0].Deleted == nil {
			t.Fatalf("expected user %s deleted_at to not be nil", resp.Items[0].Id)
			t.Fail()
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

func userSvcFndUser(ctx context.Context, conn *grpc.ClientConn, clnt pb.UserSvcClient, req *pb.FndUsersReq) (string, error) {
	r, err := clnt.FndUsers(ctx, req)
	if err != nil {
		return "", fmt.Errorf("unable to find user: %w", err)
	}
	if len(r.Items) > 0 {
		return r.Items[0].Id, nil
	}
	return "", nil
}
