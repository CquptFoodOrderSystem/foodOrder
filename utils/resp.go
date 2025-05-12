package utils

import (
	"context"
	"github.com/CquptFoodOrderSystem/foodOrder/module"
	"github.com/cloudwego/hertz/pkg/app"
)

func SuccessResp(ctx context.Context, c *app.RequestContext) {
	c.JSON(200, module.Response{
		Code: 200,
		Msg:  "success",
	})
	return
}

func FailResp(ctx context.Context, c *app.RequestContext, err string) {
	c.JSON(400, module.Response{
		Code: 400,
		Msg:  err,
	})
	return
}

func RespWithData(ctx context.Context, c *app.RequestContext, any2 interface{}) {
	c.JSON(200, module.Response{
		Code: 200,
		Msg:  "success",
		Data: any2,
	})
	return
}
