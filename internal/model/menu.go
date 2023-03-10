package model

type MenuItem struct {
	ID        uint   `gorm:"primaryKey" gorm:"autoIncrement" json:"id"`
	Name      string `json:"name"`
	OrderCode string `json:"order_code"`
	Price     int    `json:"price"`
	Type      string `json:"type"`
}
