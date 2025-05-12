package api

//
//import (
//	"context"
//	"github.com/CquptFoodOrderSystem/foodOrder/dao"
//	"github.com/CquptFoodOrderSystem/foodOrder/module"
//	"github.com/CquptFoodOrderSystem/foodOrder/utils"
//	"github.com/cloudwego/hertz/pkg/app"
//	"time"
//)
//
//func PostOrder(ctx context.Context, c *app.RequestContext) {
//	var addOrderReq module.AddOrderReq
//	var order module.Order
//	err := c.BindAndValidate(&addOrderReq)
//	if err != nil {
//		utils.FailResp(ctx, c, "参数绑定或校验错误："+err.Error())
//		return
//	}
//	order = module.Order{
//		UserID:    addOrderReq.UserID,
//		OrderTime: time.Now(),
//		TotalAmount:,
//		Status:    "",
//	}
//	dao.DB.Create(&order)
//	var Amout float64
//	for _, s := range addOrderReq.OrderDetailsID {
//		var dish module.Dish
//		var orderDetail module.OrderDetail
//		err = dao.DB.Select("id","name").Where("id = ?", s).First(&dish).Error
//		if err != nil {
//			utils.FailResp(ctx, c, "查询菜品错误："+err.Error())
//			return
//		}
//		Amout = Amout + dish.Price
//		orderDetail = module.OrderDetail{
//			OrderID:   order.OrderID,
//			DishID:    0,
//			DishName:  "",
//			Quantity:  0,
//			UnitPrice: 0,
//			Subtotal:  0,
//		}
//	}
//	order = module.Order{
//		OrderID:   "",
//		UserID:    addOrderReq.UserID,
//		OrderTime: time.Now(),
//		TotalAmount:,
//		Status:    "",
//	}
//}
//
//func MyOrder(ctx context.Context, c *app.RequestContext) {
//
//}
//
//func OrderDetail(ctx context.Context, c *app.RequestContext) {
//
//}
