package main

import (
	"log"

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
	app := routes.New()
	log.Fatal(app.Listen(":5000"))
}
