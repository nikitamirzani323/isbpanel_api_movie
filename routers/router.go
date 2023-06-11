package routers

import (
	"time"

	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/controllers"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        500,
		Expiration: 20 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"status":  fiber.StatusTooManyRequests,
				"message": "error many request",
				"record":  nil,
			})
		},
	}))
	app.Post("/api/init", controllers.CheckLogin)
	app.Post("/api/banner", controllers.Bannerhome)

	app.Post("/api/home", middleware.JWTProtected(), controllers.Home)
	app.Post("/api/genre", middleware.JWTProtected(), controllers.Moviegenre)
	app.Post("/api/movie", middleware.JWTProtected(), controllers.Moviehome)
	app.Post("/api/season", middleware.JWTProtected(), controllers.Movieseason)
	app.Post("/api/episode", middleware.JWTProtected(), controllers.Movieepisode)

	return app
}
