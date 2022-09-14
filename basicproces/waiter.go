// The current go program is in package basicproces
package basicproces

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"

	//Package time provides functionality for measuring and displaying time.
	"time"

	"github.com/rs/zerolog/log"
)

// A structure or struct in Golang is a user-defined type that allows to group/combine items of possibly different types into a single type.
// Any real-world entity which has some set of properties/fields can be represented as a struct.
// This concept is generally compared with the classes in object-oriented programming.

// Waiter - has the majority of the properties/fields as channels
// (because he works with channels of tables, with channel of order and distribution)
// and which are taken from the json file
type Waiter struct {
	// This Order struct will hold all waiter variables of the application
	Id               int
	CurrentOrder     Order
	DistributionChan chan Distribution
	OrderChan        <-chan Order
	TablesChans      []chan Order
}

// function to return all the waiters with their information, data
func NewWaiter(id int, orderChan <-chan Order, tablesChans []chan Order) *Waiter {
	return &Waiter{
		Id:               id,
		DistributionChan: make(chan Distribution),
		OrderChan:        orderChan,
		TablesChans:      tablesChans,
	}
}

// function for running a waiter
func (w *Waiter) Run() {
	// infinite loop
	for {
		select {
		// for sending order to kitchen by waiter
		case order := <-w.OrderChan:
			pickupTime := time.Duration(scfg.TimeUnit*(rand.Intn(scfg.MaxPickupTime)+1)) * time.Millisecond
			// sleep a little bit at picking up the order (realism)
			time.Sleep(pickupTime)
			// set the time of picking up the order
			order.PickUpTime = time.Now().Unix()
			order.WaiterId = w.Id
			// use the Marshal() function in package encoding/json to pack or code the data from JSON to a struct
			jsonBody, err := json.Marshal(order)
			if err != nil {
				log.Fatal().Err(err).Msg("Error marshalling order")
			}
			contentType := "application/json"

			_, err = http.Post(scfg.KitchenUrl+"/order", contentType, bytes.NewReader(jsonBody))
			if err != nil {
				log.Fatal().Err(err).Msg("Error sending order to kitchen")
			}

			log.Info().Int("waiter_id", w.Id).Int64("order_id", order.OrderId).Msg("A waiter sent order to kitchen")
		// for distributing waiter
		case distribution := <-w.DistributionChan:
			order := distribution.Order
			log.Info().Int("waiter_id", w.Id).Int64("order_id", order.OrderId).Int("cooking_time", distribution.CookingTime).Float64("max_wait", distribution.MaxWait).Msgf("Waiter received distribution")
			w.TablesChans[order.TableId] <- order
		}
	}
}
