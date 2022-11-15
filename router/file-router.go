package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saint-rivers/saint-ark/handler"
	"go.mongodb.org/mongo-driver/mongo"
)

func ApplyFileRouter(app *fiber.App, client *mongo.Client) *fiber.App {

	app.Get("/api/v1/images", func(c *fiber.Ctx) error {
		return handler.GetListedImages(c, client)
	})

	app.Post("/api/v1/images", func(c *fiber.Ctx) error {
		return handler.InsertImage(c, client)
	})

	return app
}
