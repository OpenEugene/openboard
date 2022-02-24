package postsvc

import (
	"context"
	"database/sql"
	"io/fs"
	"io/ioutil"

	"github.com/OpenEugene/openboard/back/internal/pb"
	"github.com/OpenEugene/openboard/back/internal/postsvc/internal/asset"
	"github.com/OpenEugene/openboard/back/internal/postsvc/internal/postdb"
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
func New(relDb *sql.DB, driver string, offset uint64) (*PostSvc, error) {
	db, err := postdb.New(relDb, driver, offset)
	if err != nil {
		return nil, err
	}

	s := PostSvc{
		db: db,
	}

	return &s, nil
}

// AddType implements part of the pb.PostServer interface.
func (s *PostSvc) AddType(ctx context.Context, req *pb.AddTypeReq) (*pb.TypeResp, error) {
	return s.db.AddType(ctx, req)
}

// FndTypes implements part of the pb.PostServer interface.
func (s *PostSvc) FndTypes(ctx context.Context, req *pb.FndTypesReq) (*pb.TypesResp, error) {
	return s.db.FndTypes(ctx, req)
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

	afs, err := asset.NewFS()
	if err != nil {
		return name, nil
	}

	sqls, err := fs.Glob(afs, "*.sql")
	if err != nil {
		return name, nil
	}

	for _, sql := range sqls {
		f, err := afs.Open(sql)
		if err != nil {
			return name, nil
		}

		m[sql], err = ioutil.ReadAll(f)
		if err != nil {
			return name, nil
		}
	}

	return name, m
}

// Regularize ...
func (s *PostSvc) Regularize(ctx context.Context) error {
	return nil
}
