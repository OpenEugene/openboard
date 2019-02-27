package usersvc

import (
	"context"

	"github.com/champagneabuelo/openboard/back/pb"
	"google.golang.org/grpc"
)

var _ pb.UserServer = &UserSvc{}

// UserSvc encapsulates dependencies and data required to implement the
// pb.UserServer interface.
type UserSvc struct {
	// TODO: implement UserSvc
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
