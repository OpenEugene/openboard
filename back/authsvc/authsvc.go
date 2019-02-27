package authsvc

import (
	"context"

	"github.com/champagneabuelo/openboard/back/pb"
	"google.golang.org/grpc"
)

var _ pb.AuthServer = &AuthSvc{}

// AuthSvc ecapsulates dependencies and data required to implement the
// pb.AuthServer interface.
type AuthSvc struct {
	// TODO: implement AuthSvc
}

// RegisterWithGRPCServer implements the grpcsrv.Registerable interface.
func (s *AuthSvc) RegisterWithGRPCServer(g *grpc.Server) error {
	pb.RegisterAuthServer(g, s)

	return nil
}

// AddAuth implements part of the pb.AuthServer interface.
func (s *AuthSvc) AddAuth(ctx context.Context, req *pb.AddAuthReq) (*pb.AuthResp, error) {
	// TODO: implement AddAuth

	return nil, nil
}

// RmvAuth implements part of the pb.AuthServer interface.
func (s *AuthSvc) RmvAuth(ctx context.Context, req *pb.RmvAuthReq) (*pb.RmvAuthResp, error) {
	// TODO: implement RmvAuth

	return nil, nil
}

// AddVoucher implements part of the pb.AuthServer interface.
func (s *AuthSvc) AddVoucher(ctx context.Context, req *pb.AddVoucherReq) (*pb.AddVoucherResp, error) {
	// TODO: implement AddVoucher

	return nil, nil
}
