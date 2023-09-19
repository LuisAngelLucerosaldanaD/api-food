package category

import (
	"fmt"

	"api-food/internal/logger"
)

type PortsServerCategory interface {
	CreateCategory(name string, roleId int64) (*Category, int, error)
	UpdateCategory(id int64, name string, roleId int64) (*Category, int, error)
	DeleteCategory(id int64) (int, error)
	GetCategoryByID(id int64) (*Category, int, error)
	GetAllCategory() ([]*Category, error)
}

type service struct {
	repository ServicesCategoryRepository
	txID       string
}

func NewCategoryService(repository ServicesCategoryRepository, TxID string) PortsServerCategory {
	return &service{repository: repository, txID: TxID}
}

func (s *service) CreateCategory(name string, roleId int64) (*Category, int, error) {
	m := NewCreateCategory(name, roleId)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Category :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateCategory(id int64, name string, roleId int64) (*Category, int, error) {
	m := NewCategory(id, name, roleId)
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Category :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteCategory(id int64) (int, error) {
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

func (s *service) GetCategoryByID(id int64) (*Category, int, error) {
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

func (s *service) GetAllCategory() ([]*Category, error) {
	return s.repository.getAll()
}
