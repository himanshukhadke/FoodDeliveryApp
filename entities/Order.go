package entities

import (
	"github.com/himanshukhadke/Zomato/constants"
)

type Order struct {
	OrderId      string              `json:"orderId"`
	PlacedAt     string              `json:"placedAt"`
	Items        []Item              `json:"items"`
	TotalPrice   float64             `json:"totalPrice"`
	OrderStatus  constants.OrderStatus `json:"status"`
	User         User                `json:"user"`
	OrderedFrom  Restaurant          `json:"orderedFrom"`
	DeliveredAt  Address             `json:"deliveredAt"`
}

// Getters
func (o *Order) GetOrderId() string {
	return o.OrderId
}

func (o *Order) GetPlacedAt() string {
	return o.PlacedAt
}

func (o *Order) GetItems() []Item {
	return o.Items
}

func (o *Order) GetTotalPrice() float64 {
	return o.TotalPrice
}

func (o *Order) GetStatus() constants.OrderStatus {
	return o.OrderStatus
}


func (o *Order) GetOrderedFrom() Restaurant {
	return o.OrderedFrom
}

func (o *Order) GetDeliveredAt() Address {
	return o.DeliveredAt
}

// Setters
func (o *Order) SetOrderId(orderId string) {
	o.OrderId = orderId
}

func (o *Order) SetPlacedAt(placedAt string) {
	o.PlacedAt = placedAt
}

func (o *Order) SetItems(items []Item) {
	o.Items = items
}

func (o *Order) SetTotalPrice(totalPrice float64) {
	o.TotalPrice = totalPrice
}

func (o *Order) SetStatus(status constants.OrderStatus) {
	o.OrderStatus = status
}

func (o *Order) SetOrderedFrom(orderedFrom Restaurant) {
	o.OrderedFrom = orderedFrom
}

func (o *Order) SetDeliveredAt(deliveredAt Address) {
	o.DeliveredAt = deliveredAt
}