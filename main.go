package main

import (
	// "douyin/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"douyin/public"
)

func main() {
	public.InitDatabase()
	// go service.RunMessageServer()
	app := fiber.New()
	app.Use(logger.New())
	initRouter(app)
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		Output: os.Stdout,
	}))
	app.Listen(":8080")
}
