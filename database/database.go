package database

import (
	"interface/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("Db.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	//db.AutoMigrate(&models.Account{}, &models.Transaction{}, &models.Balance{})
	db.AutoMigrate(&models.Account{}, &models.Transaction{})

	Database = DbInstance{
		Db: db,
	}
}
