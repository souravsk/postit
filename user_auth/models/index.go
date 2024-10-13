package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// struct holds the configuration values for the database connection
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var DB *gorm.DB //global variable that will store the instance of the database connection.

// InitDB function is used to initialize the database connection using the values from the Config struct.
func InitDB(cfg Config) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	//It automatically migrates the User model by calling AutoMigrate. This creates the necessary table in the database if it doesn't exist. If there is an error, it panics.
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}

	fmt.Println("Migrated database")

	// It sets the global DB variable to the instance of the database connection.
	DB = db

}
