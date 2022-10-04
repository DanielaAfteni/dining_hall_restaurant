package main

import (
	"math/rand"
)

type Order struct {
	OrderId            int     `json:"order_id"`
	Priority           int     `json:"priority"`
	MenuItemIds        []int   `json:"items"`
	MaxPreparationTime float32 `json:"max_wait"`
}

type FoodCookedByCook struct {
	FoodId int `json:"food_id"`
	CookId int `json:"cook_id"`
}

type OrderPrepared struct {
	Order
	WaiterId       int                `json:"waiter_id"`
	TableId        int                `json:"table_id"`
	PickUpTime     int64              `json:"pick_up_time"`
	CookingTime    int64              `json:"cooking_time"`
	CookingDetails []FoodCookedByCook `json:"cooking_details"`
}

func CreateOrder() Order {
	var quantityFoods int = rand.Intn(10) + 1
	var menuItemIds []int
	var maxPrepTime float32
	var prior int
	for i := 0; i <= quantityFoods; i++ {
		menuItemIds = append(menuItemIds, rand.Intn(10)+1)
	}
	//prior = (10 - len(menuItemIds)) / (10 /5)
	//prior = (10 - len(menuItemIds)) / (10 /5) + 1
	prior = (10-quantityFoods)/(10/5) + 1
	maxPrepTime = GetMaxPrepTime(menuItemIds, quantityFoods)
	return Order{
		OrderId:            rand.Intn(10000) + 1,
		MenuItemIds:        menuItemIds,
		MaxPreparationTime: maxPrepTime,
		Priority:           prior,
	}
}

func GetMaxPrepTime(foodsId []int, quantityFoods int) float32 {
	var maxPrepTimeFoods []float32
	for i := 0; i <= quantityFoods; i++ {
		maxPrepTimeFoods = append(maxPrepTimeFoods, Menu[foodsId[i]-1].preparationTime)
	}
	maxPreparationTimeFood := maxPrepTimeFoods[0]
	for _, eachPreparationTimeFood := range maxPrepTimeFoods {
		if eachPreparationTimeFood > maxPreparationTimeFood {
			maxPreparationTimeFood = eachPreparationTimeFood
		}
	}
	maxPreparationTimeFood *= 1.3
	return maxPreparationTimeFood
}
