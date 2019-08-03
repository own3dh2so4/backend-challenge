package model

import (
	"testing"
)

func TestCreateBasket(t *testing.T) {
	basket := GetNewBasket()
	if len(basket.GetItems()) != 0 {
		t.Error("New baskets must be empty")
	}
}

func TestCreateTwoBaskets(t *testing.T) {
	basket1 := GetNewBasket()
	basket2 := GetNewBasket()
	if basket1.GetID() == basket2.GetID() {
		t.Errorf("Two diffent baskets can not have the same ID. (basket1 %d, basket2 %d)",
			basket1.GetID(), basket2.GetID())
	}
}

func TestAddItemBasket(t *testing.T) {
	basket := GetNewBasket()
	item1 := &Item{
		Code:       "Item1",
		Name:       "Item1",
		Price:      10,
		FinalPrice: 10,
	}
	item2 := &Item{
		Code:       "Item2",
		Name:       "Item2",
		Price:      15,
		FinalPrice: 15,
	}
	basket.AddItem(item1)
	itemsInBasket := basket.GetItems()
	if len(itemsInBasket) != 1 {
			t.Error("Basket must contains one item")
	}
	assetBasketContainsItemNTimes(basket, item1, 1, t)
	assetBasketContainsItemNTimes(basket, item2, 0, t)

	basket.AddItem(item2)
	itemsInBasket = basket.GetItems()
	if len(itemsInBasket) != 2 {
		t.Error("Basket must contains two item")
	}
	assetBasketContainsItemNTimes(basket, item1, 1, t)
	assetBasketContainsItemNTimes(basket, item2, 1, t)

	basket.AddItem(item2)
	itemsInBasket = basket.GetItems()
	if len(itemsInBasket) != 3 {
		t.Error("Basket must contains three item")
	}
	assetBasketContainsItemNTimes(basket, item1, 1, t)
	assetBasketContainsItemNTimes(basket, item2, 2, t)
}

func assetBasketContainsItemNTimes(b Basket, item *Item, times int, t *testing.T) {
	timesFound := 0
	for _, itemInBasket := range b.GetItems() {
		if item.Code == itemInBasket.Code {
			timesFound++
		}
	}
	if timesFound != times {
		t.Errorf("Item %s found %d times, expected %d times", item.Code, timesFound, times)
	}
}