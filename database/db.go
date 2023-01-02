package database

import (
	"API/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Unable to connect database")

	}
	return db
}
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Partner{})
	db.AutoMigrate(&model.Customer{})
	db.AutoMigrate(&model.Address{})
	db.AutoMigrate(&model.Country{})
	db.AutoMigrate(&model.CustomerAddress{})
	db.AutoMigrate(&model.DeliveryDriver{})
	db.AutoMigrate(&model.OrderStatus{})
	db.AutoMigrate(&model.OrderMenuItem{})
	db.AutoMigrate(&model.OderStatus{})
	db.AutoMigrate(&model.MenuItem{})
	db.AutoMigrate(&model.Restaurant{})
	db.AutoMigrate(&model.Payment{})
	db.AutoMigrate(&model.Contract{})

}
