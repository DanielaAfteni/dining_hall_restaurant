package main

type Waiter struct {
	id           int
	waiterTables []*Table
}

const NrWaiters int = 4

var Waiters [NrWaiters]*Waiter

type OrderToSend struct {
	Order
	Table_Id   int   `json:"table_id"`
	WaiterId   int   `json:"waiter_id"`
	PickUpTime int64 `json:"pick_up_time"`
}
