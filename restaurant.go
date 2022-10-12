package main

type Restaurant struct {
	RestaurantId      int     `json:"restaurant_id"`
	NameRestaurant    string  `json:"restaurant_name"`
	AddressRestaurant string  `json:"restaurant_address"`
	MenuItems         int     `json:"menu_items"`
	Menu              []Food  `json:"menu"`
	Rating            float32 `json:"rating"`
}
