package router

import (
	"github.com/eko/gocache/cache"
	"github.com/gin-gonic/gin"

	mywire "sme-jwk/internal/wire"
)

// SetupJwkRouter SetupJwkRouter
func SetupJwkRouter(g *gin.Engine, chaincache *cache.ChainCache) *gin.Engine {

	// initialize API
	_jwkAPI := mywire.InitJWKAPI(chaincache)

	r := g.Group("/jwk")
	{
		r.GET("/:kid", _jwkAPI.GetJwk)

		// Ping test
		r.POST("/", _jwkAPI.CreateJwk)

		r.POST("/signin", _jwkAPI.Signin)
		r.POST("/verify", _jwkAPI.VerifyToken)
	}

	return g
}
