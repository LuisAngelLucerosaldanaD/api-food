package rol

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Rol  Model struct Rol
type Rol struct {
	ID        int64     `json:"id" db:"id" valid:"-"`
	Name      string    `json:"name" db:"name" valid:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewRol(id int64, name string) *Rol {
	return &Rol{
		ID:   id,
		Name: name,
	}
}

func NewCreateRol(name string) *Rol {
	return &Rol{
		Name: name,
	}
}

func (m *Rol) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
