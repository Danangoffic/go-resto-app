package model

type OrderStatus string
type ProductOrderStatus string

type Order struct {
	UserID        string         `json:"-"`
	ID            string         `gorm:"primaryKey" gorm:"size:255" json:"id"`
	Status        OrderStatus    `json:"status"`
	ProductOrders []ProductOrder `json:"product_orders"`
	ReferenceID   string         `gorm:"unique" json:"reference_id"`
}

type ProductOrder struct {
	ID         int                `gorm:"primaryKey" json:"id"`
	OrderID    string             `gorm:"size:255" json:"order_id"`
	OrderCode  string             `json:"order_code"`
	Quantity   int                `json:"quantity"`
	TotalPrice int64              `json:"total_price"`
	Status     ProductOrderStatus `json:"status"`
}

type OrderMenuProductRequest struct {
	OrderCode string `json:"order_code"`
	Quantity  int    `json:"quantity"`
}

type OrderMenuRequest struct {
	UserID        string                    `json:"-"`
	OrderProducts []OrderMenuProductRequest `json:"order_products"`
	ReferenceID   string                    `json:"reference_id"`
}

type GetOrderDataRequest struct {
	UserID  string `json:"-"`
	OrderID string `json:"order_id"`
}

// constant
const (
	OrderStatusProcessed OrderStatus = "Processed"
	OrderStatusFinished  OrderStatus = "Finished"
	OrderStatusFailed    OrderStatus = "Failed"
)

const (
	ProductOrderStatusPreparing ProductOrderStatus = "Preparing"
	ProductOrderStatusFinished  ProductOrderStatus = "Finished"
)
