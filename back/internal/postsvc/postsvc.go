package postsvc

import (
	"context"

	"github.com/OpenEugene/openboard/back/internal/pb"
	"google.golang.org/grpc"
)

var _ pb.PostServer = &PostSvc{}

// var _ grpcsrv.Registerable = &PostSvc{}

// PostSvc encapsulates dependencies and data required to implement the
// pb.PostServer interface.
type PostSvc struct{}

// New returns a pointer to a PostSvc instance or an error.
func New() (*PostSvc, error) {
	return &PostSvc{}, nil
}

// AddType implements part of the pb.PostServer interface.
func (s *PostSvc) AddType(ctx context.Context, req *pb.AddTypeReq) (*pb.TypeResp, error) {
	return nil, nil
}

// AddPost implements part of the pb.PostServer interface.
func (s *PostSvc) AddPost(ctx context.Context, req *pb.AddPostReq) (*pb.PostResp, error) {
	return nil, nil
}

// FndPosts implements part of the pb.PostServer interface.
func (s *PostSvc) FndPosts(ctx context.Context, req *pb.FndPostsReq) (*pb.PostsResp, error) {
	return nil, nil
}

// OvrPost implements part of the pb.PostServer interface.
func (s *PostSvc) OvrPost(ctx context.Context, req *pb.OvrPostReq) (*pb.PostResp, error) {
	return nil, nil
}

// RmvPost implements part of the pb.PostServer interface.
func (s *PostSvc) RmvPost(ctx context.Context, req *pb.RmvPostReq) (*pb.RmvPostResp, error) {
	return nil, nil
}

// RegisterWithGRPCServer implements the grpcsrv.Registerable interface.
func (s *PostSvc) RegisterWithGRPCServer(g *grpc.Server) error {
	pb.RegisterPostServer(g, s)
	return nil
}
