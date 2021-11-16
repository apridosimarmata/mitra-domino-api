package main

import (
	"domino/domino/Config"
	"domino/domino/Models"
	"domino/domino/Routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status: ", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(
		&Models.User{},
		&Models.Order{},
	)

	Config.DB.Model(&Models.Order{}).AddForeignKey("user_id", "user(id)", "RESTRICT", "RESTRICT")

	r := Routes.SetupRouter()

	r.Run()
}
