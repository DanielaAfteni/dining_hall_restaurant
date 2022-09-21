// The current go program is in package basicproces
package basicproces

import (
	// Package time provides functionality for measuring and displaying time.
	"time"

	"github.com/rs/zerolog/log"
)

// A structure or struct in Golang is a user-defined type that allows to group/combine items of possibly different types into a single type.
// Any real-world entity which has some set of properties/fields can be represented as a struct.
// This concept is generally compared with the classes in object-oriented programming.

// Order - has the majority of the properties/fields as int (because are integer numbers)
// and which are taken from the json file
type Order struct {
	// This Order struct will hold all order details variables of the application that we read from file
	OrderId    int64   `json:"order_id"`
	TableId    int     `json:"table_id"`
	WaiterId   int     `json:"waiter_id"`
	Items      []int   `json:"items"`
	Priority   int     `json:"priority"`
	MaxWait    float64 `json:"max_wait"`
	PickUpTime int64   `json:"pick_up_time"`
}

// function to calculate the rating
func (o Order) CalculateRating() int {
	// calculate the order time
	orderTime := float64((time.Now().Unix() - o.PickUpTime) * 1000 / int64(scfg.TimeUnit))
	// show the maximum waiting time for this order
	maxWaitTime := o.MaxWait
	// show the corresponding message
	log.Info().Int64("orderId", o.OrderId).Float64("orderTime", orderTime).Float64("maxWait", maxWaitTime).Msg("Rating calculation")
	// in case if the order time is smaller than maximum waiting time
	if orderTime < maxWaitTime {
		// then return 5
		return 5
	}
	// in case if the order time is smaller than maximum waiting time multiplied by 1.1
	if orderTime < maxWaitTime*1.1 {
		// then return 4
		return 4
	}
	// in case if the order time is smaller than maximum waiting time multiplied by 1.2
	if orderTime < maxWaitTime*1.2 {
		// then return 3
		return 3
	}
	// in case if the order time is smaller than maximum waiting time multiplied by 1.3
	if orderTime < maxWaitTime*1.3 {
		// then return 2
		return 2
	}
	// in case if the order time is smaller than maximum waiting time multiplied by 1.4
	if orderTime < maxWaitTime*1.4 {
		// then return 1
		return 1
	}
	// otherwise return just raiting 0
	return 0
}
