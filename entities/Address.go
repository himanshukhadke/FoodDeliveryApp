package entities

type Address struct {
	Location Location `json:"location"`
	Street   string   `json:"street"`
	Pincode  int      `json:"pincode"`
}

func NewAddress(location Location, street string, pincode int) *Address {
	return &Address{Location: location, Street: street, Pincode: pincode}
}

func (a *Address) GetLocation() Location {
	return a.Location
}

func (a *Address) GetStreet() string {
	return a.Street
}

func (a *Address) GetPincode() int {
	return a.Pincode
}