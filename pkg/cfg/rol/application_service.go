package rol

import (
	"fmt"

	"api-food/internal/logger"
)

type PortsServerRol interface {
	CreateRol(name string) (*Rol, int, error)
	UpdateRol(id int64, name string) (*Rol, int, error)
	DeleteRol(id int64) (int, error)
	GetRolByID(id int64) (*Rol, int, error)
	GetAllRol() ([]*Rol, error)
}

type service struct {
	repository ServicesRolRepository
	txID       string
}

func NewRolService(repository ServicesRolRepository, TxID string) PortsServerRol {
	return &service{repository: repository, txID: TxID}
}

func (s *service) CreateRol(name string) (*Rol, int, error) {
	m := NewCreateRol(name)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Rol :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateRol(id int64, name string) (*Rol, int, error) {
	m := NewRol(id, name)
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Rol :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteRol(id int64) (int, error) {
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

func (s *service) GetRolByID(id int64) (*Rol, int, error) {
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

func (s *service) GetAllRol() ([]*Rol, error) {
	return s.repository.getAll()
}
