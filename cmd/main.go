package main

import (
	"mweather-go/config"
	"mweather-go/internal/app"
	"mweather-go/internal/app/weather/service/city"
	"os"
	"os/signal"
	"syscall"

	"gitlab.com/makeblock-go/log"
)

func main() {
	config.Load()
	isProduction := config.Items().ProjectEnv == config.Prod
	log.SetUp(isProduction, log.Any("app_name", config.Items().AppName))
	defer log.Sync()

	city.Register()
	// cnf := mysql.NewConfig(
	// 	config.Items().Mysql.User,
	// 	config.Items().Mysql.Pwd,
	// 	config.Items().Mysql.Host,
	// 	config.Items().Mysql.Port,
	// 	config.Items().Mysql.DBName,
	// 	config.Items().Mysql.Charset,
	// 	logger.Info)
	// mysql.Register(cnf)
	// defer mysql.Close()
	// gormCtxLogger := trace.NewGormCtxLogger(logger.Config{LogLevel: logger.Info})
	// mysql.SetLogger(mysql.GetDB(), gormCtxLogger)

	// redis.SetUp(
	// 	config.Items().Redis.Host,
	// 	config.Items().Redis.Port,
	// 	config.Items().Redis.Pwd)
	// defer redis.Close()

	app.RunServer(":8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.ShutdownServer()
	log.Println("Shutdown Server ...")
}
