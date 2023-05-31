package migration

import (
	"fmt"
	"rest-go/database"
	"rest-go/model/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database migration successful")
}
