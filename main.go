package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"iqbalatma/go-iqbalatma/app/enum"
	"iqbalatma/go-iqbalatma/cmd"
	"iqbalatma/go-iqbalatma/config"
	"iqbalatma/go-iqbalatma/middleware"
	"iqbalatma/go-iqbalatma/route"
	"os"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

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
	fmt.Println("Running server")
	router := gin.Default()
	router.Use(middleware.ErrorHandler())
	route.RegisterRoute(router)

	err := router.Run(":" + config.AppConfig.AppPort)
	if err != nil {
		return
	}
}
