package main

import (
	"fmt"
	"log"
	"test/database"
	"test/pkg/mysql"
	"test/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))
	mysql.DatabaseConnection()
	database.RunMigration()
	routes.RouteInit(e.Group("/api/v1"))

	fmt.Println("Runing on Port : 5200")
	e.Logger.Fatal(e.Start("localhost:5200"))
}