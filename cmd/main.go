package main

import (
	"github.com/CquptFoodOrderSystem/foodOrder/dao"
	"github.com/CquptFoodOrderSystem/foodOrder/router"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	dao.DBINit()
	router.RouterInit()
}
