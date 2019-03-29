package helodb

import (
	"context"
	"database/sql"

	"github.com/champagneabuelo/openboard/back/internal/pb"
	"github.com/codemodus/sqlo"
	"github.com/codemodus/uidgen"
)

var _ pb.HelloServer = &HeloDB{}

// HeloDB encapsulates dependencies and data required to implement the
// pb.HelloServer interface.
type HeloDB struct {
	db  *sqlo.SQLO
	drv string
	ug  *uidgen.UIDGen
}

// New returns a pointer to a HeloDB instance or an error.
func New(relDB *sql.DB, driver string, offset uint64) (*HeloDB, error) {
	db := HeloDB{
		db:  sqlo.New(relDB),
		drv: driver,
		ug:  uidgen.New(offset, uidgen.VARCHAR26),
	}

	return &db, nil
}

// AddHello implements part of the pb.HelloServer interface.
func (s *HeloDB) AddHello(ctx context.Context, req *pb.AddHelloReq) (*pb.HelloResp, error) {
	r := &pb.HelloResp{}
	if err := s.upsertHello(ctx, "", req, r); err != nil {
		return nil, err
	}
	return r, nil
}

// OvrHello implements part of the pb.HelloServer interface.
func (s *HeloDB) OvrHello(ctx context.Context, req *pb.OvrHelloReq) (*pb.HelloResp, error) {
	r := &pb.HelloResp{}
	if err := s.upsertHello(ctx, req.Id, req.Req, r); err != nil {
		return nil, err
	}
	return r, nil
}

// RmvHello implements part of the pb.HelloServer interface.
func (s *HeloDB) RmvHello(ctx context.Context, req *pb.RmvHelloReq) (*pb.RmvHelloResp, error) {
	if err := s.deleteHello(ctx, req.Id); err != nil {
		return nil, err
	}
	return &pb.RmvHelloResp{}, nil
}

// FndHellos implements part of the pb.HelloServer interface.
func (s *HeloDB) FndHellos(ctx context.Context, req *pb.FndHellosReq) (*pb.HellosResp, error) {
	r := &pb.HellosResp{}
	if err := s.findHellos(ctx, req, r); err != nil {
		return nil, err
	}
	return r, nil
}
