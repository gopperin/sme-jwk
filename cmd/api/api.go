package api

import (
	"net/http"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	myconfig "sme-jwk/internal/config"
	mymiddleware "sme-jwk/internal/middleware"
	myrouter "sme-jwk/internal/router"
)

// StartCmd api
var (
	StartCmd = &cobra.Command{
		Use:   "start",
		Short: "start sme-jwk api",
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			//启动API服务
			run()

			logrus.Println("sme-jwk end")
		},
	}
)

func setup() {

	//1. 读取配置
	myconfig.Setup("./")

}

func run() {
	router := gin.Default()

	router.Use(mymiddleware.Cors())

	_cacheChain, _ := InitChainCache()

	/* api base */
	myrouter.SetupBaseRouter(router)

	/* jwk */
	myrouter.SetupJwkRouter(router, _cacheChain)

	server := &http.Server{
		Addr:         ":" + myconfig.Case.Application.Port,
		Handler:      router,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
	}

	logrus.Println("sme-stage start on:", myconfig.Case.Application.Port)
	gracehttp.Serve(server)
}
