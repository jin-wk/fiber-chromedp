package routes

import (
	fiberSwagger "github.com/arsmn/fiber-swagger"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
	"github.com/gofiber/fiber/middleware/logger"
	"github.com/jin-wk/fiber-chromedp/config"
	_ "github.com/jin-wk/fiber-chromedp/docs"
	"github.com/jin-wk/fiber-chromedp/handlers"
)

func New() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST",
		AllowCredentials: true,
	}))
	app.Use(logger.New(logger.Config{
		Format:     "${yellow}[${time}] ${blue}[${path}] ${green}[${method}] ${white}${body} ${white}${resBody} > ${yellow}${status}\n",
		TimeFormat: "2006-01-02 15:04:05.000",
		TimeZone:   "Asia/Seoul",
	}))

	if config.Get("APP_ENV") != "production" {
		app.Get("/swagger/*", fiberSwagger.Handler)
	}
	api := app.Group(("/api"))
	api.Get("/health", handlers.CheckHealth)
	api.Get("/crawl/youtube", handlers.Crawl)

	return app
}
