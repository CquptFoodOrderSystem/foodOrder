package module

// 商品
type Dish struct {
	// 热量（千卡）
	Calories int64 `json:"calories"`
	// 碳水化合物（克）
	Carbs int64 `json:"carbs"`
	// 类别：主食、配菜、饮品等
	Category string `json:"category"`
	// 菜品id
	DishID int64 `json:"dish_id"`
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
