package main

import (
	"API/database"
	"API/route"
	config "API/util"
	"fmt"
	"log"
)

//cách mặc định foreign key= struct owner's name+ owner's primary key
// type FoodOrder struct {
// 	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
// 	OrderStatusID int
// 	OrderStatus   OrderStatus
// }

// type OrderStatus struct {
// 	ID   int
// 	Name string `gorm:"type:varchar(100)"`
// }

// cách đổi tên foreign key -> sử dụng `gorm:"ForeignKey:???"`
// type FoodOrder struct {
// 	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
// 	OrderRefer int
// 	OrderStatus OrderStatus `gorm:"foreignKey:OrderRefer"`
// }

// type OrderStatus struct {
// 	ID   int
// 	Name string `gorm:"type:varchar(100)"`
// }
// type User struct {
// 	gorm.Model
// 	ID           int
// 	MemberNumber string
// 	CreditCards  []CreditCard `gorm:"foreignKey:UserNumber;references:MemberNumber"`
// }

// type CreditCard struct {
// 	gorm.Model
// 	ID         int
// 	Number     string
// 	UserNumber string
// }
// User has many CreditCards, UserID is the foreign key
func main() {
	
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Enviroment file no found")

	}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.Db.Username, config.Db.Password, config.Db.Host, config.Db.Port, config.Db.Dbname)
	db := database.ConnectDB(dsn)
	database.Migrate(db)
	route.SetupRoutes(db)
	

}
