package main

import (
	"github.com/CquptFoodOrderSystem/foodOrder/dao"
	"github.com/CquptFoodOrderSystem/foodOrder/router"
)

func main() {
	dao.DBINit()
	router.RouterInit()
}
