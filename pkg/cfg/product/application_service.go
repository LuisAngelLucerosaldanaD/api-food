package product

import (
	"fmt"

	"api-food/internal/logger"
)

type PortsServerProduct interface {
	CreateProduct(name string, urlImg string, time string, description string, categoryId int64, calorie string) (*Product, int, error)
	UpdateProduct(id int64, name string, urlImg string, time string, description string, categoryId int64, calorie string) (*Product, int, error)
	DeleteProduct(id int64) (int, error)
	GetProductByID(id int64) (*Product, int, error)
	GetAllProduct() ([]*Product, error)
}

type service struct {
	repository ServicesProductRepository
	txID       string
}

func NewProductService(repository ServicesProductRepository, TxID string) PortsServerProduct {
	return &service{repository: repository, txID: TxID}
}

func (s *service) CreateProduct(name string, urlImg string, time string, description string, categoryId int64, calorie string) (*Product, int, error) {
	m := NewCreateProduct(name, urlImg, time, description, categoryId, calorie)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Product :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateProduct(id int64, name string, urlImg string, time string, description string, categoryId int64, calorie string) (*Product, int, error) {
	m := NewProduct(id, name, urlImg, time, description, categoryId, calorie)
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Product :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteProduct(id int64) (int, error) {
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return 15, fmt.Errorf("id is required")
	}

	if err := s.repository.delete(id); err != nil {
		if err.Error() == "ecatch:108" {
			return 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't update row:", err)
		return 20, err
	}
	return 28, nil
}

func (s *service) GetProductByID(id int64) (*Product, int, error) {
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return nil, 15, fmt.Errorf("id is required")
	}
	m, err := s.repository.getByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAllProduct() ([]*Product, error) {
	return s.repository.getAll()
}
