package helodb

import (
	"context"
	"database/sql"

	"github.com/champagneabuelo/openboard/back/internal/pb"
)

var _ pb.HelloServer = &HeloDB{}

// HeloDB encapsulates dependencies and data required to implement the
// pb.HelloServer interface.
type HeloDB struct {
	db  *sql.DB
	drv string
}

// New returns a pointer to a HeloDB instance or an error.
func New(relDB *sql.DB, driver string) (*HeloDB, error) {
	db := HeloDB{
		db:  relDB,
		drv: driver,
	}

	return &db, nil
}

// AddHello implements part of the pb.HelloServer interface.
func (s *HeloDB) AddHello(ctx context.Context, req *pb.AddHelloReq) (*pb.HelloResp, error) {
	return nil, nil
}

// OvrHello implements part of the pb.HelloServer interface.
func (s *HeloDB) OvrHello(ctx context.Context, req *pb.OvrHelloReq) (*pb.HelloResp, error) {
	return nil, nil
}

// RmvHello implements part of the pb.HelloServer interface.
func (s *HeloDB) RmvHello(ctx context.Context, req *pb.RmvHelloReq) (*pb.RmvHelloResp, error) {
	return nil, nil
}

// FndHellos implements part of the pb.HelloServer interface.
func (s *HeloDB) FndHellos(ctx context.Context, req *pb.FndHellosReq) (*pb.HellosResp, error) {
	return nil, nil
}
