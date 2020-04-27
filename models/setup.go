package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/raviwu/gobookstore/config"
)

func SetupModels() *gorm.DB {
	config := config.Load("../config.yaml")

	fmt.Printf("Check the loaded config: %+v\n", config)

	driver := config.Database.Driver
	database := config.Database.Database

	db, err := gorm.Open(driver, database)

	if err != nil {
		panic("failing db connection")
	}

	db.AutoMigrate(&Book{})

	return db
}
