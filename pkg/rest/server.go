package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var _ http.Handler = (*Server)(nil)

type Server struct {
	eng *gin.Engine
}

func NewServer() *Server {
	eng := gin.Default()
	s := &Server{
		eng: gin.Default(),
	}
	s.AddEndpoints(eng)
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.eng.ServeHTTP(w, r)
}
