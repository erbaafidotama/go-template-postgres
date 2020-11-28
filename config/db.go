package config

import (
	"admin-rt/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB declare gorm
// var DB *gorm.DB

// func GetDB() *gorm.DB {
// 	return db
// }

// InitDB config
// func InitDB() {
// 	// get ENV
// 	gotenv.Load()

// 	var err error
// 	dbName := os.Getenv("DB_NAME")

// 	// connect to db
// 	dsn := "user=postgres password=postgres dbname=" + dbName + " port=5432 sslmode=disable TimeZone=Asia/Shanghai"

// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	fmt.Printf("\n\n%#v\n\n", DB)
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	// migrate table
// 	DB.AutoMigrate(&models.Account{})
// 	// DB.AutoMigrate(&models.Item{})
// 	// DB.AutoMigrate(&models.TransactionItem{})
// }

func GetDB() *gorm.DB {
	godotenv.Load()
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	fmt.Println(dbName)
	dsn := "user=" + dbUsername + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Printf("\n\n%#v\n\n", db)
	if err != nil {
		panic("Connecting database failed:" + err.Error())
	}
	db.AutoMigrate(&models.Account{})
	return db
}
