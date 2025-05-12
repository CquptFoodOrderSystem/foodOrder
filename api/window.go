package api

import (
	"context"
	"github.com/CquptFoodOrderSystem/foodOrder/dao"
	"github.com/CquptFoodOrderSystem/foodOrder/module"
	"github.com/CquptFoodOrderSystem/foodOrder/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

func AddWindow(ctx context.Context, c *app.RequestContext) {
	var addReq module.AddWindowReq
	var window module.Window
	err := c.BindAndValidate(&addReq)
	if err != nil {
		utils.FailResp(ctx, c, "参数绑定或校验错误："+err.Error())
		return
	}
	window = module.Window{
		CurrentQueueLength: 0,
		WindowName:         addReq.Name,
	}
	err = dao.DB.Create(&window).Error
	if err != nil {
		utils.FailResp(ctx, c, "数据库操作错误："+err.Error())
		return
	}
	utils.RespWithData(ctx, c, window)
}

func GetAllWindows(ctx context.Context, c *app.RequestContext) {
	var windows []module.Window
	err := dao.DB.Find(&windows).Error
	if err != nil {
		utils.FailResp(ctx, c, "数据库操作错误："+err.Error())
		return
	}
	utils.RespWithData(ctx, c, windows)
}

func GetDetails(ctx context.Context, c *app.RequestContext) {
	var dishs []module.Dish
	WinId := c.Query("id")
	err := dao.DB.Where("window_id = ?", WinId).Find(&dishs).Error
	if err != nil {
		utils.FailResp(ctx, c, "数据库操作错误："+err.Error())
		return
	}
	utils.RespWithData(ctx, c, dishs)
}
