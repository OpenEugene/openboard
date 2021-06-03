package httpsrv

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/OpenEugene/openboard/back/internal/httpsrv/internal/embedded"
	"github.com/OpenEugene/openboard/back/internal/pb"
	"github.com/codemodus/chain/v2"
	"github.com/codemodus/hedrs"
	"github.com/codemodus/mixmux"
	"github.com/codemodus/swagui"
	"github.com/codemodus/swagui/suidata3"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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
	mux, err := multiplexer(time.Now(), gmux, origins)
	if err != nil {
		return nil, err
	}

	s := HTTPSrv{
		Server: &http.Server{
			Handler: mux,
		},
		gmux: gmux,
	}

	return &s, nil
}

func multiplexer(start time.Time, gmux *runtime.ServeMux, origins []string) (http.Handler, error) {
	handleSwagger, err := swaggerHandler(start, "/v", "apidocs.swagger.json")
	if err != nil {
		return nil, err
	}

	m := mixmux.NewTreeMux(nil)

	gm := http.StripPrefix("/v", gmux)
	m.Any("/v/", gm)
	m.Any("/v/*x", gm)

	m.Any("/v/docs/swagger.json", handleSwagger)

	if ui, err := swagui.New(http.NotFoundHandler(), suidata3.New()); err == nil {
		sh := http.StripPrefix("/v/docs", ui.Handler("/v/docs/swagger.json"))
		m.Get("/v/docs/", sh)
		m.Get("/v/docs/*x", sh)
	}

	origins = append(hedrs.DefaultOrigins, origins...)
	corsOrigins := hedrs.CORSOrigins(hedrs.NewAllowed(origins...))
	corsMethods := hedrs.CORSMethods(hedrs.NewValues(hedrs.AllMethods...))
	corsHeaders := hedrs.CORSHeaders(hedrs.NewValues(hedrs.DefaultHeaders...))

	cmn := chain.New(
		corsOrigins,
		corsMethods,
		corsHeaders,
	)

	return cmn.End(m), nil
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

	err = pb.RegisterPostHandler(ctx, s.gmux, conn)
	if err != nil {
		return err
	}

	err = pb.RegisterUserSvcHandler(ctx, s.gmux, conn)
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

func swaggerHandler(start time.Time, basePath, name string) (http.Handler, error) {
	fs, err := embedded.NewFS()
	if err != nil {
		return nil, err
	}

	f, err := fs.Open("apidocs.swagger.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	d, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	if basePath != "" {
		if d, err = setSwaggerBasePath(basePath, d); err != nil {
			return nil, err
		}
	}

	mt := start
	i, err := f.Stat()
	if err == nil {
		mt = i.ModTime()
	}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b := bytes.NewReader(d)
		http.ServeContent(w, r, name, mt, b)
	})

	return h, nil
}

func setSwaggerBasePath(path string, d []byte) ([]byte, error) {
	var j swaggerJSON
	if err := json.Unmarshal(d, &j); err != nil {
		return nil, err
	}

	j.BasePath = path

	return json.Marshal(&j)
}

type swaggerJSON struct {
	Swagger string `json:"swagger"`
	/*Info struct {
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		Version     string `json:"version,omitempty"`
	}*/
	Info        json.RawMessage `json:"info,omitempty"`
	BasePath    string          `json:"basePath,omitempty"`
	Schemes     json.RawMessage `json:"schemes,omitempty"`  // []string
	Consumes    json.RawMessage `json:"consumes,omitempty"` // []string
	Produces    json.RawMessage `json:"produces,omitempty"` // []string
	Paths       json.RawMessage `json:"paths,omitempty"`
	Definitions json.RawMessage `json:"definitions,omitempty"`
}
