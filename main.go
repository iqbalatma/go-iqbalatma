package main

import (
	"fmt"
	"iqbalatma/go-iqbalatma/app/enum"
	"iqbalatma/go-iqbalatma/app/model"
	"iqbalatma/go-iqbalatma/cmd"
	"iqbalatma/go-iqbalatma/config"
	"iqbalatma/go-iqbalatma/middleware"
	iqbalatma_go_jwt_authentication "iqbalatma/go-iqbalatma/packages/iqbalatma-go-jwt-authentication"
	"iqbalatma/go-iqbalatma/route"
	"iqbalatma/go-iqbalatma/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	config.LoadLogger()

	iqbalatma_go_jwt_authentication.LoadJWTConfig()

	uuid := "6d5b857e-2832-477e-a0b8-7f3a289ef942"
	hash, err := utils.MakeHash(uuid)
	if err != nil {
		return
	}
	fmt.Println("APAKAH VALID")
	fmt.Println(utils.CheckHash(hash, uuid))
	fmt.Println(utils.CheckHash("$2a$10$37xyIWIBLsmSYKxuyPJp7e2Q5DvE0ndc4iuRpWS8L5YDtiG4snESy", uuid))

	var user model.User
	config.DB.Where("email = ?", "iqbalatma@gmail.com").First(&user)
	jwtTokenString, atv, err := iqbalatma_go_jwt_authentication.Encode(&user, iqbalatma_go_jwt_authentication.ACCESS_TOKEN, true, "", "")
	jwtTokenStringRefresh, _, err := iqbalatma_go_jwt_authentication.Encode(&user, iqbalatma_go_jwt_authentication.REFRESH_TOKEN, true, "", "")
	if err != nil {
		return
	}

	fmt.Println("USER ID : " + user.ID.String())
	fmt.Println("ATV : " + atv)
	fmt.Println("JWT TOKEN ACCESS : " + jwtTokenString)
	fmt.Println("JWT TOKEN REFRESH : " + jwtTokenStringRefresh)
	//return

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
