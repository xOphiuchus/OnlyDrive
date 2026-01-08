package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xOphiuchus/OnlyDrive/middlewares"
	"github.com/xOphiuchus/OnlyDrive/services"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.Router, db *gorm.DB) {
	(*app).Use(middlewares.Recovery())
	(*app).Use(middlewares.Logger())
	(*app).Use(middlewares.CORS())
	(*app).Use(middlewares.RateLimiter())
	(*app).Get("/healthz", func(c *fiber.Ctx) error {
		return services.HealthCheck(c)
	})

	// Start Routing Setup Here
}
