package main

import (
	"fmt"
	"todolist/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func helloworld(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func initDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=1234 dbname=todo port=9920 "
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connection database!")
	}
		fmt.Println("Database connected")
	    database.DBConn.AutoMigrate(&models.Todo{})
		fmt.Println('Migrated DB')
}

func setupRoutes(app *fiber.App)  {
	app.Get("/todos",models.GetTodos)
}

func main() {
	app := fiber.New()

	app.Get("/", helloworld)
    setupRoutes(app)
	app.Listen(":8000")
}
