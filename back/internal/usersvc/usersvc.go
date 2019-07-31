package usersvc

import (
	"context"
	"database/sql"

	"github.com/OpenEugene/openboard/back/internal/pb"
	"github.com/OpenEugene/openboard/back/internal/usersvc/internal/userdb"
	"google.golang.org/grpc"
)

var _ pb.UserServer = &UserSvc{}

//var _ sqlmig.DataProvider = &UserSvc{}
//var _ sqlmig.Regularizer = &UserSvc{}

// UserSvc encapsulates dependencies and data required to implement the
// pb.UserServer interface.
type UserSvc struct {
	// TODO: implement UserSvc
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
	pb.RegisterUserServer(g, s)

	return nil
}

// AddRole implements part of the pb.UserServer interface.
func (s *UserSvc) AddRole(ctx context.Context, req *pb.AddRoleReq) (*pb.RoleResp, error) {
	// TODO: implement AddRole

	return nil, nil
}

// FndRoles implements part of the pb.UserServer interface.
func (s *UserSvc) FndRoles(ctx context.Context, req *pb.FndRolesReq) (*pb.RolesResp, error) {
	// TODO: implement FndRoles

	return nil, nil
}

// AddUser implements part of the pb.UserServer interface.
func (s *UserSvc) AddUser(ctx context.Context, req *pb.AddUserReq) (*pb.UserResp, error) {
	// TODO: implement AddUser

	return nil, nil
}

// OvrUser implements part of the pb.UserServer interface.
func (s *UserSvc) OvrUser(ctx context.Context, req *pb.OvrUserReq) (*pb.UserResp, error) {
	// TODO: implement OvrUser

	return nil, nil
}

// FndUsers implements part of the pb.UserServer interface.
func (s *UserSvc) FndUsers(ctx context.Context, req *pb.FndUsersReq) (*pb.UsersResp, error) {
	// TODO: implement FndUsers

	return nil, nil
}

// RmvUser implements part of the pb.UserServer interface.
func (s *UserSvc) RmvUser(ctx context.Context, req *pb.RmvUserReq) (*pb.RmvUserResp, error) {
	// TODO: implement RmvUser

	return nil, nil
}

// MigrationData ...
func (s *UserSvc) MigrationData() (string, map[string][]byte) {
	name := "usersvc"
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

//Regularize ...
func (s *UserSvc) Regularize(ctx context.Context) error {
	return nil
}
