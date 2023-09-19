package rol

import (
	"github.com/jmoiron/sqlx"

	"api-food/internal/logger"
)

const (
	Postgresql = "postgres"
)

type ServicesRolRepository interface {
	create(m *Rol) error
	update(m *Rol) error
	delete(id int64) error
	getByID(id int64) (*Rol, error)
	getAll() ([]*Rol, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesRolRepository {
	var s ServicesRolRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newRolPsqlRepository(db, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
