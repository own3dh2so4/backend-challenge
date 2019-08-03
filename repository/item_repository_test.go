package repository

import (
	"prueba_cabify/model"
	"reflect"
	"testing"
)

func TestGetItems(t *testing.T) {
	items := GetItemsRepository().GetItems()
	expectedItems := []*model.Item{
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

	if !reflect.DeepEqual(expectedItems, items) {
		t.Errorf("expected items %+v, given items %+v", expectedItems, items)
	}
}