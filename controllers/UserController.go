package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/himanshukhadke/Zomato/entities"
)

// UserController is a controller for user-related operations
var userController *UserController

type UserController struct {
	list []*entities.User
}

// init initializes the UserController
func init() {
	if userController == nil {
		fmt.Println("Creating User controller")
		userController = &UserController{
			list: []*entities.User{},
		}
	}
}

// AddUser adds a new user to the UserController's list
func (uc *UserController) AddUser(user *entities.User) {
	uc.list = append(uc.list, user)
}

// RegisterUser handles the registration of a new user
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var userPayload struct {
			Profile entities.Profile `json:"profile"`
			Address entities.Address `json:"address"`
		}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&userPayload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Create a new user
		fmt.Println("Registering a new user", userPayload.Profile)
		// Generate a unique ID for the user
		userId := generateUserId()
		user := entities.NewUser(userId, userPayload.Profile, userPayload.Address)

		// Add user to the list
		userController.AddUser(user)

		fmt.Fprintf(w, "Registered a new user with ID: %s", userId)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func MyOrders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		userId := r.URL.Query().Get("userId")
		user := getUserById(userId)
		if user == nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		orders := user.GetOrders()
		ordersJson, err := json.Marshal(orders)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(ordersJson)
		default:
	}
}

// generateUserId generates a unique user ID
func generateUserId() string {
	// Implement a unique ID generation logic here
	return fmt.Sprintf("U%d", len(userController.list)+1)
}

func getUserById(userId string) *entities.User {
	// Implement a logic to get user by ID search in userController list
	for _, user := range userController.list {
		if user.Id == userId {
			return user
		}
	}
	return nil
}