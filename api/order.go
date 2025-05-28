package api

import (
	"context"
	"github.com/CquptFoodOrderSystem/foodOrder/dao"
	"github.com/CquptFoodOrderSystem/foodOrder/module"
	"github.com/CquptFoodOrderSystem/foodOrder/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"time"
)

const (
	OrderCreate = "waiting"
	OrderDoing  = "doing"
	OrderDone   = "done"
)

func PostOrder(ctx context.Context, c *app.RequestContext) {
	var addOrderReq module.AddOrderReq
	var order module.Order
	err := c.BindAndValidate(&addOrderReq)
	if err != nil {
		utils.FailResp(ctx, c, "参数绑定或校验错误："+err.Error())
		return
	}
	order = module.Order{
		UserID:      addOrderReq.UserID,
		OrderTime:   time.Now(),
		TotalAmount: 0,
		Status:      OrderCreate,
	}
	err = dao.DB.Create(&order).Error
	if err != nil {
		utils.FailResp(ctx, c, "订单创建错误："+err.Error())
		return
	}
	var Amout float64
	for _, s := range addOrderReq.OrderDetailsID {
		var dish module.Dish
		var orderDetail module.OrderDetails
		err = dao.DB.Select("id", "name", "price").Where("id = ?", s).First(&dish).Error
		if err != nil {
			utils.FailResp(ctx, c, "查询菜品错误："+err.Error())
			return
		}
		var window module.Window
		err = dao.DB.Where("id = ?", dish.WindowID).First(&window).Error
		if err != nil {
			utils.FailResp(ctx, c, "查询窗口错误："+err.Error())
			return
		}
		err = dao.DB.Model(&window).Update("current_queue_length", window.CurrentQueueLength+1).Error
		if err != nil {
			utils.FailResp(ctx, c, "更新排队人数错误："+err.Error())
			return
		}
		Amout = Amout + dish.Price
		orderDetail = module.OrderDetails{
			OrderID:   strconv.Itoa(int(order.ID)),
			DishID:    int64(dish.ID),
			DishName:  dish.Name,
			Quantity:  1,
			UnitPrice: dish.Price,
			Subtotal:  dish.Price,
		}
		err = dao.DB.Create(&orderDetail).Error
		if err != nil {
			utils.FailResp(ctx, c, "创建商品订单详细错误："+err.Error())
			return
		}
	}
	err = dao.DB.Model(&order).Updates(module.Order{OrderTime: time.Now(), TotalAmount: Amout}).Error
	if err != nil {
		utils.FailResp(ctx, c, "更新商品订单错误："+err.Error())
		return
	}
	utils.RespWithData(ctx, c, map[string]interface{}{
		"order_id":         order.ID,
		"totalAmount":      Amout,
		"order_details_id": addOrderReq.OrderDetailsID,
	})
	return
}

func MyOrder(ctx context.Context, c *app.RequestContext) {
	var orders []module.Order
	uid := c.Query("id")
	err := dao.DB.Where("user_id = ?", uid).Find(&orders).Error
	if err != nil {
		utils.FailResp(ctx, c, "查询订单错误："+err.Error())
		return
	}
	utils.RespWithData(ctx, c, orders)
	return
}

func OrderDetail(ctx context.Context, c *app.RequestContext) {
	var orderDetails []module.OrderDetails
	orderId := c.Query("id")
	err := dao.DB.Where("order_id = ?", orderId).Find(&orderDetails).Error
	if err != nil {
		utils.FailResp(ctx, c, "查询订单详细错误："+err.Error())
		return
	}
	utils.RespWithData(ctx, c, orderDetails)
	return
}
