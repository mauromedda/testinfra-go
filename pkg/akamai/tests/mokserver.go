package mockserver

import (
	"context"
	"net/http"
	"time"
)

type mockServer struct {
	// TODO: add auth mock

	server http.Server
}

func (s *mockServer) GenericHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Pragma") != "" {
		rw.Header().Set("X-Cache-Key", "S/L/8888/666666/3h/www.mockorig.com/it/donna")
		rw.Header().Set("X-Check-Cacheable", "YES")
	}
	return
}

var server *mockServer

// Run mock server
func Run() error {
	initServer()
	return server.server.ListenAndServe()
}

func initServer() {
	server = &mockServer{}
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(server.GenericHandler))
	server.server.Handler = mux
	server.server.Addr = ":8080"
}

// Close mock server
func Close() error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()

	return server.server.Shutdown(ctx)
}
