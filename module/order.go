package module

import (
	"gorm.io/gorm"
	"time"
)

// 订单
type Order struct {
	gorm.Model
	UserID      string    `gorm:"size:50;not null" json:"user_id"`
	OrderTime   time.Time `gorm:"not null" json:"order_time"`
	TotalAmount float64   `gorm:"not null" json:"total_amount"`
	Status      string    `gorm:"size:20;not null" json:"status"` // 已完成 / 取消 / 待取餐等
}

type Subtotal struct {
	Double  *float64
	Integer *int64
}

// 订单详细表
type OrderDetail struct {
	gorm.Model
	OrderID   string  `gorm:"size:50;not null;index" json:"order_id"`
	DishID    int64   `gorm:"not null" json:"dish_id"`
	DishName  string  `gorm:"size:100;not null" json:"dish_name"`
	Quantity  int64   `gorm:"not null" json:"quantity"`
	UnitPrice float64 `gorm:"not null" json:"unit_price"`
	Subtotal  float64 `gorm:"not null" json:"subtotal"`
}

type AddOrderReq struct {
	OrderDetailsID []string `form:"orderDetailsID" validate:"required"`
	OrderTime      string   `form:"orderTime"`
	UserID         string   `form:"userID" validate:"required"`
}
