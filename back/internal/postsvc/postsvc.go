package postsvc

import (
	"context"
	"database/sql"

	"github.com/OpenEugene/openboard/back/internal/pb"
	"github.com/OpenEugene/openboard/back/internal/postsvc/internal/postdb"
	"github.com/OpenEugene/openboard/back/internal/postsvc/internal/postdb/mysqlmig"
	"google.golang.org/grpc"
)

var _ pb.PostServer = &PostSvc{}

// var _ grpcsrv.Registerable = &PostSvc{}
// var _ sqlmig.DataProvider = &PostSvc{}
// var _ sqlmig.Regularizer = &PostSvc{}

type relDB interface {
	pb.PostServer
}

// PostSvc encapsulates dependencies and data required to implement the
// pb.PostServer interface.
type PostSvc struct {
	db relDB
}

// New returns a pointer to a PostSvc instance or an error.
func New(relDb *sql.DB, driver string) (*PostSvc, error) {
	db, err := postdb.New(relDB, driver)

	if err != nil {
		return nil, err
	}

	s := PostSvc {
		db: db
	}

	return &s, nil
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

// MigrationData ...
func (s *PostSvc) MigrationData() (string, map[string][]byte) {
	name := "postsvc"
	m := make(map[string][]byte) 
	
	ids, err := mysqlmig.AssetDir("")
	if err != nil {
		return name, nil
	}

	for _, id := range ids {
		d, err := mysqlmig.Asset(id)
		if err != nil {
			return name, nil
		}

		m[id] = d
	}

	return name, m
}

// Regularize ...
func (s *PostSvc) Regularize (ctx context.Context) error {
	return nil
}