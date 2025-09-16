package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func NewServer() *Server {
	log.Println("INFO NewServer: creating new http server")
	router := gin.Default()

	server := &Server{
		Router: router,
	}

	log.Println("INFO NewServer: http server created")
	return server
}

func (s *Server) Start(address string) error {
	log.Printf("INFO Server.Start: starting server at %s", address)
	return s.Router.Run(address)
}
