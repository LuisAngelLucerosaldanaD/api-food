package api

import (
	"api-food/api/handlers/category"
	"api-food/api/handlers/ingredient"
	"api-food/api/handlers/rol"
	_ "api-food/docs"
	"github.com/ansrivas/fiberprometheus/v2"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func routes(db *sqlx.DB, loggerHttp bool, allowedOrigins string) *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit: 30 * 1024 * 1024,
	})

	prometheus := fiberprometheus.New("OnlyOne Smart Contract")
	prometheus.RegisterAt(app, "/metrics")

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))

	app.Use(recover.New())
	app.Use(prometheus.Middleware)
	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowHeaders: "Origin, X-Requested-With, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST",
	}))
	if loggerHttp {
		app.Use(logger.New())
	}
	TxID := uuid.New().String()

	loadRoutes(app, db, TxID)

	return app
}

func loadRoutes(app *fiber.App, db *sqlx.DB, TxID string) {
	category.RouterCategory(app, db, TxID)
	ingredient.RouterIngredient(app, db, TxID)
	rol.RouterRol(app, db, TxID)
}
