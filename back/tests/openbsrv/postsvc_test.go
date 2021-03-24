package main_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/OpenEugene/openboard/back/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
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
	t.Run("Find all posts", postSvcFindAllPosts(ctx, conn, postClnt))
	t.Run("Add, find and edit post", postSvcEdtPostFn(ctx, conn, postClnt))
	t.Run("Find and delete post", postSvcDelPostFn(ctx, conn, postClnt))
	t.Run("Find posts by keywords", postSvcKeywordSearch(ctx, conn, postClnt))
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

			if !proto.Equal(got, tc.want) {
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

		if !proto.Equal(got, want) {
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

func postsContain(resp *pb.PostsResp, post *pb.PostResp) bool {
	for _, j := range resp.Posts {
		if j.Title == post.Title && j.Body == post.Body && j.TypeId == post.TypeId {
			return true
		}
	}
	return false
}

func postSvcFindAllPosts(ctx context.Context, conn *grpc.ClientConn, clnt pb.PostClient) func(*testing.T) {
	return func(t *testing.T) {
		tests := []struct {
			fndReq *pb.FndPostsReq
			want   *pb.PostsResp
		}{
			{
				&pb.FndPostsReq{Keywords: []string{}},
				&pb.PostsResp{
					Posts: []*pb.PostResp{
						{
							Title:  "test title postA first",
							Body:   "test body of first post",
							TypeId: "2",
						},
						{
							Title:  "test title postB second",
							Body:   "test body of second postB",
							TypeId: "3",
						},
					},
				},
			},
		}

		for _, tc := range tests {
			resp, err := clnt.FndPosts(ctx, tc.fndReq)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if len(tc.want.Posts) != len(resp.Posts) {
				t.Errorf("mismatch between response length and post length, "+
					"want: %d got: %d", len(tc.want.Posts), len(resp.Posts))
			}

			for _, post := range tc.want.Posts {
				if !postsContain(resp, post) {
					t.Errorf("couldn't find post with title: %s", post.Title)
				}
			}
		}
	}
}

func postSvcKeywordSearch(ctx context.Context, conn *grpc.ClientConn, clnt pb.PostClient) func(*testing.T) {
	return func(t *testing.T) {
		tests := []struct {
			fndReq *pb.FndPostsReq
			want   *pb.PostsResp
		}{
			{
				// Test without any results
				&pb.FndPostsReq{Keywords: []string{"randomgiberish"}},
				&pb.PostsResp{
					Posts: []*pb.PostResp{},
				},
			},
			{
				&pb.FndPostsReq{Keywords: []string{"postB", "postD"}},
				&pb.PostsResp{
					Posts: []*pb.PostResp{
						{
							Title:  "test title postB second",
							Body:   "test body of second postB",
							TypeId: "3",
						},
						{
							Title:  "test title postD first",
							Body:   "test body of fourth post",
							TypeId: "4",
						},
					},
				},
			},
			{
				&pb.FndPostsReq{Keywords: []string{"edited"}},
				&pb.PostsResp{
					Posts: []*pb.PostResp{
						{
							Title:  "Post C (edited)",
							Body:   "This is post C after edits.",
							TypeId: "2",
						},
					},
				},
			},
		}

		for _, tc := range tests {
			resp, err := clnt.FndPosts(ctx, tc.fndReq)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if len(tc.want.Posts) != len(resp.Posts) {
				t.Errorf("mismatch between response length and post length, "+
					"want: %d got: %d", len(tc.want.Posts), len(resp.Posts))
			}

			for _, post := range tc.want.Posts {
				if !postsContain(resp, post) {
					t.Errorf("couldn't find post with title: %s", post.Title)
				}
			}
		}
	}
}
