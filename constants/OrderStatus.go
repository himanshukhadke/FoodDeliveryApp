package constants

type OrderStatus string

const (
	OrderPlaced    OrderStatus = "Order Placed"
	OrderConfirmed OrderStatus = "Order Confirmed"
	OrderCancelled OrderStatus = "Order Cancelled"
	OrderDelivered OrderStatus = "Order Delivered"
)