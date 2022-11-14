package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"mail.blast/Config"

	"mail.blast/Models"
	"mail.blast/Routes"
)

func main() {

	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := Config.DbURL(Config.BuildDBConfig())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("Status:", err)
	}
	close, _ := db.DB()
	defer close.Close()
	db.AutoMigrate(&Models.EmailAccount{})
	db.AutoMigrate(&Models.EmailTemplate{})
	db.AutoMigrate(&Models.Publisher{})
	db.AutoMigrate(&Models.PublisherParticipant{})
	db.AutoMigrate(&Models.User{})

	r := Routes.SetupRouter()
	//running

	r.Run()

}
