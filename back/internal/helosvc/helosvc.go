package helosvc

import (
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

/*type HelloServer interface {
	AddHello(context.Context, *AddHelloReq) (*HelloResp, error)
	OvrHello(context.Context, *OvrHelloReq) (*HelloResp, error)
	RmvHello(context.Context, *RmvHelloReq) (*RmvHelloResp, error)
	FndHellos(context.Context, *FndHellosReq) (*HellosResp, error)
}*/
