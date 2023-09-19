package ingredient

import (
	"github.com/jmoiron/sqlx"

	"api-food/internal/logger"
)

const (
	Postgresql = "postgres"
)

type ServicesIngredientRepository interface {
	create(m *Ingredient) error
	update(m *Ingredient) error
	delete(id int64) error
	getByID(id int64) (*Ingredient, error)
	getAll() ([]*Ingredient, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesIngredientRepository {
	var s ServicesIngredientRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newIngredientPsqlRepository(db, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
