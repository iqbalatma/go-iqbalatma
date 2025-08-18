package seeder

import (
	"fmt"
	"github.com/jaswdr/faker"
	"iqbalatma/go-iqbalatma/app/model"
	"iqbalatma/go-iqbalatma/config"
	"iqbalatma/go-iqbalatma/utils"
)

func UserSeeder() {
	fmt.Println("Seed User")

	fake := faker.New()
	for i := 0; i < 10; i++ {
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
