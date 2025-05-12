package module

type Response struct {
	Code int
	Msg  string
	Data interface{} `json:"data,omitempty"`
}
