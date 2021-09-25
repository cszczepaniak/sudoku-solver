package rest

import "github.com/gin-gonic/gin"

func (s *Server) AddEndpoints(eng *gin.Engine) {
	api := s.eng.Group(`/api`)
	api.POST(`/solve`, s.solve)
}
