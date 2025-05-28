package module

import "gorm.io/gorm"

// 商品
type Dish struct {
	// 热量（千卡）
	Calories int64 `json:"calories"`
	// 碳水化合物（克）
	Carbs int64 `json:"carbs"`
	// 类别：主食、配菜、饮品等
	Category string `json:"category"`
	// 菜品id
	gorm.Model
	// 脂肪（克）
	Fat int64 `json:"fat"`
	// 是否高温菜（如油炸）
	IsHot bool `json:"is_hot"`
	// 是否辣
	IsSpicy bool `json:"is_spicy"`
	// 菜品名称
	Name string `json:"name"`
	// 单价
	Price float64 `json:"price"`
	// 蛋白质（克）
	Protein int64 `json:"protein"`
	// 窗口id
	WindowID int64 `json:"window_id"`
}

type AddReq struct {
	// 热量
	Calories int64 `form:"calories" validate:"required"`
	// 碳水
	Carbs int64 `form:"carbs" validate:"required"`
	// 特色描述
	Category string `form:"category" validate:"required"`
	// 脂肪
	Fat int64 `form:"fat" validate:"required"`
	// 是否高温菜（如油炸）
	IsHot bool `form:"isHot" validate:"required"`
	// 是否辛辣
	IsSpicy bool `form:"isSpicy" validate:"required"`
	// 名称
	Name string `form:"name" validate:"required"`
	// 价格
	Price int64 `form:"price" validate:"required"`
	// 蛋白质
	Protein int64 `form:"protein" validate:"required"`
	// 窗口id
	WindowID int64 `form:"windowId" validate:"required"`
}
