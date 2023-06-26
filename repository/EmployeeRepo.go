package repository

import (
	"API/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type employeeRepository struct {
	Db *gorm.DB
}
type EmployeeRepository interface {
	ListContract() (model.Contract, error)
	GetContract(uuid.UUID) (model.Contract, error)
}
