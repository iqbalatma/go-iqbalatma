package seeder

import (
	"fmt"
	"iqbalatma/go-iqbalatma/app/model"
	"iqbalatma/go-iqbalatma/config"

	"github.com/jaswdr/faker"
)

func AccountSiteSeeder() {
	fmt.Println("Seed account site")

	fake := faker.New()

	for i := 0; i < 10; i++ {
		name := fake.Internet().Email()
		description := fake.Lorem().Paragraph(3)

		fmt.Printf("%d Create account_sites for name : %s \n", i+1, name)
		accountSite := model.AccountSite{
			Name:        name,
			Url:         fake.Internet().Domain(),
			Description: &description,
			Icon:        nil,
		}

		config.DB.Create(&accountSite)
	}
}
