package router

import (
	"github.com/gin-gonic/gin"

	mywire "sme-jwk/internal/wire"
)

// SetupBaseRouter SetupBaseRouter
func SetupBaseRouter(g *gin.Engine) {

	// initialize API
	_baseAPI := mywire.InitBaseAPI()

	r := g.Group("/")
	{
		r.GET("/favicon.ico", func(c *gin.Context) {
			return
		})

		r.GET("release", _baseAPI.Release)

		// Ping test
		r.POST("/health", _baseAPI.Health)
	}
}
