package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jin-wk/fiber-chromedp/database"
	"github.com/jin-wk/fiber-chromedp/routes"
)

// @title Fiber-Chromedp
// @version 1.0
// @description Fiber Web Application with Chromedp
// @contact.name jin-wk
// @contact.url https://github.com/jin-wk
// @contact.email note@kakao.com
// @host localhost:5000
// @BasePath /
func main() {
	if err := database.Connect(); err != nil {
		log.Panic("Can't connect database: ", err.Error())
	}
	app := fiber.New(fiber.Config{
		Prefork: true,
	})
	routes.InitRoute(app)
	log.Fatal(app.Listen(":5000"))
}
