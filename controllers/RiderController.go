package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/himanshukhadke/Zomato/constants"
	"github.com/himanshukhadke/Zomato/entities"
)

// riderController is a controller for user related operations

var riderController *RiderController = nil

type RiderController struct {
	list []*entities.Rider
}

// NewRiderController creates a new user controller
func init() {
	if riderController == nil {
		fmt.Println("Creating Rider controller")
		riderController = &RiderController{
			list: []*entities.Rider{},
		}
	}
}

// Add new user in RiderController list 
func (uc *RiderController) AddRider(rider *entities.Rider) {
	uc.list = append(uc.list, rider)
}

// Define the new struct for the combined payload
type RiderPayload struct {
	Profile entities.Profile `json:"profile"`
	Address entities.Address `json:"address"`
}

func RegisterRider(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var riderPayload RiderPayload
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&riderPayload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// create a new rider
		fmt.Println("Registering a new Rider ", riderPayload.Profile)
		// Generate a unique ID for the user
		riderId := generateRiderId()
		rider := entities.NewRider(riderId, riderPayload.Profile, riderPayload.Address)

		// Add user to the list
		riderController.AddRider(rider)

		fmt.Fprintf(w, "Registered a new user with ID: %s", riderId)
	}
}

func ShowOrders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		riderId := r.URL.Query().Get("riderId")
		rider := getOrdersByRiderId(riderId)


		orderJson, err := json.Marshal(rider.GetOrders())

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if orderJson == nil {
			http.Error(w, "Rider not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(orderJson)
	}

}

func getOrdersByRiderId(riderId string) *entities.Rider {
	for _, rider := range riderController.list {
		if rider.GetRiderId() == riderId {
			return rider
		}
	}
	return nil
}

func generateRiderId() string {
	return fmt.Sprintf("RIDER-%d", len(riderController.list)+1)
}

func AssignNearestRider(location entities.Location) *entities.Rider {
	var nearestRider *entities.Rider
	var minDistance float64 = 999999
	for _, rider := range riderController.list {
		distance := location.CalculateDistance(rider.GetLocation())
		if distance < minDistance && rider.GetRiderStatus() == constants.IDLE {
			minDistance = distance
			nearestRider = rider
		}
	}
	if nearestRider != nil {
		return nearestRider
	}
	return nil
}