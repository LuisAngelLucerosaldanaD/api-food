package category

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Category  Model struct Category
type Category struct {
	ID        int64     `json:"id" db:"id" valid:"-"`
	Name      string    `json:"name" db:"name" valid:"required"`
	RoleId    int64     `json:"role_id" db:"role_id" valid:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewCategory(id int64, name string, roleId int64) *Category {
	return &Category{
		ID:     id,
		Name:   name,
		RoleId: roleId,
	}
}

func NewCreateCategory(name string, roleId int64) *Category {
	return &Category{
		Name:   name,
		RoleId: roleId,
	}
}

func (m *Category) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
