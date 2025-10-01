package main

import (
	"fmt"
	"iqbalatma/go-iqbalatma/config"
	"iqbalatma/go-iqbalatma/middleware"
	iqbalatma_go_jwt_authentication "iqbalatma/go-iqbalatma/packages/iqbalatma-go-jwt-authentication"
	"iqbalatma/go-iqbalatma/route"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	config.LoadLogger()

	iqbalatma_go_jwt_authentication.LoadJWTConfig()

	jwtTokenString, err := iqbalatma_go_jwt_authentication.Encode()
	if err != nil {
		return
	}

	payload, err := iqbalatma_go_jwt_authentication.Decoding(jwtTokenString)
	if err != nil {
		panic(err)
	}

	fmt.Println("Payload: ", payload.SUB)
	//token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
	//	return []byte(iqbalatma_go_jwt_authentication.Config.JWTSecretKey), nil
	//})
	//
	//if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	//	fmt.Println("✅ Token valid")
	//	for key, val := range claims {
	//		fmt.Printf("%s: %v\n", key, val)
	//	}
	//
	//	// contoh akses spesifik
	//	fmt.Println("UserID:", claims["user_id"])
	//	fmt.Println("ExpiresAt:", claims["exp"])
	//} else {
	//	fmt.Println("❌ Token invalid")
	//}
	//fmt.Println(jwtToken)

	//var (
	//	key []byte
	//	t   *jwt.Token
	//	s   string
	//)
	//key = []byte("8pQ1c3Hj8aX3vF-2fWc7P4lWmZ9N8dS9pV2Y5sQ4a2c")
	//t = jwt.NewWithClaims(iqbalatma_go_jwt_authentication.GetSigningMethod(),
	//	jwt.MapClaims{
	//		"iss": "my-auth-server",
	//		"sub": "john",
	//		"foo": 2,
	//	})
	//s, _ = t.SignedString(key)
	//fmt.Println(s)

	//fake := faker.New()
	//user := model.User{
	//	FirstName: fake.Person().FirstName(),
	//	LastName:  fake.Person().LastName(),
	//	Email:     fake.Internet().Email(),
	//	Password:  "Password",
	//}
	//config.DB.Create(&user)
	//fmt.Println(user.GetSubjectKey())
	//
	//encode, err := iqbalatma_go_jwt_authentication.Encode()
	//if err != nil {
	//	return
	//}
	//
	//fmt.Println(encode)
	//
	//var (
	//	key []byte
	//	t   *jwt.Token
	//	s   string
	//)
	//
	//key = []byte("8pQ1c3Hj8aX3vF-2fWc7P4lWmZ9N8dS9pV2Y5sQ4a2c")
	//t = jwt.NewWithClaims(jwt.SigningMethodHS256,
	//	jwt.MapClaims{
	//		"iss": "my-auth-server",
	//		"sub": "john",
	//		"foo": 2,
	//	})
	//s, _ = t.SignedString(key)
	//fmt.Println(s)
	//if len(os.Args) < 2 {
	//	runServer()
	//} else {
	//	switch os.Args[1] {
	//	case "server":
	//	case string(enum.CommandSeeder):
	//		cmd.RunningSeeder()
	//	default:
	//		fmt.Println("❌ Unknown command:", os.Args[1])
	//	}
	//}
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
