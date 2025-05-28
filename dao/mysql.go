package dao

import (
	"github.com/CquptFoodOrderSystem/foodOrder/module"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func MysqlInit() {
	dsn := "root:123456@tcp(127.0.0.1:3307)/foodOrder?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
		log.Println("AutoMigrate1 err : ", err)
	}
}
