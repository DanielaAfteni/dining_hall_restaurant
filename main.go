package main

import (
	// Package json implements encoding and decoding of JSON.
	// The mapping between JSON and Go values is described in the documentation for the Marshal and Unmarshal functions.

	// Package ioutil implements some I/O utility functions.

	// Package os provides a platform-independent interface to operating system functionality.

	// importing the gin, because is a high-performance HTTP web framework written in Golang (Go).
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const TIME_UNIT = 250

func main() {
	router := gin.Default()
	router.POST("/distribution", recieveOrder)
	router.POST("/v2/order", recieveOrderFromFoodOrdering)

	rand.Seed(time.Now().UnixNano())
	for id := 0; id < NrTables; id++ {
		tables = append(tables, &Table{
			Table_Id: id + 1,
			state:    TableAvailable,
		})
	}
	for id := range Waiters {
		Waiters[id] = &Waiter{
			id:           id + 1,
			waiterTables: tables,
		}
	}
	go generateOrders()
	for id := range Waiters {
		go Waiters[id].OrdersToLookFor()
	}
	go RestaurantRegistration()

	router.Run(":8080")
}

func RestaurantRegistration() {
	var restaurant = Restaurant{
		RestaurantId:      1,
		NameRestaurant:    "McDonald's",
		AddressRestaurant: "http://dining_hall_restaurant:8080",
		MenuItems:         10,
		Menu:              Menu,
	}
	jsonBody, err := json.Marshal(restaurant)
	if err != nil {
		log.Err(err).Msg("Error!!!")
	}
	contentType := "application/json"
	//_, err = http.Post("http://food_ordering_service_restaurant:8090/register", contentType, bytes.NewReader(jsonBody))
	_, err = http.Post("http://localhost:8090/register", contentType, bytes.NewReader(jsonBody))
	if err != nil {
		log.Err(err).Msg("Error!!")
	}
}

func recieveOrder(c *gin.Context) {
	var order *OrderPrepared
	if err := c.BindJSON(&order); err != nil {
		return
	}
	GetRating(order.MaxPreparationTime, order.CookingTime)
	GetTable(order.TableId - 1).SetState(TableAvailable)
	log.Printf("Already prepared order was recieved from kitchen %+v \n", order)
	c.IndentedJSON(http.StatusCreated, order)
}

func recieveOrderFromFoodOrdering(c *gin.Context) {
	var order *Order
	if err := c.BindJSON(&order); err != nil {
		log.Err(err).Msg("Error!!")
		return
	}
	log.Printf("The restaurant received from client new order: %+v", order)
	c.IndentedJSON(http.StatusCreated, order)
}

func (w *Waiter) OrdersToLookFor() {
	for {
		for _, eachtable := range w.waiterTables {
			if eachtable.GetState() == TableWaitingToMakeOrder {
				order := eachtable.OrderCreationASForTable()
				log.Printf("From table with id = %d, the waiter with id = %d took the order %+v", eachtable.Table_Id, w.id, order)
				sendOrder := &OrderToSend{
					Order:      order,
					Table_Id:   eachtable.Table_Id,
					WaiterId:   w.id,
					PickUpTime: time.Now().Unix(),
				}
				jsonBody, err := json.Marshal(sendOrder)
				if err != nil {
					log.Err(err).Msg("Error!!!")
				}
				contentType := "application/json"
				//_, err = http.Post("http://kitchen_restaurant:8081/order", contentType, bytes.NewReader(jsonBody))
				_, err = http.Post("http://localhost:8081/order", contentType, bytes.NewReader(jsonBody))
				if err != nil {
					log.Err(err).Msg("Error!!")
				}
			}
		}
	}
}

func CalculateFinalRating() {
	var totalPointsGathered float32
	var nrOfTotalOrders float32
	for _, points := range Ratings {
		nrOfTotalOrders += points
	}
	totalPointsGathered = Ratings[1]*1 + Ratings[2]*2 + Ratings[3]*3 + Ratings[4]*4 + Ratings[5]*5
	finalRating := totalPointsGathered / nrOfTotalOrders
	log.Printf("Final rating of the restaurant = %f", finalRating)
}
