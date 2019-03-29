package helosvc

import (
	"context"
	"database/sql"

	"github.com/champagneabuelo/openboard/back/internal/helosvc/internal/helodb"
	"github.com/champagneabuelo/openboard/back/internal/helosvc/internal/helodb/mysqlmig"
	"github.com/champagneabuelo/openboard/back/internal/pb"
	"google.golang.org/grpc"
)

var _ pb.HelloServer = &HeloSvc{}

//var _ grpcsrv.Registerable = &HeloSvc{}
//var _ sqlmig.DataProvider = &HeloSvc{}
//var _ sqlmig.Regularizer = &HeloSvc{}

type relDB interface {
	pb.HelloServer
}

// HeloSvc encapsulates dependencies and data required to implement the
// pb.HelloServer interface.
type HeloSvc struct {
	db relDB
}

// New returns a pointer to a HeloSvc instance or an error.
func New(relDB *sql.DB, driver string) (*HeloSvc, error) {
	db, err := helodb.New(relDB, driver)
	if err != nil {
		return nil, err
	}

	s := HeloSvc{
		db: db,
	}

	return &s, nil
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

// MigrationData ...
func (s *HeloSvc) MigrationData() (string, map[string][]byte) {
	name := "helosvc"
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
func (s *HeloSvc) Regularize(ctx context.Context) error {
	return nil
}
