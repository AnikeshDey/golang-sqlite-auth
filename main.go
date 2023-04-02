package main

import (
	"fmt"
	"log"

	"auth/database"
	"auth/models"
	"auth/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "auth.db")
	if err != nil {
		panic("Failed to connect to database...")
	}

	fmt.Println("Connected to Database...")

	database.DBConn.AutoMigrate(&models.User{})
	fmt.Println("Database Migrated Successfully...")
}

func main() {
	app := fiber.New()

	initDatabase()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.SetupRoutes(app)

	defer database.DBConn.Close()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":8080"))
}
