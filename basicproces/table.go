// Packages are used to organize related Go source files together into a single unit
// The current go program is in package basicproces
package basicproces

import (
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/rs/zerolog/log"
)

// A structure or struct in Golang is a user-defined type that allows to group/combine items of possibly different types into a single type.
// Any real-world entity which has some set of properties/fields can be represented as a struct.
// This concept is generally compared with the classes in object-oriented programming.

// Table - has the majority of the properties/fields as channels
// (because he works with channel of sending, with channel of receiving and rating)
// and which are taken from the json file
type Table struct {
	Id           int
	Menu         Menu
	State        string
	CurrentOrder Order
	SendChan     chan<- Order
	ReceiveChan  chan Order
	RatingChan   chan<- int
}

const (
	available = "available"
	ready     = "ready"
	waiting   = "waiting"
)

var orderId int64

// function to return all the tables with their information, data
func NewTable(id int, menu Menu, orderChan chan<- Order, ratingChan chan<- int) *Table {
	return &Table{
		Id:          id,
		Menu:        menu,
		State:       available,
		SendChan:    orderChan,
		ReceiveChan: make(chan Order),
		RatingChan:  ratingChan,
	}
}

// The main idea about tables: is that it represent a person which can:
// wait, send an order and receive it back
func (t *Table) Run() {
	// infinite loops
	for {
		t.waitAvailable()
		t.sendOrder()
		t.receiveOrder()
	}
}

// function which defines the next step of people:
// from: available table -> send, send -> waiting, waiting -> available table
func (t *Table) nextState() {
	// available table -> send
	if t.State == available {
		t.State = ready
	} else if t.State == ready {
		// send -> waiting
		t.State = waiting
	} else if t.State == waiting {
		// waiting -> available table
		t.State = available
	}
}

// function for occupying a table
func (t *Table) waitAvailable() {
	// in case if the table is not then we return
	if t.State != available {
		return
	}
	availableTime := time.Duration(scfg.TimeUnit*(rand.Intn(scfg.MaxTableAvailableTime)+1)) * time.Millisecond
	// we sleep during this time, (like people are staying at the table)
	time.Sleep(availableTime)
	// then we move on to the next stage from available table to ready to send
	t.nextState()

	log.Info().Int("tableId", t.Id).Msg("A table is occupied")
}

// function for a table placing a new order
func (t *Table) sendOrder() {
	// in case if the table is not then we return
	if t.State != ready {
		return
	}
	foodCount := rand.Intn(scfg.MaxOrderItemsCount) + 1

	order := Order{
		OrderId:  atomic.AddInt64(&orderId, 1),
		TableId:  t.Id,
		Items:    make([]int, foodCount),
		Priority: scfg.MaxOrderItemsCount - foodCount,
	}

	maxTime := 0
	for i := 0; i < foodCount; i++ {
		order.Items[i] = rand.Intn(t.Menu.FoodsCount) + 1
		prepTime := t.Menu.Foods[i].PreparationTime
		if prepTime > maxTime {
			maxTime = prepTime
		}
	}
	order.MaxWait = float64(maxTime) * scfg.MaxWaitTimeCoefficient
	t.CurrentOrder = order
	t.SendChan <- order
	t.nextState()

	log.Info().Int("tableId", t.Id).Int64("orderId", order.OrderId).Msg("A table placed new order")
}

// function for a table receiving an order
func (t *Table) receiveOrder() {
	if t.State != waiting {
		return
	}
	// we are looping over elements in received channel
	for order := range t.ReceiveChan {
		// if everything is wrong, then show that: received wrong order
		if order.TableId != t.Id || order.OrderId != t.CurrentOrder.OrderId {
			log.Err(nil).Int("tableId", t.Id).Int64("orderId", order.OrderId).Msg("A table received wrong order")
			continue
		}
		// call the function responsible for the rating calculation
		rating := order.CalculateRating()
		// we send the data of rating to rating channel
		t.RatingChan <- rating
		// moving to another stage: give the people their order back
		t.nextState()
		// show that a table received its order
		log.Info().Int("tableId", t.Id).Int64("orderId", order.OrderId).Int("rating", rating).Msg("A table received its order")
		return
	}
}
