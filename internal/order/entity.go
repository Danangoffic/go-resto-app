package order

type OrderStatus string
type ProductOrderStatus string

type Order struct {
	ID            string `gorm:"primaryKey"`
	Status        OrderStatus
	ProductOrders []ProductOrder
}

type ProductOrder struct {
	ID         int `gorm:"primaryKey"`
	OrderID    string
	OrderCode  string
	Quantity   int
	TotalPrice int64
	Status     ProductOrderStatus
}

type OrderMenuProductRequest struct {
	OrderCode string
	Quantity  int
}

type OrderMenuRequest struct {
	OrderProducts []OrderMenuProductRequest
}

type GetOrderDataRequest struct {
	OrderID string
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
