package postdb

import (
	"context"
	"database/sql"

	"github.com/OpenEugene/openboard/back/internal/pb"
	"github.com/codemodus/sqlo"
	"github.com/codemodus/uidgen"
)

var _ pb.PostServer = &PostDB{}

// PostDB encapsulates dependencies and data required to implement the
// pb.PostServer interface.
type PostDB struct {
	db  *sqlo.SQLO
	drv string
	ug  *uidgen.UIDGen
}

// New returns a pointer to a PostDB instance or an error.
func New(relDB *sql.DB, driver string, offset uint64) (*PostDB, error) {
	db := PostDB{
		db:  sqlo.New(relDB),
		drv: driver,
		ug:  uidgen.New(offset, uidgen.VARCHAR26),
	}

	return &db, nil
}

// AddType implements part of the pb.PostServer interface.
func (s *PostDB) AddType(ctx context.Context, req *pb.AddTypeReq) (*pb.TypeResp, error) {
	r := &pb.TypeResp{}
	if err := s.upsertType(ctx, "", req, r); err != nil {
		return nil, err
	}
	return r, nil
}

// FndTypes implements part of the pb.PostServer interface.
func (s *PostDB) FndTypes(ctx context.Context, req *pb.FndTypesReq) (*pb.TypesResp, error) {
	r := &pb.TypesResp{}
	if err := s.findTypes(ctx, req, r); err != nil {
		return nil, err
	}
	return r, nil
}

// AddPost implements part of the pb.PostServer interface.
func (s *PostDB) AddPost(ctx context.Context, req *pb.AddPostReq) (*pb.PostResp, error) {
	r := &pb.PostResp{}
	if err := s.upsertPost(ctx, "", req, r); err != nil {
		return nil, err
	}
	return r, nil
}

// FndPosts implements part of the pb.PostServer interface.
func (s *PostDB) FndPosts(ctx context.Context, req *pb.FndPostsReq) (*pb.PostsResp, error) {
	r := &pb.PostsResp{}
	if err := s.findPosts(ctx, req, r); err != nil {
		return nil, err
	}
	return r, nil
}

// OvrPost implements part of the pb.PostServer interface.
func (s *PostDB) OvrPost(ctx context.Context, req *pb.OvrPostReq) (*pb.PostResp, error) {
	r := &pb.PostResp{}
	if err := s.upsertPost(ctx, string(req.Id), req.Req, r); err != nil {
		return nil, err
	}
	return r, nil
}

// RmvPost implements part of the pb.PostServer interface.
func (s *PostDB) RmvPost(ctx context.Context, req *pb.RmvPostReq) (*pb.RmvPostResp, error) {
	if err := s.deletePost(ctx, string(req.Id)); err != nil {
		return nil, err
	}
	return &pb.RmvPostResp{}, nil
}
