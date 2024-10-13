package main

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/souravsk/go-zero-to-hero/user_auth/models"
	"github.com/souravsk/go-zero-to-hero/user_auth/routes"
)

func main() {
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},  // Allow requests from your Next.js app
		AllowMethods:     []string{"POST", "GET", "OPTIONS"}, // Allow POST, GET, and OPTIONS methods
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Fetch environment variables from Docker
	config := models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// Initialize DB
	models.InitDB(config)

	// Load the router
	routes.AuthRoutes(r)

	// RUN the Server
	r.Run(":8080")
}
