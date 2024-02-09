package server

import (
	"context"
	"net/http"
	"testTaskMedods/config"
	"testTaskMedods/pkg"
	"time"
)

type Server struct {
	httpServer http.Server
}

func NewServer(config config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: http.Server{
			Addr:           config.Server.Host + ":" + config.Server.Port,
			Handler:        handler,
			MaxHeaderBytes: 1024 * 1024,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}

}

func (s *Server) Run() error {
	pkg.InfoLog.Printf("Starting server on  %s", s.httpServer.Addr)
	if err := s.httpServer.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.httpServer.Close(); err != nil {
		return err
	}
	return nil
}
