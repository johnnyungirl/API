package repository

import (
	"API/model"

	"gorm.io/gorm"
)

type customerRepository struct {
	DB *gorm.DB
}
type CustomerRepository interface {
	ListPartner() ([]model.Partner, error)
	ListMenu() ([]model.MenuItem, error)
	GetMenuItem(string) (model.MenuItem, error)
	GetPartner(string) (model.Partner, error)
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return customerRepository{
		DB: db,
	}
}
func (cust customerRepository) ListPartner() (partner []model.Partner, err error) {
	return partner, cust.DB.Find(&partner).Error
}
func (cust customerRepository) ListMenu() (menu []model.MenuItem, err error) {
	return menu, cust.DB.Find(&menu).Error
}
func (cust customerRepository) GetMenuItem(Name string) (menu model.MenuItem, err error) {
	return menu, cust.DB.First(&menu, "Name", Name).Error
}
func (cust customerRepository) GetPartner(Name string) (partner model.Partner, err error) {
	return partner, cust.DB.First(&partner).Error
}
