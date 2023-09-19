package product

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Product  Model struct Product
type Product struct {
	ID          int64     `json:"id" db:"id" valid:"-"`
	Name        string    `json:"name" db:"name" valid:"required"`
	UrlImg      string    `json:"url_img" db:"url_img" valid:"required"`
	Time        string    `json:"time" db:"time" valid:"required"`
	Description string    `json:"description" db:"description" valid:"required"`
	CategoryId  int64     `json:"category_id" db:"category_id" valid:"required"`
	Calorie     string    `json:"calorie" db:"calorie" valid:"required"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func NewProduct(id int64, name string, urlImg string, time string, description string, categoryId int64, calorie string) *Product {
	return &Product{
		ID:          id,
		Name:        name,
		UrlImg:      urlImg,
		Time:        time,
		Description: description,
		CategoryId:  categoryId,
		Calorie:     calorie,
	}
}

func NewCreateProduct(name string, urlImg string, time string, description string, categoryId int64, calorie string) *Product {
	return &Product{
		Name:        name,
		UrlImg:      urlImg,
		Time:        time,
		Description: description,
		CategoryId:  categoryId,
		Calorie:     calorie,
	}
}

func (m *Product) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
