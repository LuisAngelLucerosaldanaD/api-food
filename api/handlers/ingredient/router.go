package ingredient

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterIngredient(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerWork{DB: db, TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/ingredient")
	user.Post("/", h.CreateIngredient)
	user.Put("/", h.UpdateIngredient)
	user.Delete("/:id", h.DeleteIngredient)
	user.Get("/all", h.GetAllIngredients)
	user.Get("/:id", h.GetIngredient)
}
