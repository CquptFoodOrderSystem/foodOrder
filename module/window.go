package module

// 窗口
type Window struct {
	// 当前排队人数
	CurrentQueueLength int64 `json:"current_queue_length"`
	// 数据最后更新时间（UTC）
	LastUpdated string `json:"last_updated"`
	// 窗口唯一标识符
	WindowID string `json:"window_id"`
	// 窗口名称
	WindowName string `json:"window_name"`
}
