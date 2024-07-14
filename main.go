package main

import (
	"net/http"
	"github.com/himanshukhadke/Zomato/controllers"
)

func main() {
	http.HandleFunc("/user/registerUser", controllers.RegisterUser)
	http.HandleFunc("/user/myOrders", controllers.MyOrders)
	http.HandleFunc("/rider/registerRider", controllers.RegisterRider)
	http.HandleFunc("/rider/showOrders", controllers.ShowOrders)
	http.HandleFunc("/restaurant/registerRestaurant", controllers.RegisterRestaurant)
	http.HandleFunc("/restaurant/showMenu", controllers.ShowMenu)
	http.HandleFunc("/placeOrder", controllers.PlaceOrder)
 	http.ListenAndServe(":8080", nil)
}