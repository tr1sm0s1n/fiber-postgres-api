package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/tr1sm0s1n/fiber-postgres-api/controllers"
	"github.com/tr1sm0s1n/fiber-postgres-api/db"
	"github.com/tr1sm0s1n/fiber-postgres-api/middlewares"
	"github.com/tr1sm0s1n/fiber-postgres-api/models"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Fatal("Failed to connect the database")
	}

	db.AutoMigrate(&models.Certificate{})

	app := fiber.New()
	app.Use(logger.New())
	app.Use(func(ctx *fiber.Ctx) error {
		return middlewares.Authority(ctx)
	})

	controllers := controllers.NewControllers(db)

	app.Post("/create", controllers.CreateOne)
	app.Get("/read", controllers.ReadAll)
	app.Get("/read/:id", controllers.ReadOne)
	app.Put("/update/:id", controllers.UpdateOne)
	app.Delete("/delete/:id", controllers.DeleteOne)

	app.Listen(":8080")
}
