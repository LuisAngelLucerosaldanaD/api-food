package ingredient

import (
	"fmt"

	"api-food/internal/logger"
)

type PortsServerIngredient interface {
	CreateIngredient(name string, idProduct int64) (*Ingredient, int, error)
	UpdateIngredient(id int64, name string, idProduct int64) (*Ingredient, int, error)
	DeleteIngredient(id int64) (int, error)
	GetIngredientByID(id int64) (*Ingredient, int, error)
	GetAllIngredient() ([]*Ingredient, error)
}

type service struct {
	repository ServicesIngredientRepository
	txID       string
}

func NewIngredientService(repository ServicesIngredientRepository, TxID string) PortsServerIngredient {
	return &service{repository: repository, txID: TxID}
}

func (s *service) CreateIngredient(name string, idProduct int64) (*Ingredient, int, error) {
	m := NewCreateIngredient(name, idProduct)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Ingredient :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateIngredient(id int64, name string, idProduct int64) (*Ingredient, int, error) {
	m := NewIngredient(id, name, idProduct)
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Ingredient :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteIngredient(id int64) (int, error) {
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

func (s *service) GetIngredientByID(id int64) (*Ingredient, int, error) {
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

func (s *service) GetAllIngredient() ([]*Ingredient, error) {
	return s.repository.getAll()
}
