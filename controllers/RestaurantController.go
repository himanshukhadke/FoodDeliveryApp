package controllers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/himanshukhadke/Zomato/entities"
)

// UserController is a controller for user related operations

var  resaturantController *ResaturantController = nil

type ResaturantController struct {
	list []*entities.Restaurant
}

// NewUserController creates a new user controller
func init() {
	if resaturantController == nil {
		fmt.Println("Creating Restaurant controller")
		resaturantController = &ResaturantController{
			list: []*entities.Restaurant{},
		}
	}

}

// Add new ser in usercontroller list 
func (rc *ResaturantController) AddRestaurant(restaurant *entities.Restaurant) {
	rc.list = append(rc.list, restaurant)
}

func RegisterRestaurant(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "POST":
			var restaurant entities.Restaurant
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&restaurant)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// create a new user
			fmt.Println("Registering a new resturant ", restaurant)
			// add user to the list

			resaturantController.AddRestaurant(&restaurant)

			fmt.Fprintf(w, "Registered a Given resturant")
	}
}

func ShowMenu(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			restaurantId := r.URL.Query().Get("restaurantId")
			restaurant := getRestaurantById(restaurantId)
			if restaurant == nil {
				http.Error(w, "Restaurant not found", http.StatusNotFound)
				return
			}
			// return the menu
			menu, err := json.Marshal(restaurant.Menu)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(menu)
	}

}

func getRestaurantById(id string) *entities.Restaurant {
	for _, restaurant := range resaturantController.list {
		if restaurant.RestaurantId == id {
			return restaurant
		}
	}
	return nil
}