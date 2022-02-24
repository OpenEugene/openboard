package usersvc

import (
	"context"
	"database/sql"
	"io/fs"
	"io/ioutil"

	"github.com/OpenEugene/openboard/back/internal/pb"
	"github.com/OpenEugene/openboard/back/internal/usersvc/internal/asset"
	"github.com/OpenEugene/openboard/back/internal/usersvc/internal/userdb"
	"google.golang.org/grpc"
)

var _ pb.UserSvcServer = &UserSvc{}

//var _ grpcsrv.Registerable = &UserSvc{}
//var _ sqlmig.DataProvider = &UserSvc{}
//var _ sqlmig.Regularizer = &UserSvc{}

type relDb interface {
	pb.UserSvcServer
}

// UserSvc encapsulates dependencies and data required to implement the
// pb.UserServer interface.
type UserSvc struct {
	db relDb
}

// New returns a pointer to a UserSvc instance or an error.
func New(relDb *sql.DB, driver string, offset uint64) (*UserSvc, error) {
	db, err := userdb.New(relDb, driver, offset)
	if err != nil {
		return nil, err
	}

	s := UserSvc{
		db: db,
	}

	return &s, nil
}

// RegisterWithGRPCServer implements the grpcsrv.Registerable interface.
func (s *UserSvc) RegisterWithGRPCServer(g *grpc.Server) error {
	pb.RegisterUserSvcServer(g, s)

	return nil
}

// AddRole implements part of the pb.UserServer interface.
func (s *UserSvc) AddRole(ctx context.Context, req *pb.AddRoleReq) (*pb.RoleResp, error) {
	return s.db.AddRole(ctx, req)
}

// FndRoles implements part of the pb.UserServer interface.
func (s *UserSvc) FndRoles(ctx context.Context, req *pb.FndRolesReq) (*pb.RolesResp, error) {
	return s.db.FndRoles(ctx, req)
}

// AddUser implements part of the pb.UserServer interface.
func (s *UserSvc) AddUser(ctx context.Context, req *pb.AddUserReq) (*pb.UserResp, error) {
	return s.db.AddUser(ctx, req)
}

// OvrUser implements part of the pb.UserServer interface.
func (s *UserSvc) OvrUser(ctx context.Context, req *pb.OvrUserReq) (*pb.UserResp, error) {
	return s.db.OvrUser(ctx, req)
}

// FndUsers implements part of the pb.UserServer interface.
func (s *UserSvc) FndUsers(ctx context.Context, req *pb.FndUsersReq) (*pb.UsersResp, error) {
	return s.db.FndUsers(ctx, req)
}

// RmvUser implements part of the pb.UserServer interface.
func (s *UserSvc) RmvUser(ctx context.Context, req *pb.RmvUserReq) (*pb.RmvUserResp, error) {
	return s.db.RmvUser(ctx, req)
}

// MigrationData ...
func (s *UserSvc) MigrationData() (string, map[string][]byte) {
	name := "usersvc"
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

//Regularize ...
func (s *UserSvc) Regularize(ctx context.Context) error {
	return nil
}
