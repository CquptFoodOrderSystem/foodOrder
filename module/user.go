package module

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string  `json:"username"`
	PasswordHash string  `json:"password_hash"`
	Gender       string  `json:"gender"` // male / female / other
	Age          int     `json:"age"`
	Height       float64 `json:"height"` // cm
	Weight       float64 `json:"weight"` // kg
	IsVegetarian bool    `json:"is_vegetarian"`
	Allergies    string  `json:"allergies"`
}

type RegisterReq struct {
	// 过敏原
	Allergen string `form:"allergen"`
	// 身高
	Height int64 `form:"height"`
	// 密码
	PassWord string `form:"passWord" validate:"required"`
	// 确认密码
	RePassWord string `form:"rePassWord" validate:"required"`
	// 用户名
	UserName string `form:"userName" validate:"required"`
	// 体重
	Weight int64 `form:"weight"`
}

type LoginReq struct {
	// 密码
	PassWord string `form:"passWord" validate:"required"`
	// 用户名
	UserName string `form:"userName" validate:"required"`
}
