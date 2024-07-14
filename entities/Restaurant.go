package entities

type Restaurant struct {
	RestaurantId string  `json:"restaurantId"`
	Name         string  `json:"name"`
	Address      Address `json:"address"`
	Phone        string  `json:"phone"`
	Menu         []Item  `json:"menu"`
}

func NewRestaurant(restaurantId string, name string, address Address, phone string, Menu []Item) *Restaurant {
	return &Restaurant{RestaurantId: restaurantId, Name: name, Address: address, Phone: phone, Menu: Menu}
}

func (r *Restaurant) GetRestaurantId() string {
	return r.RestaurantId
}

func (r *Restaurant) GetName() string {
	return r.Name
}

func (r *Restaurant) GetAddress() Address {
	return r.Address
}

func (r *Restaurant) GetPhone() string {
	return r.Phone
}

func (r *Restaurant) GetMenu() []Item {
	return r.Menu
}

func (r *Restaurant) AddItem(item Item) {
	r.Menu = append(r.Menu, item)
}