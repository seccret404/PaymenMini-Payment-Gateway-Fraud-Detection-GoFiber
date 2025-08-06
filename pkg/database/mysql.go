package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
	var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Connected to MySQL!")
			//auto migrate
			err = db.AutoMigrate(
				&domain.User{},
				&domain.Product{},
				&domain.Cart{},
			)

			DB = db
			return db, nil
		}
		log.Printf(" Waiting for database... (%d/10)", i+1)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to DB after retries: %w", err)
}
