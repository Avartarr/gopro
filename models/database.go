package models

import (
	"log"
	// "os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// dsn := "host=" + os.Getenv("DB_HOST") +
	// 	" user=" + os.Getenv("DB_USER") +
	// 	" password=" + os.Getenv("DB_PASSWORD") +
	// 	" dbname=" + os.Getenv("DB_NAME") +
	// 	" port=" + os.Getenv("DB_PORT") +
	// 	" sslmode=disable"

	dsn := "host=localhost user=postgres password=Chiamaka@06 dbname=dukiago port=5433 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to open database", err)
	}

	db.AutoMigrate(&User{})
	DB = db
}
