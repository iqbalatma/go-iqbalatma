package main

import (
	"fmt"
	config2 "github.com/iqbalatma/gofortress/config"
	"iqbalatma/go-iqbalatma/app/enum"
	"iqbalatma/go-iqbalatma/cmd"
	"iqbalatma/go-iqbalatma/config"
	"iqbalatma/go-iqbalatma/middleware"
	"iqbalatma/go-iqbalatma/route"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	config.LoadLogger()
	config2.LoadJWTConfig()
	config2.ConnectRedis()

	if len(os.Args) < 2 {
		runServer()
	} else {
		switch os.Args[1] {
		case "server":
		case string(enum.CommandSeeder):
			cmd.RunningSeeder()
		default:
			fmt.Println("❌ Unknown command:", os.Args[1])
		}
	}
}

func runServer() {

	router := gin.New()
	router.
		Use(middleware.RequestLatencyMiddleware()).
		Use(middleware.ErrorHandler()).
		Use(middleware.RequestIDMiddleware())
	route.RegisterRoute(router)

	err := router.Run(":" + config.AppConfig.AppPort)
	if err != nil {
		return
	}
}
