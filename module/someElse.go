package module

type QueueStatus struct {
	WindowID           string  `json:"window_id"`
	WindowName         string  `json:"window_name"`
	CurrentQueueLength int     `json:"current_queue_length"`
	AverageWaitTime    float64 `json:"average_wait_time"` // 单位：分钟
	LastUpdated        string  `json:"last_updated"`      // ISO 8601 时间字符串
}

type DishRecommendation struct {
	RecID   int     `json:"rec_id"`
	UserID  int     `json:"user_id"`
	DishID  int     `json:"dish_id"`
	Score   float64 `json:"score"`
	RecTime string  `json:"rec_time"` // ISO 8601 时间字符串
}

type UserPreference struct {
	UserID               int    `json:"user_id"`
	EncryptedPreferences string `json:"encrypted_preferences"`
}
