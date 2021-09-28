package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) AddEndpoints(eng *gin.Engine) {
	api := s.eng.Group(`/api`)
	api.POST(`/solve`, s.solve)
	api.GET(`/health`, func(c *gin.Context) {
		c.String(http.StatusOK, `Healthy!`)
	})
}
