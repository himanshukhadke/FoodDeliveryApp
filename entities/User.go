package entities

type User struct {
	Id      string   `json:"id"`
	Profile Profile  `json:"profile"`
	Orders  []Order  `json:"orders"`
	Address Address  `json:"address"`
}

// NewUser creates a new User instance with the provided profile and initializes the orders slice.
func NewUser(id string, userProfile Profile, address Address) *User {
	return &User{
		Id:      id,
		Profile: userProfile,
		Orders:  []Order{},
		Address: address,
	}
}

// GetProfile returns the profile of the user.
func (u *User) GetProfile() Profile {
	return u.Profile
}

// AddOrder adds a new order to the user's orders slice.
func (u *User) AddOrder(order Order) {
	u.Orders = append(u.Orders, order)
}

// GetOrders returns the list of orders associated with the user.
func (u *User) GetOrders() []Order {
	return u.Orders
}

// GetAddress returns the address of the user.
func (u *User) GetAddress() Address {
	return u.Address
}

// SetAddress sets a new address for the user.
func (u *User) SetAddress(address Address) {
	u.Address = address
}