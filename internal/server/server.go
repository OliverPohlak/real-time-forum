package server

import (
	"log"
	"net/http"

	"github.com/mrsih/gorouter"
	"github.com/mrsih/real-time-forum/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(conf *config.Conf, router *gorouter.Router) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + conf.API.Port,
			Handler: router,
		},
	}
}

func (s *Server) Run() error {
	log.Printf("API server is online at %v", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
