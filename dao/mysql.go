package dao

import (
	"github.com/CquptFoodOrderSystem/foodOrder/module"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func MysqlInit() {
	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		log.Println("mysql init err : ", err)
	}
	DB = db
	err = DB.AutoMigrate(
		&module.User{},
		&module.Dish{},
		&module.DishRecommendation{},
		&module.OrderDetails{},
		&module.Order{},
		&module.QueueStatus{},
		&module.Subtotal{},
		&module.UserPreference{},
		&module.Window{},
	)
	if err != nil {
		log.Println("AutoMigrate err : ", err)
	}
}
