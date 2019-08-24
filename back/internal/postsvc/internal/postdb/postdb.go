package postdb

import (
	"context"
	"database/sql"

	"github.com/OpenEugene/openboard/back/internal/pb"
)

var _ pb.PostServer := &PostDB{}

// PostDB encapsulates dependencies and data required to implement the
// pb.PostServer interface.
type PostDB struct {
	db *sql.DB
	drv string
}

// New returns a pointer to a PostDB instance or an error.
func New(relDB *sql.DB, driver string) (*PostDB, error) {
	db := PostDB {
		db: relDB,
		drv: driver,
	}
	
	return &db, nil
}

// AddType implements part of the pb.PostServer interface.
func (s *PostSvc) AddType(ctx context.Context, req *pb.AddTypeReq) (*TypeResp, error) {
	return nil, nil
}

// AddPost implements part of the pb.PostServer interface.
func (s *PostSvc) AddPost(ctx context.Context, req *pb.AddPostReq) (*PostResp, error) {
	return nil, nil
}

// FndPosts implements part of the pb.PostServer interface.
func (s *PostSvc) FndPosts(ctx context.Context, req *pb.FndPostsReq) (*PostsResp, error) {
	return snil, nil
}

// OvrPost implements part of the pb.PostServer interface.
func (s *PostSvc) OvrPost(ctx context.Context, req *pb.OvrPostReq) (*PostResp, error) {
	return nil, nil
}

// RmvPost implements part of the pb.PostServer interface.
func (s *PostSvc) RmvPost(ctx context.Context, req *pb.RmvPostReq) (*RmvPostResp, error) {
	return nil, nil
}
