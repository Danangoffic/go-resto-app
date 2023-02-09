package menu

type MenuItem struct {
	Name      string `json:"name"`
	OrderCode string `json:"order_code"`
	Price     int    `json:"price"`
	Type      string `json:"type"`
}
