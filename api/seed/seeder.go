package seed

import (
	"log"

	"github.com/INI-MED/gocrud/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "USER1",
		Email:    "USER1@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "USER2",
		Email:    "USER2@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "USER3",
		Email:    "USER3@gmail.com",
		Password: "password",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
