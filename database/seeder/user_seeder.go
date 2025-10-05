package seeder

import (
	"fmt"
	"iqbalatma/go-iqbalatma/app/model"
	"iqbalatma/go-iqbalatma/config"
	"iqbalatma/go-iqbalatma/utils"

	"github.com/jaswdr/faker"
)

var defaultUser []model.User = []model.User{
	{
		FirstName: "Iqbal",
		LastName:  "Atma Muliawan",
		Password:  "admin",
		Email:     "iqbalatma@gmail.com",
	},
}

func UserSeeder() {
	fmt.Println("Seed User")
	for _, u := range defaultUser {
		hashedPassword, _ := utils.MakeHash(u.Password)

		u.Password = hashedPassword
		config.DB.Create(&u)
	}

	fake := faker.New()
	for i := 0; i < 100; i++ {
		hashedPassword, _ := utils.MakeHash(fake.Internet().Password())
		user := model.User{
			FirstName: fake.Person().FirstName(),
			LastName:  fake.Person().LastName(),
			Email:     fake.Internet().Email(),
			Password:  hashedPassword,
		}
		config.DB.Create(&user)
	}

	fmt.Println("Seed Completed")
}
