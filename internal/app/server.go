package app

import (
	"context"
	"mweather-go/config"
	"mweather-go/internal/pkg/middleware"
	"net/http"
	"time"

	"gitlab.com/makeblock-go/log"

	"github.com/gin-gonic/gin"
)

const (
	timeoutInSeconds          = 30
	maxWaitTimeBeforeShutdown = 10
)

var srv *http.Server

//RunServer run server with port
func RunServer(port string) {
	setGinMode()
	router := gin.Default()
	router.Use(middleware.Cors())
	registerRouters(router)
	startServer(router, port)
}

func ShutdownServer() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*maxWaitTimeBeforeShutdown)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Print("shutdown server: ", err)
	}
}

func setGinMode() {
	env := config.Items().ProjectEnv
	if env == config.Dev {
		gin.SetMode(gin.DebugMode)
	} else if env == config.Test {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func startServer(router *gin.Engine, port string) {
	// Listen and serve on 0.0.0.0:8080
	// router.Run(":80") 这样写就可以了，下面所有代码（go1.8+）是为了优雅处理重启等动作。可根据实际情况选择。
	srv = &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  timeoutInSeconds * time.Second,
		WriteTimeout: timeoutInSeconds * time.Second,
	}

	go func() {
		log.Println("Start Http Server ", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.FatalE("Failed to serve: ", err)
		}
	}()
}
