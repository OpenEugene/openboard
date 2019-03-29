package helosvc

import (
	"context"

	"github.com/champagneabuelo/openboard/back/internal/pb"
)

var _ pb.HelloServer = &HeloSvc{}

// HeloSvc encapsulates dependencies and data required to implement the
// pb.HelloServer interface.
type HeloSvc struct{}

// New returns a pointer to a HeloSvc instance or an error.
func New() (*HeloSvc, error) {
	return &HeloSvc{}, nil
}

// AddHello implements part of the pb.HelloServer interface.
func (s *HeloSvc) AddHello(ctx context.Context, req *pb.AddHelloReq) (*pb.HelloResp, error) {
	return nil, nil
}

// OvrHello implements part of the pb.HelloServer interface.
func (s *HeloSvc) OvrHello(ctx context.Context, req *pb.OvrHelloReq) (*pb.HelloResp, error) {
	return nil, nil
}

// RmvHello implements part of the pb.HelloServer interface.
func (s *HeloSvc) RmvHello(ctx context.Context, req *pb.RmvHelloReq) (*pb.RmvHelloResp, error) {
	return nil, nil
}

// FndHellos implements part of the pb.HelloServer interface.
func (s *HeloSvc) FndHellos(ctx context.Context, req *pb.FndHellosReq) (*pb.HellosResp, error) {
	return nil, nil
}
