package TaskManager

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	Server *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.Server = &http.Server{
		Addr:           port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
	}

	return s.Server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
