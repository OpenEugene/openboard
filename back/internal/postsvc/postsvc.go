package postsvc

import (
	"context"

	"github.com/OpenEugene/openboard/back/internal/pb"
	"google.golang.org/grpc"
)

var _ pb.PostServer = &PostSvc{}

// var _ grpcsrv.Registerable = &PostSvc{}

type relDB interface {
	pb.PostServer
}

// PostSvc encapsulates dependencies and data required to implement the
// pb.PostServer interface.
type PostSvc struct {
	db relDB
}

// New returns a pointer to a PostSvc instance or an error.
func New() (*PostSvc, error) {
	return &PostSvc{}, nil
}

// AddType implements part of the pb.PostServer interface.
func (s *PostSvc) AddType(ctx context.Context, req *pb.AddTypeReq) (*pb.TypeResp, error) {
	return s.db.AddType(ctx, req)
}

// AddPost implements part of the pb.PostServer interface.
func (s *PostSvc) AddPost(ctx context.Context, req *pb.AddPostReq) (*pb.PostResp, error) {
	return s.db.AddPost(ctx, req)
}

// FndPosts implements part of the pb.PostServer interface.
func (s *PostSvc) FndPosts(ctx context.Context, req *pb.FndPostsReq) (*pb.PostsResp, error) {
	return s.db.FndPosts(ctx, req)
}

// OvrPost implements part of the pb.PostServer interface.
func (s *PostSvc) OvrPost(ctx context.Context, req *pb.OvrPostReq) (*pb.PostResp, error) {
	return s.db.OvrPost(ctx, req)
}

// RmvPost implements part of the pb.PostServer interface.
func (s *PostSvc) RmvPost(ctx context.Context, req *pb.RmvPostReq) (*pb.RmvPostResp, error) {
	return s.db.RmvPost(ctx, req)
}

// RegisterWithGRPCServer implements the grpcsrv.Registerable interface.
func (s *PostSvc) RegisterWithGRPCServer(g *grpc.Server) error {
	pb.RegisterPostServer(g, s)
	return nil
}
