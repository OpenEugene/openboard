package artifact_tests

import (
	"context"
	"fmt"
	"reflect"

	"github.com/OpenEugene/openboard/back/internal/pb"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
)

func unsetUntestedFields(item interface{}) {
	val := reflect.Indirect(reflect.ValueOf(item))
	if val.Kind() != reflect.Struct {
		return
	}

	strFldNames := []string{"Id"}
	for _, name := range strFldNames {
		fv := val.FieldByName(name)
		if fv.IsValid() && fv.Kind() == reflect.String && fv.CanSet() {
			fv.SetString("")
		}
	}

	byteFldNames := []string{"XXX_unrecognized"}
	b := new([]byte)
	bt := reflect.TypeOf(b)

	for _, name := range byteFldNames {
		fv := val.FieldByName(name)
		if fv.IsValid() && fv.Type() == bt && fv.CanSet() {
			fv.Set(reflect.Zero(bt))
		}
	}

	timeFldNames := []string{
		"LastLogin",
		"Created",
		"Updated",
		"Deleted",
		"Blocked",
	}
	t := new(timestamp.Timestamp)
	tt := reflect.TypeOf(t)

	for _, name := range timeFldNames {
		fv := val.FieldByName(name)
		if fv.IsValid() && fv.Type() == tt && fv.CanSet() {
			fv.Set(reflect.Zero(tt))
		}
	}

	intFieldNames := []string{"XXX_sizecache"}
	for _, name := range intFieldNames {
		fv := val.FieldByName(name)
		if fv.IsValid() && fv.Kind() == reflect.Int && fv.CanSet() {
			fv.SetInt(0)
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
