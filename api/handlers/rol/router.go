package rol

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterRol(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerWork{DB: db, TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/rol")
	user.Post("/", h.CreateRol)
	user.Put("/", h.UpdateRol)
	user.Delete("/:id", h.DeleteRol)
	user.Get("/all", h.GetAllRoles)
	user.Get("/:id", h.GetRol)
}
