package userdb

import (
	"context"
	"database/sql"

	"github.com/OpenEugene/openboard/back/internal/pb"
	"github.com/codemodus/uidgen"
	"github.com/jmoiron/sqlx"
)

var _ pb.UserSvcServer = &UserDB{}

// UserDB encapsulates dependencies and data required to implement the
// pb.UserServer interface.
type UserDB struct {
	db  *sqlx.DB
	drv string
	ug  *uidgen.UIDGen
}

// New returns a pointer to a UserDB instance or an error.
func New(relDB *sql.DB, driver string, offset uint64) (*UserDB, error) {
	db := UserDB{
		db:  sqlx.NewDb(relDB, driver),
		drv: driver,
		ug:  uidgen.New(offset, uidgen.VARCHAR26),
	}

	return &db, nil
}

// AddUser implements part of the pb.UserServer interface.
func (s *UserDB) AddUser(ctx context.Context, req *pb.AddUserReq) (*pb.UserResp, error) {
	r := &pb.UserResp{}
	if err := s.upsertUser(ctx, "", req, r); err != nil {
		return nil, err
	}
	return r, nil
}

// OvrUser implements part of the pb.UserServer interface.
func (s *UserDB) OvrUser(ctx context.Context, req *pb.OvrUserReq) (*pb.UserResp, error) {
	r := &pb.UserResp{}
	if err := s.upsertUser(ctx, string(req.Id), req.Req, r); err != nil {
		return nil, err
	}
	return r, nil
}

// RmvUser implements part of the pb.UserServer interface.
func (s *UserDB) RmvUser(ctx context.Context, req *pb.RmvUserReq) (*pb.RmvUserResp, error) {
	if err := s.deleteUser(ctx, string(req.Id)); err != nil {
		return nil, err
	}
	return &pb.RmvUserResp{}, nil
}

// FndUsers implements part of the pb.UserServer interface.
func (s *UserDB) FndUsers(ctx context.Context, req *pb.FndUsersReq) (*pb.UsersResp, error) {
	r := &pb.UsersResp{}
	if err := s.findUsers(ctx, req, r); err != nil {
		return nil, err
	}
	return r, nil
}

// AddRole implements part of the pb.UserServer interface.
func (s *UserDB) AddRole(ctx context.Context, req *pb.AddRoleReq) (*pb.RoleResp, error) {
	r := &pb.RoleResp{}
	if err := s.upsertRole(ctx, "", req, r); err != nil {
		return nil, err
	}
	return r, nil
}

// FndRoles implements part of the pb.UserServer interface.
func (s *UserDB) FndRoles(ctx context.Context, req *pb.FndRolesReq) (*pb.RolesResp, error) {
	r := &pb.RolesResp{}
	if err := s.findRoles(ctx, req, r); err != nil {
		return nil, err
	}
	return r, nil
}
