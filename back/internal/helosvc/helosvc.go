package helosvc

import (
	"context"

	"github.com/champagneabuelo/openboard/back/internal/pb"
	"google.golang.org/grpc"
)

var _ pb.HelloServer = &HeloSvc{}

//var _ grpcsrv.Registerable = &HeloSvc{}

type relDB interface {
	pb.HelloServer
}

// HeloSvc encapsulates dependencies and data required to implement the
// pb.HelloServer interface.
type HeloSvc struct {
	db relDB
}

// New returns a pointer to a HeloSvc instance or an error.
func New() (*HeloSvc, error) {
	return &HeloSvc{}, nil
}

// AddHello implements part of the pb.HelloServer interface.
func (s *HeloSvc) AddHello(ctx context.Context, req *pb.AddHelloReq) (*pb.HelloResp, error) {
	return s.db.AddHello(ctx, req)
}

// OvrHello implements part of the pb.HelloServer interface.
func (s *HeloSvc) OvrHello(ctx context.Context, req *pb.OvrHelloReq) (*pb.HelloResp, error) {
	return s.db.OvrHello(ctx, req)
}

// RmvHello implements part of the pb.HelloServer interface.
func (s *HeloSvc) RmvHello(ctx context.Context, req *pb.RmvHelloReq) (*pb.RmvHelloResp, error) {
	return s.db.RmvHello(ctx, req)
}

// FndHellos implements part of the pb.HelloServer interface.
func (s *HeloSvc) FndHellos(ctx context.Context, req *pb.FndHellosReq) (*pb.HellosResp, error) {
	return s.db.FndHellos(ctx, req)
}

// RegisterWithGRPCServer implements the grpcsrv.Registerable interface.
func (s *HeloSvc) RegisterWithGRPCServer(g *grpc.Server) error {
	pb.RegisterHelloServer(g, s)
	return nil
}
