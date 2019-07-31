package userdb

import (
	"context"
	"database/sql"

	"github.com/OpenEugene/openboard/back/internal/pb"
)

var _ pb.UserServer = &UserDB{}

// UserDB encapsulates dependencies and data required to implement the
// pb.UserServer interface.
type UserDB struct {
	db  *sqlo.SQLO
	drv string
}

// New returns a pointer to a UserDB instance or an error.
func New(relDB *sql.DB, driver string, offset uint64) (*UserDB, error) {
	db := UserDB{
		db:  sqlo.New(relDB),
		drv: driver,
		ug:  uidgen.New(offset, uidgen.VARCHAR26),
	}

	return &db, nil
}

// AddUser implements part of the pb.UserServer interface.
func (s *UserDB) AddUser(ctx context.Context, req *pb.AddUserReq) (*pb.UserResp, error) {
	return nil, nil
}

// OvrUser implements part of the pb.UserServer interface.
func (s *UserDB) OvrUser(ctx context.Context, req *pb.OvrUserReq) (*.pb.UserResp, error) {
	return nil, nil
}

// RmvUser implements part of the pb.UserServer interface.
func (s *UserDB) RmvUser(ctx context.Context, req *pb.RmvUserReq) (*pb.RmvUserResp, error) {
	return nil, nil
}

// FndUsers implements part of the pb.UserServer interface.
func (s *UserDB) FndUsers(ctx context.Context, req *pb.FndUsersReq) (*pb.UsersResp, error) {
	return nil, nil
}

// AddRole implements part of the pb.UserServer interface.
func (s *UserDB) AddRole(ctx context.Context, req *pb.AddRoleReq) (*pb.RoleResp, error) {
	return nil, nil
}

// FndRoles implements part of the pb.UserServer interface.
func (s *UserDB) FndRoles(ctx context.Context, req *pb.FndRolesReq) (*pb.RolesResp, error) {
	return nil, nil
}