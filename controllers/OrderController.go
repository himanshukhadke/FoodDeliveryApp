package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/himanshukhadke/Zomato/constants"
	"github.com/himanshukhadke/Zomato/entities"
)

   var orderController *OrderController = nil

   type OrderController struct {
       orders []entities.Order
   }

   func init() {
       if orderController == nil {
           orderController = &OrderController{
               orders: []entities.Order{},
           }
       }
   }

   // PlaceOrder handles the order placement
func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var orderPayload struct {
			Items        []entities.Item `json:"items"`
			TotalPrice   float64         `json:"totalPrice"`
			UserId       string          `json:"user"`
			RestaurantId string          `json:"restaurantId"`
		}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&orderPayload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Fetch user and restaurant information
		user := getUserById(orderPayload.UserId)
		if user == nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		restaurant := getRestaurantById(orderPayload.RestaurantId)
		if restaurant == nil {
			http.Error(w, "Restaurant not found", http.StatusNotFound)
			return
		}

		// Assign the nearest available rider
		rider := AssignNearestRider(user.Address.Location)
		if rider == nil {
			http.Error(w, "No available rider", http.StatusInternalServerError)
			return
		}

		// Calculate delivery time based on distance
		distance := user.Address.Location.CalculateDistance(restaurant.Address.Location)  // ditance of user and restaurant
		distance += restaurant.Address.Location.CalculateDistance(rider.GetLocation()) // distance of restaurant and rider
		deliveryTime := time.Duration(distance/30*60) * time.Minute // assuming average speed of 30 km/h

		// Create and log the order
		order := entities.Order{
			OrderId:     generateOrderId(),
			PlacedAt:    time.Now().Format(time.RFC3339),
			Items:       orderPayload.Items,
			TotalPrice:  orderPayload.TotalPrice,
			OrderStatus: constants.OrderPlaced,
			User:        *user,
			OrderedFrom: *restaurant,
			DeliveredAt: user.Address,
		}
		rider.AddOrder(order)
		rider.SetRiderStatus(constants.BUSY)
		user.AddOrder(order)

		// Return rider information and delivery time
		response := map[string]interface{}{
			"rider": map[string]interface{}{
				"id":       rider.GetRiderId(),
				"name":     rider.GetProfile().Username,
				"location": rider.GetLocation(),
				"status":   rider.GetRiderStatus(),
			},
			"deliveryTime": deliveryTime.String() + " minutes",
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}
}

func generateOrderId() string {
	// Implement a function to generate a unique order ID
	return "order" + string(len(orderController.orders)+1)
}