package module

type User struct {
	UserID       int     `json:"user_id"`
	Username     string  `json:"username"`
	PasswordHash string  `json:"password_hash"`
	Gender       string  `json:"gender"` // male / female / other
	Age          int     `json:"age"`
	Height       float64 `json:"height"` // cm
	Weight       float64 `json:"weight"` // kg
	IsVegetarian bool    `json:"is_vegetarian"`
	Allergies    string  `json:"allergies"`
}
