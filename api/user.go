package api

import (
	"context"
	"errors"
	"github.com/CquptFoodOrderSystem/foodOrder/dao"
	"github.com/CquptFoodOrderSystem/foodOrder/module"
	"github.com/CquptFoodOrderSystem/foodOrder/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

const (
	JwtKey = "123"
)

func Register(ctx context.Context, c *app.RequestContext) {
	var req module.RegisterReq
	var usr module.User
	err := c.BindAndValidate(&req)
	if err != nil {
		utils.FailResp(ctx, c, "参数绑定或校验错误："+err.Error())
		return
	}
	if req.PassWord != req.RePassWord {
		utils.FailResp(ctx, c, "请输入两次相同密码")
		return
	}
	if err = dao.DB.Where("username = ?", req.UserName).First(&usr).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.FailResp(ctx, c, "该用户名已注册："+err.Error())
		return
	}
	passwordHash, err := utils.HashPassword(req.PassWord)
	if err != nil {
		utils.FailResp(ctx, c, err.Error())
		return
	}
	usr = module.User{
		Username:     req.UserName,
		PasswordHash: passwordHash,
		Gender:       "",
		Height:       float64(req.Height),
		Weight:       float64(req.Weight),
		Allergies:    req.Allergen,
	}
	if err = dao.DB.Create(&usr).Error; err != nil {
		utils.FailResp(ctx, c, "数据库操作错误："+err.Error())
		return
	}
	token, err := utils.GenerateJwt([]byte(JwtKey), jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  usr.ID,
		"name": usr.Username,
	})
	if err != nil {
		utils.FailResp(ctx, c, "token产生错误："+err.Error())
		return
	}
	utils.RespWithData(ctx, c, map[string]interface{}{
		"id":       usr.ID,
		"userName": usr.Username,
		"token":    token,
	})
}

func Login(ctx context.Context, c *app.RequestContext) {
	var req module.LoginReq
	var usr module.User
	err := c.BindAndValidate(&req)
	if err != nil {
		utils.FailResp(ctx, c, "参数绑定或校验错误："+err.Error())
		return
	}
	if err = dao.DB.Where("username = ?", req.UserName).Find(&usr).Error; err != nil {
		utils.FailResp(ctx, c, "该用户名尚未注册："+err.Error())
		return
	}
	if is := utils.CheckPassword(usr.PasswordHash, req.PassWord); !is {
		utils.FailResp(ctx, c, "密码错误："+err.Error())
		return
	}
	token, err := utils.GenerateJwt([]byte(JwtKey), jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  usr.ID,
		"name": usr.Username,
	})
	if err != nil {
		utils.FailResp(ctx, c, "token产生错误："+err.Error())
		return
	}
	utils.RespWithData(ctx, c, map[string]interface{}{
		"id":       usr.ID,
		"userName": usr.Username,
		"token":    token,
	})
}
