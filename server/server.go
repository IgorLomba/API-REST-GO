package server

// By: DARTHxIKE

import (
	"log"

	"github.com/IgorLomba/API-REST-GO/API-REST-GO/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "3000",
		server: gin.Default(),
	}
}

// run server in 3000
func (s *Server) Run() {
	router := routes.LoadRoutes(s.server)
	log.Println("server is running at port", s.port)
	log.Fatal(router.Run(":" + s.port))
}
