package artifact_tests

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/OpenEugene/openboard/back/internal/pb"
	"google.golang.org/grpc"
)

func TestPostClientServices(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.Dial(":4242", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	// Post Service Tests
	postClnt := pb.NewPostClient(conn)
	t.Run("Add type", postSvcAddAndFndTypesFn(ctx, conn, postClnt))
	t.Run("Add and find posts", postSvcAddAndFndPostsFn(ctx, conn, postClnt))
	t.Run("Add, find and edit post", postSvcEdtPostFn(ctx, conn, postClnt))
	t.Run("Find and delete post", postSvcDelPostFn(ctx, conn, postClnt))
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
			unsetUntestedFields(got)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %#v, want: %#v", got, tc.want)
			}

			gotPostID, err := postSvcFndPost(ctx, conn, clnt, tc.fndReq)
			if err != nil {
				t.Error(err)
			}

			if gotPostID != wantPostID {
				t.Fatalf("got: %s, want: %s", gotPostID, wantPostID)
			}
		}
	}
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

		unsetUntestedFields(got)

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got: %#v, want: %#v", got, want)
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
