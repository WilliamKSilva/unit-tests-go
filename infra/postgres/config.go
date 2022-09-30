package postgres

import (
	"log"

	"github.com/WilliamKSilva/unit-tests-go/application/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://docker:admin@db/app"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&entities.User{}, &entities.Post{})

	return db
}
