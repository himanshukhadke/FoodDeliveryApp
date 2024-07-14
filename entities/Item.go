package entities

type Item struct {
	ItemId   string  `json:"itemId"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func NewItem(itemId string, name string, price float64, quantity int) *Item {
	return &Item{ItemId: itemId, Name: name, Price: price, Quantity: quantity}
}

func (i *Item) GetItemId() string {
	return i.ItemId
}

func (i *Item) GetName() string {
	return i.Name
}

func (i *Item) GetPrice() float64 {
	return i.Price
}

func (i *Item) GetQuantity() int {
	return i.Quantity
}

func (i *Item) SetQuantity(quantity int) {
	i.Quantity = quantity
}

func (i *Item) SetPrice(price float64) {
	i.Price = price
}

func (i *Item) SetName(name string) {
	i.Name = name
}

func (i *Item) SetItemId(itemId string) {
	i.ItemId = itemId
}

func (i *Item) AddQuantity(quantity int) {
	i.Quantity += quantity
}