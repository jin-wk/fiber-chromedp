package routes

import (
	fiberSwagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jin-wk/fiber-chromedp/config"
	_ "github.com/jin-wk/fiber-chromedp/docs"
	"github.com/jin-wk/fiber-chromedp/handlers"
)

func InitRoute(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST",
		AllowCredentials: true,
	}))

	if config.Get("APP_ENV") != "production" {
		app.Get("/docs/*", fiberSwagger.HandlerDefault)
	}

	api := app.Group(("/api"))
	api.Use(logger.New(logger.Config{
		Format:     "${yellow}[${time}] ${blue}[${path}] ${green}[${method}] ${white}${body} ${white}${resBody} > ${yellow}${status}\n",
		TimeFormat: "2006-01-02 15:04:05.000",
		TimeZone:   "Asia/Seoul",
	}))
	api.Get("/health", handlers.CheckHealth)
	api.Get("/crawl/youtube", handlers.Crawl)
}
