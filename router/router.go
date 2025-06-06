package router

import (
	"github.com/CquptFoodOrderSystem/foodOrder/api"
	"github.com/CquptFoodOrderSystem/foodOrder/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RouterInit() {
	h := server.Default()

	user := h.Group("/user")
	{
		user.POST("/register", api.Register)
		user.GET("/login", api.Login)
	}
	dish := h.Group("/dish")
	dish.Use(middleware.AuthMiddleware())
	{
		dish.POST("/add", api.AddDish)
		dish.GET("/get", api.GetDish)
	}
	window := h.Group("/window")
	window.Use(middleware.AuthMiddleware())
	{
		window.POST("/add", api.AddWindow)
		window.GET("/getAllWindows", api.GetAllWindows)
		window.GET("/getDetails", api.GetDetails)
	}
	order := h.Group("/order")
	order.Use(middleware.AuthMiddleware())
	{
		order.POST("/post", api.PostOrder)
		order.GET("/myOrder", api.MyOrder)
		order.GET("/detail", api.OrderDetail)
	}
	someElse := h.Group("/someElse")
	someElse.Use(middleware.AuthMiddleware())
	{
		someElse.GET("/recommend", api.Recommand)
		someElse.GET("/timingQueue", api.TimingQueue)
		someElse.GET("/charge", api.Charge)
		someElse.GET("/message", api.Message)
	}

	h.Spin()
}
