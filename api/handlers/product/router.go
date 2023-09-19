package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterProduct(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerWork{DB: db, TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/product")
	user.Post("/", h.CreateProduct)
	user.Put("/", h.UpdateProduct)
	user.Delete("/:id", h.DeleteProduct)
	user.Get("/all", h.GetAllProducts)
	user.Get("/:id", h.GetProduct)
}
