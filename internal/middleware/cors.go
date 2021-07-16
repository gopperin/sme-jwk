package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	myconfig "sme-jwk/internal/config"
)

// Cors 跨域请求中间件
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     myconfig.CORS.AllowOrigins,
		AllowMethods:     myconfig.CORS.AllowMethods,
		AllowHeaders:     myconfig.CORS.AllowHeaders,
		AllowCredentials: myconfig.CORS.AllowCredentials,
		MaxAge:           time.Second * time.Duration(myconfig.CORS.MaxAge),
	})
}
