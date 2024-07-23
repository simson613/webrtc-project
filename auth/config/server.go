package config

import (
	"os"
)

type ServerInterface interface {
	Port() string
}

type Server struct {
	port string
}

func initServerConfig() *Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5011"
	}

	return &Server{
		port: port,
	}
}

func (server *Server) Port() string {
	return server.port
}
