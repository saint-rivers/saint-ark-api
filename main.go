package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/saint-rivers/saint-ark/config"
	"github.com/saint-rivers/saint-ark/router"
)

func main() {
	client := config.Setup()

	// creating routers
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app = router.ApplyFileRouter(app, client)

	fmt.Println("Application running on port 8080")
	app.Listen(":8080")

	client.Disconnect(context.TODO())
}
