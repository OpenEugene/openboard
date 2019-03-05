package httpsrv

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/champagneabuelo/openboard/back/httpsrv/internal/embed/assets"
	"github.com/champagneabuelo/openboard/back/pb"
	"github.com/codemodus/chain/v2"
	"github.com/codemodus/hedrs"
	"github.com/codemodus/mixmux"
	"github.com/codemodus/swagui"
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
	mux := multiplexer(time.Now(), gmux, origins)

	s := HTTPSrv{
		Server: &http.Server{
			Handler: mux,
		},
		gmux: gmux,
	}

	return &s, nil
}

func multiplexer(start time.Time, gmux *runtime.ServeMux, origins []string) http.Handler {
	origins = append(hedrs.DefaultOrigins, origins...)
	corsOrigins := hedrs.CORSOrigins(hedrs.NewAllowed(origins...))
	corsMethods := hedrs.CORSMethods(hedrs.NewValues(hedrs.AllMethods...))
	corsHeaders := hedrs.CORSHeaders(hedrs.NewValues(hedrs.DefaultHeaders...))

	cmn := chain.New(
		corsOrigins,
		corsMethods,
		corsHeaders,
	)

	m := mixmux.NewTreeMux(nil)
	m.Get("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello, world")
	}))
	m.Any("/v/*x", gmux)
	// TODO: add swagger json merging
	m.Any("/v/docs/auth.swagger.json", swaggerJSONHandler(start, "auth"))
	m.Any("/v/docs/user.swagger.json", swaggerJSONHandler(start, "user"))

	if ui, err := swagui.New(nil); err == nil {
		sh := http.StripPrefix("/v/docs/", ui.Handler("/v/docs/user.swagger.json"))
		m.Get("/v/docs/", sh)
		m.Get("/v/docs/*x", sh)
	}

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

	err = pb.RegisterAuthHandler(ctx, s.gmux, conn) // TODO: []callback from new func args
	if err != nil {
		return err
	}

	err = pb.RegisterUserHandler(ctx, s.gmux, conn)
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

func swaggerJSONHandler(start time.Time, prefix string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := prefix + ".swagger.json"

		d, err := assets.Asset(name)
		if err != nil {
			stts := http.StatusInternalServerError
			http.Error(w, http.StatusText(stts), stts)
			return
		}

		mt := start
		i, err := assets.AssetInfo(name)
		if err == nil {
			mt = i.ModTime()
		}

		b := bytes.NewReader(d)
		http.ServeContent(w, r, name, mt, b)
	})
}
