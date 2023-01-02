package model

import (
	"time"

	"github.com/google/uuid"
)

// Partner data model
type Partner struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();PrimaryKey" json:"id"`
	Name               string    `gorm:"type:varchar(500)" json:"name"`
	PhoneNumber        string    `gorm:"type:varchar(50)" json:"phone number"`
	Address            string    `gorm:"type:varchar(500)" json:"address"`
	Email              string    `gorm:"type:varchar(100)" json:"email"`
	Representative     string    `gorm:"type:varchar(400)" json:"representative"`
	TypeOfFood         string    `gorm:"type:varchar(300)" json:"type of food"`
	AverageOrderPerDay string    `gorm:"type:int" json:"average order per day"`
	BranchCount        string    `gorm:"type:int" json:"branch count"`
	Restaurant         []Restaurant
	Contract           []Contract
}

//Delivery guys data model
type DeliveryDriver struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();PrimaryKey" json:"id"`
	Name              string    `gorm:"type:varchar(400)" json:"name"`
	PhoneNumber       string    `gorm:"type:varchar(20)" json:"phone number"`
	Address           string    `gorm:"type:varchar(500)" json:"address"`
	Email             string    `gorm:"type:varchar(200)" json:"email"`
	LicensePlates     string    `gorm:"type:varchar(30)" json:"license plates"`
	ActiveArea        string    `gorm:"type:varchar(500)" json:"active area"`
	IDNumber          string    `gorm:"type:varchar(100)" json:"id number"`
	BankAccountNumber string    `gorm:"type:varchar(50)" json:"bank account number" `
	FoodOrder         []FoodOrder
}
type Customer struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();PrimaryKey" json:"id"`
	Name            string    `gorm:"type:varchar(400)" json:"name"`
	PhoneNumber     string    `gorm:"type:varchar(50)" json:"phone number"`
	Address         string    `gorm:"type:varchar(500)" json:"address"`
	Email           string    `gorm:"type:varchar(400)" json:"email"`
	CustomerAddress []CustomerAddress
	FoodOrder       []FoodOrder
	Payment         []Payment
}

//Country name data model
type Country struct {
	ID          string `gorm:"type:varchar(10);PrimaryKey" json:"id"`
	CountryName string `gorm:"type:varchar(500)" json:"country name"`
	Address     []Address
}

//Restaurant/Store and Customer address data model
type Address struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();PrimaryKey" json:"id"`
	UnitNumber      int       `gorm:"type:int" json:"No"`
	Alley           string    `gorm:"type:varchar(50)"  json:"alley"`
	SubDistrict     string    `gorm:"type:varchar(100)" json:"sub district"`
	Street          string    `gorm:"type:varchar(150)" json:"street"`
	District        string    `gorm:"type:varchar(100)" json:"district"`
	City            string    `gorm:"type:varchar(200)" json:"city"`
	Region          string    `gorm:"type:varchar(300)" json:"region"`
	PostalCode      int       `gorm:"type:int" json:"postal code"`
	CountryID       string    `gorm:"type:varchar(10)" json:"country id"`
	Restaurant      []Restaurant
	CustomerAddress []CustomerAddress
}

//Store/Restaurant of Partner data model
type Restaurant struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();PrimaryKey" json:"ID"`
	Name      string    `gorm:"type:varchar(500)" json:"name"`
	OpenTime  string    `gorm:"type:varchar(100)" json:"open time"`
	Status    string    `gorm:"type:varchar(200)" json:"status"`
	AddressID uuid.UUID
	PartnerID uuid.UUID
	FoodOrder []FoodOrder
	MenuItem  []MenuItem
}
type OderStatus struct {
	ID          int    `gorm:"type:int;Serial PrimaryKey" json:"id"`
	StatusValue string `gorm:"varchar(100)" json:"status value"`
}

//FoodOrder data model
type FoodOrder struct {
	ID                        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();PrimaryKey" json:"ID"`
	Status                    string    `gorm:"type:varchar(200)" json:"status"`
	OrderDateTime             time.Time
	DeliveryFee               float32 `gorm:"type:float" json:"delivery fee"`
	TotalAmount               int     `gorm:"type:int" json:"total amount"`
	RequestedDeliveryDateTime time.Time
	CustomerDeliveryRating    string `gorm:"type:varchar(500)" json:"delivery rating"`
	CustomerRestaurantRating  string `gorm:"type:varchar(500)" json:"restaurant rating"`
	CustomerAddressID         uuid.UUID
	DeliveryDriverID          uuid.UUID
	CustomerID                uuid.UUID
	RestaurantID              uuid.UUID
	OrderStatusID             int
	PaymentID                 uuid.UUID
	OrderMenuItem             []OrderMenuItem
}

//Customer address data model
type CustomerAddress struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();PrimaryKey"`
	AddressID  uuid.UUID
	CustomerID uuid.UUID
	FoodOrder  []FoodOrder
}
type OrderStatus struct {
	ID          int    `gorm:"type:int;Serial PrimaryKey"`
	StatusValue string `gorm:"varchar(150)" json:"status value"`
	FoodOrder   []FoodOrder
}

//Menu data model
type MenuItem struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();PrimaryKey" json:"id"`
	ItemName      string    `gorm:"type:varchar(300)" json:"item name"`
	Describe      string    `gorm:"type:varchar(300)" json:"describe"`
	Price         float32   `gorm:"type:float" json:"price"`
	RestaurantID  uuid.UUID
	OrderMenuItem []OrderMenuItem
}

//OrderMenuItem show item that ordered
type OrderMenuItem struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	QuantityOrdered int       `gorm:"type:int" json:"quantity ordered"`
	MenuItemID      uuid.UUID
	FoodOrderID     uuid.UUID
}

//Payment show type customer can pay
type Payment struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	PayMentDate time.Time
	Total       float32
	Type        string
	FoodOrderID uuid.UUID
	CustomerID  uuid.UUID
	FoodOrder   []FoodOrder
}
type Contract struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Deputy            string    `gorm:"type:varchar(500)" json:"deputy"`
	TaxCode           string    `gorm:"type:varchar(100)" json:"tax code"`
	BranchResignCount int       `gorm:"type:int" json:"branch number"`
	BankAccountNumber string    `gorm:"type:varchar(50)" json:"account number"`
	BankName          string    `gorm:"type:varchar(500)" json:"bank name"`
	BankBranch        string    `gorm:"type:varchar(500)" json:"bank branch"`
	PartnerID         uuid.UUID
}
