package module

import "gorm.io/gorm"

// 窗口
type Window struct {
	// 当前排队人数
	CurrentQueueLength int64 `json:"current_queue_length"`
	// 数据最后更新时间（UTC）
	// 窗口唯一标识符
	gorm.Model
	// 窗口名称
	WindowName string `json:"window_name"`
}

type AddWindowReq struct {
	// 当前排队人数
	CurrentQueueLen int64 `form:"currentQueueLen" validate:"required"`
	// 更新时间
	LastUpdated string `form:"lastUpdated" validate:"required"`
	// 名称
	Name string `form:"name" validate:"required"`
}
