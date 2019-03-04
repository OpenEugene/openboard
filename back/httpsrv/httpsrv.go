package httpsrv

import (
	"context"
	"fmt"
	"net/http"

	"github.com/champagneabuelo/openboard/back/pb"
	"github.com/codemodus/chain/v2"
	"github.com/codemodus/hedrs"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// HTTPSrv ...
type HTTPSrv struct {
	*http.Server
	gmux *runtime.ServeMux
}

// New ...
func New(origins []string) (*HTTPSrv, error) {
	gmux := runtime.NewServeMux() // TODO: set options
	mux := multiplexer(gmux, origins)

	s := HTTPSrv{
		Server: &http.Server{
			Handler: mux,
		},
		gmux: gmux,
	}

	return &s, nil
}

func multiplexer(gmux *runtime.ServeMux, origins []string) http.Handler {
	origins = append(hedrs.DefaultOrigins, origins...)
	corsOrigins := hedrs.CORSOrigins(hedrs.NewAllowed(origins...))
	corsMethods := hedrs.CORSMethods(hedrs.NewValues(hedrs.AllMethods...))
	corsHeaders := hedrs.CORSHeaders(hedrs.NewValues(hedrs.DefaultHeaders...))

	cmn := chain.New(
		corsOrigins,
		corsMethods,
		corsHeaders,
	)

	m := http.NewServeMux()

	m.Handle("/", gmux)
	m.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello, world")
	})

	return cmn.End(m)
}

// Serve ...
func (s *HTTPSrv) Serve(rpcPort, httpPort string) error {
	opts := []grpc.DialOption{grpc.WithInsecure()}

	conn, err := grpc.Dial(rpcPort, opts...)
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterAuthHandler(ctx, s.gmux, conn)
	if err != nil {
		return err
	}

	s.Server.Addr = httpPort

	if err = s.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

// Stop ...
func (s *HTTPSrv) Stop() error {
	// TODO: setup shutdown context
	return s.Server.Shutdown(context.Background())
}
