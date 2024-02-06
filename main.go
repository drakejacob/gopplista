package main

import (
	"gopplista/app/routes"
	"gopplista/db"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {
	db := db.Init()

	f := fiber.New(fiber.Config{
		Views: html.New("./app", ".html"),
	})

	f.Use(logger.New())

	f.Static("/", "./app/static")

	app := f.Group("/")
	routes.RegisterRoutes(app, db)

	f.Use(func(c *fiber.Ctx) error {
		return c.Status(404).Render("routes/404", nil, "layouts/base")
	})

	log.Fatal(f.Listen("localhost:3333"))
}