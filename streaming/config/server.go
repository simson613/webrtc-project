package config

import (
	"os"
)

type ServerInterface interface {
	Env() string
	Port() string
}

type Server struct {
	env  string
	port string
}

func initServerConfig() *Server {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}

	return &Server{
		env:  env,
		port: port,
	}
}

func (server *Server) Env() string {
	return server.env
}

func (server *Server) Port() string {
	return server.port
}
