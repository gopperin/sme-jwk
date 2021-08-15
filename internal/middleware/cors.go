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
		AllowOrigins:     myconfig.Case.CORS.AllowOrigins,
		AllowMethods:     myconfig.Case.CORS.AllowMethods,
		AllowHeaders:     myconfig.Case.CORS.AllowHeaders,
		AllowCredentials: myconfig.Case.CORS.AllowCredentials,
		MaxAge:           time.Second * time.Duration(myconfig.Case.CORS.MaxAge),
	})
}
