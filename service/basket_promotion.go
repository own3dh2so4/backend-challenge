package service

import (
	"prueba_cabify/model"
	"prueba_cabify/repository"
	"sync"
)

type basketPromotion struct {
	model.Basket
	itemsWithPromo    map[string]model.Items
	priceWithPromo    float64
	promotionsApplied model.Promotions
	validCache        bool
	mutex             *sync.RWMutex
}

func GetBasketPromotion() model.Basket {
	return &basketPromotion{
		Basket:     model.GetNewBasket(),
		validCache: false,
		mutex:      &sync.RWMutex{},
	}
}

func (bp *basketPromotion) prepareBasket(){
	bp.promotionsApplied = model.Promotions{model.NoPromotion{}}
	bp.priceWithPromo = bp.Basket.GetPrice()
	items := map[string]model.Items{}
	for _, item := range bp.Basket.GetItems() {
		if items[item.Code] == nil {
			items[item.Code] = model.Items{}
		}
		items[item.Code] = append(items[item.Code], item)
	}
	bp.itemsWithPromo = items
}


func (bp *basketPromotion) applyBestPromotions() {
	bp.prepareBasket()
	itemsWithOutPromo := bp.itemsWithPromo
	for _, promotions := range repository.GetPromotionsRepository().GetPromotions() {
		price, itemsWithPromo := promotions.Apply(itemsWithOutPromo)
		if price < bp.priceWithPromo {
			bp.priceWithPromo = price
			bp.itemsWithPromo = itemsWithPromo
			bp.promotionsApplied = promotions
		}
	}
	bp.validCache = true
}

func (bp *basketPromotion) AddItem (item *model.Item)  {
	bp.mutex.Lock()
	bp.validCache = false
	bp.Basket.AddItem(item)
	bp.mutex.Unlock()
}

func (bp *basketPromotion) AddItems (items model.Items)  {
	bp.mutex.Lock()
	bp.validCache = false
	bp.Basket.AddItems(items)
	bp.mutex.Unlock()
}

func (bp basketPromotion) GetItems() model.Items {
	bp.mutex.RLock()
	defer bp.mutex.RUnlock()
	return bp.Basket.GetItems()
}

func (bp basketPromotion) GetPrice() float64 {
	bp.mutex.RLock()
	defer bp.mutex.RUnlock()
	if !bp.validCache {
		bp.applyBestPromotions()
	}
	return bp.priceWithPromo
}
