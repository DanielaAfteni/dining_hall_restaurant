package main

import (
	"sync"
	"time"
)

const NrTables int = 10

var tablesMutex sync.Mutex

type State int

type Table struct {
	Table_Id   int
	stateMutex sync.Mutex
	state      State
	orderMutex sync.Mutex
	order      Order
}

var tables []*Table

const (
	TableAvailable                  State = 1
	TableWaitingToMakeOrder         State = 2
	TableWaitingForOrderPreparation State = 3
)

func generateOrders() {
	for {
		for tableId := range tables {
			eachTable := GetTable(tableId)
			if eachTable.GetState() == TableAvailable {
				eachTable.SetState(TableWaitingToMakeOrder)
			}
		}
		time.Sleep(2 * TIME_UNIT * time.Millisecond)
	}
}

func GetTable(tableId int) *Table {
	tablesMutex.Lock()
	defer tablesMutex.Unlock()
	return tables[tableId]
}

func (t *Table) OrderCreationASForTable() Order {
	t.SetState(TableWaitingForOrderPreparation)
	t.SetOrder(CreateOrder())
	time.Sleep(2 * TIME_UNIT * time.Millisecond)

	t.GetLockUnlockOrder()
	order := t.order

	return order
}

func (t *Table) GetLockUnlockState() {
	t.stateMutex.Lock()
	defer t.stateMutex.Unlock()
}

func (t *Table) GetLockUnlockOrder() {
	t.orderMutex.Lock()
	defer t.orderMutex.Unlock()
}

func (t *Table) GetState() State {
	t.GetLockUnlockState()
	return t.state
}

func (t *Table) SetState(newState State) {
	t.GetLockUnlockState()
	t.state = newState
}

func (t *Table) SetOrder(newOrder Order) {
	t.GetLockUnlockOrder()
	t.order = newOrder
}
