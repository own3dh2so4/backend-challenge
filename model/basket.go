package model

import "sync"

type Basket interface {
	GetID() int
	AddItem(item *Item)
	AddItems(items Items)
	GetItems() Items
	GetPrice() float64
}

func GetNewBasket() Basket{
	return &basket{
		ID:    getNextBasketID(),
		items: Items{},
	}
}

var basketID = 0
var mutexBasketID = &sync.Mutex{}

func getNextBasketID() int {
	mutexBasketID.Lock()
	defer mutexBasketID.Unlock()
	basketID++
	return basketID
}

type basket struct {
	ID int
	items Items
}

func (b *basket) AddItem(item *Item)  {
	b.items = append(b.items, item)
}

func (b *basket) AddItems(items Items)  {
	b.items = append(b.items, items...)
}

func (b basket) GetItems() Items{
	return b.items
}

func (b basket) GetPrice() float64 {
	return b.items.getFinalPrice()
}

func (b basket) GetID() int {
	return b.ID
}