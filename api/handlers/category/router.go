package category

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterCategory(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerWork{DB: db, TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/category")
	user.Post("/", h.CreateCategory)
	user.Put("/", h.UpdateCategory)
	user.Delete("/:id", h.DeleteCategory)
	user.Get("/all", h.GetAllCategories)
	user.Get("/:id", h.GetCategory)
}
