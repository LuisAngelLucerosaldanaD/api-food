package ingredient

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Ingredient  Model struct Ingredient
type Ingredient struct {
	ID        int64     `json:"id" db:"id" valid:"-"`
	Name      string    `json:"name" db:"name" valid:"required"`
	IdProduct int64     `json:"id_product" db:"id_product" valid:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewIngredient(id int64, name string, idProduct int64) *Ingredient {
	return &Ingredient{
		ID:        id,
		Name:      name,
		IdProduct: idProduct,
	}
}

func NewCreateIngredient(name string, idProduct int64) *Ingredient {
	return &Ingredient{
		Name:      name,
		IdProduct: idProduct,
	}
}

func (m *Ingredient) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
