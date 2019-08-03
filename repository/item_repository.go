package repository

import (
	"errors"
	"fmt"
	"prueba_cabify/model"
	"sync"
)

type ItemsRepository interface {
	GetItems() []*model.Item
	GetItem(ID string) (*model.Item, error)
}

var itemsRepository ItemsRepository
var onceItemsRepository *sync.Once

func init()  {
	onceItemsRepository = &sync.Once{}
}

func GetItemsRepository() ItemsRepository {
	onceItemsRepository.Do(
		func() {
			itemsRepository = hardCodeItemsRepository{}
		})
	return itemsRepository
}

type hardCodeItemsRepository struct {
}

func (hc hardCodeItemsRepository) GetItems() []*model.Item {
	return []*model.Item{
		{
			Code:       "VOUCHER",
			Name:       "Cabify Voucher",
			Price:      5,
			FinalPrice: 5,
		},
		{
			Code:       "TSHIRT",
			Name:       "Cabify T-Shirt",
			Price:      20,
			FinalPrice: 20,
		},
		{
			Code:       "MUG",
			Name:       "Cabify Coffee Mug",
			Price:      7.5,
			FinalPrice: 7.5,
		},
	}
}

func (hc hardCodeItemsRepository) GetItem(ID string) (*model.Item, error) {
	for _, item := range hc.GetItems() {
		if item.Code == ID {
			return item, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("item %s not found", ID))
}
