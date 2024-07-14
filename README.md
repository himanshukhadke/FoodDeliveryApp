# Zomato
Food Delivery Application 

Going to Write Go Code

compile using : go build .
run using     : ./Zomato


Creating a Restaurant 
curl --location 'localhost:8080/restaurant/registerRestaurant' \
--header 'Content-Type: application/json' \
--data '{
  "restaurantId": "R001",
  "name": "Chitale Bandhu",
  "address": {
        "location": [40.3, 70.4],
        "street": "Balewadi High Street",
        "pincode": 12345
    },
  "phone": "123-456-7890",
  "menu": [
    {
      "itemId": "R001001",
      "name": "Bhakarwadi",
      "price": 10.99,
      "quantity": 1
    },
    {
      "itemId": "R001002",
      "name": "Kaju Kadai",
      "price": 12.99,
      "quantity": 1
    }
  ]
}'

// Creating a rider

curl --location 'localhost:8080/rider/registerRider' \
--header 'Content-Type: application/json' \
--data-raw '{
  "profile": {
    "username": "Himanshu Khadke",
    "email": "khadkeh754@gmail.com",
    "password": "1234"
  },
  "address": {
    "location": [9.0, 10.0],
    "street": "Kumar Shantiniketan",
    "pincode": 411007
  }
}'

// Registering a User
curl --location 'localhost:8080/user/registerUser' \
--header 'Content-Type: application/json' \
--data-raw '{
  "profile": {
    "username": "Jay Tendolker",
    "email": "jay@gmail.com",
    "password": "2344"
  },
  "address": {
    "location": [2.0, 4.0],
    "street": "Kumar Classic",
    "pincode": 411007
  }
}'

// Placing an Order
curl --location 'localhost:8080/placeOrder' \
--header 'Content-Type: application/json' \
--data '{
  "items": [
    {
      "itemId": "R001001",
      "quantity": 1
    },
    {
      "itemId": "R001002",
      "quantity": 2
    }
  ],
  "totalPrice": 21.97,
  "user": "U1",
  "restaurantId": "R001"
}'

// Get Orders placed by User
curl --location 'localhost:8080/user/myOrders?userId=U1'

// Get Menu of restaurant
curl --location 'localhost:8080/restaurant/showMenu?restaurantId=R001'

// Get Orders of Rider
curl --location 'localhost:8080/rider/showOrders?riderId=RIDER-1'