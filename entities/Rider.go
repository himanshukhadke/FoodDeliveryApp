package entities

import (
	"github.com/himanshukhadke/Zomato/constants"
)

type Rider struct {
	riderId string								`json:"riderId"`
	profile Profile								`json:"profile"`
	ordersLog []Order							`json:"ordersLog"`
	riderStatus constants.RiderStatus			`json:"riderStatus"`
	Address Address								`json:"address"`
}

func NewRider(riderId string,riderProfile Profile, riderAddress Address) *Rider {
	return &Rider{
		riderId: riderId,
		profile: riderProfile,
		ordersLog: []Order{},
		riderStatus: constants.IDLE,
		Address: riderAddress,
	}
}

func (r *Rider) GetRiderId() string {
	return r.riderId
}

func (r *Rider) GetProfile() Profile {
	return r.profile
}

func (r *Rider) AddOrder(order Order) {
	r.ordersLog = append(r.ordersLog, order)
}

func (r *Rider) GetRiderStatus() constants.RiderStatus {
	return r.riderStatus
}

func (r *Rider) SetRiderStatus(status constants.RiderStatus) {
	r.riderStatus = status
}

func (r *Rider) GetLocation() Location {
	return r.Address.Location
}

func (r *Rider) GetOrders() []Order {
	return r.ordersLog
}