package main

import (
	"log"
	"miniproject/config"
	"miniproject/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Database initialization
	config.InitDB()

	// Routes
	routes.RegisterRoutes(e)

	// Start the server
	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))

}
