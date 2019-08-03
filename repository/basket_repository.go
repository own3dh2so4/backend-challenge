package repository

import (
	"errors"
	"fmt"
	"prueba_cabify/model"
	"sync"
)

type BasketRepository interface {
	AddBasket(basket model.Basket) error
	GetBasket(basketID int) (model.Basket, error)
	DeleteBasket(basketID int)
}

var basketRepository BasketRepository
var onceBasketRepository *sync.Once

func init() {
	onceBasketRepository = &sync.Once{}
}

func GetBasketRepository() BasketRepository {
	onceBasketRepository.Do(
		func() {
			basketRepository = &inMemoryBasketRepository{
				mutex:   &sync.RWMutex{},
				baskets: map[int]model.Basket{},
			}
		})
	return basketRepository
}

type inMemoryBasketRepository struct {
	mutex *sync.RWMutex
	baskets map[int]model.Basket
}

func (mem *inMemoryBasketRepository) AddBasket(basket model.Basket) error {
	basketID := basket.GetID()
	mem.mutex.RLock()
	if mem.baskets[basketID] != nil {
		mem.mutex.RUnlock()
		return errors.New(fmt.Sprintf("basket '%d' already exists", basketID))
	}
	mem.mutex.RUnlock()

	mem.mutex.Lock()
	mem.baskets[basketID] = basket
	mem.mutex.Unlock()

	return nil
}

func (mem *inMemoryBasketRepository) DeleteBasket(basketID int) {
	mem.mutex.RLock()
	basket := mem.baskets[basketID]
	mem.mutex.RUnlock()
	if basket != nil {
		mem.mutex.Lock()
		delete(mem.baskets, basketID)
		mem.mutex.Unlock()
	}
}

func (mem *inMemoryBasketRepository) GetBasket(basketID int) (model.Basket, error) {
	mem.mutex.RLock()
	defer mem.mutex.RUnlock()
	basket := mem.baskets[basketID]
	if basket != nil {
		return basket, nil
	}
	return nil, errors.New(fmt.Sprintf("basket '%d' does not exists", basketID))
}