package api

import (
	"context"
	"github.com/CquptFoodOrderSystem/foodOrder/dao"
	"github.com/CquptFoodOrderSystem/foodOrder/module"
	"github.com/CquptFoodOrderSystem/foodOrder/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

func AddDish(ctx context.Context, c *app.RequestContext) {
	var addReq module.AddReq
	var dish module.Dish
	err := c.BindAndValidate(&addReq)
	if err != nil {
		utils.FailResp(ctx, c, "参数绑定或校验错误："+err.Error())
		return
	}
	dish = module.Dish{
		Calories: addReq.Calories,
		Carbs:    addReq.Carbs,
		Category: addReq.Category,
		Fat:      addReq.Fat,
		IsHot:    addReq.IsHot,
		IsSpicy:  addReq.IsSpicy,
		Name:     addReq.Name,
		Price:    float64(addReq.Price),
		Protein:  addReq.Protein,
		WindowID: addReq.WindowID,
	}
	err = dao.DB.Create(&dish).Error
	if err != nil {
		utils.FailResp(ctx, c, "数据库操作失败："+err.Error())
		return
	}
	utils.RespWithData(ctx, c, map[string]interface{}{
		"id":   dish.ID,
		"name": dish.Name,
	})
}

func GetDish(ctx context.Context, c *app.RequestContext) {
	var dish module.Dish
	ID := c.Query("id")
	if ID == "" {
		utils.FailResp(ctx, c, "参数绑定或校验错误：")
		return
	}
	err := dao.DB.Where("id = ?", ID).Find(&dish).Error
	if err != nil {
		utils.FailResp(ctx, c, "数据库操作失败："+err.Error())
		return
	}
	utils.RespWithData(ctx, c, dish)
}
