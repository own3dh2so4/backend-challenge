package service

import (
	"prueba_cabify/model"
	"prueba_cabify/repository"
	"testing"
)

func TestBasketPromotion_GetPrice(t *testing.T) {
	var tests = []struct{
		input struct{
			basket model.Basket
			items  model.Items
		}
		want float64
		message string
	}{
		{
			input: struct {
				basket model.Basket
				items  model.Items
			}{
				basket: GetBasketPromotion([]model.Promotions{}),
				items:  model.Items{},
			},
			want:    0,
			message: "No promotions and empty items",
		},
		{
			input: struct {
				basket model.Basket
				items  model.Items
			}{
				basket: GetBasketPromotion(repository.GetPromotionsRepository().GetPromotions()),
				items:  model.Items{},
			},
			want:    0,
			message: "Default promotions and empty items",
		},
		{
			input: struct {
				basket model.Basket
				items  model.Items
			}{
				basket: GetBasketPromotion(repository.GetPromotionsRepository().GetPromotions()),
				items:  model.Items{getVoucherItem()},
			},
			want:    5,
			message: "Default promotions with a voucher",
		},
		{
			input: struct {
				basket model.Basket
				items  model.Items
			}{
				basket: GetBasketPromotion(repository.GetPromotionsRepository().GetPromotions()),
				items:  model.Items{getVoucherItem(), getVoucherItem()},
			},
			want:    5,
			message: "Default promotions with two voucher",
		},
		{
			input: struct {
				basket model.Basket
				items  model.Items
			}{
				basket: GetBasketPromotion(repository.GetPromotionsRepository().GetPromotions()),
				items:  model.Items{getVoucherItem(), getVoucherItem(), getVoucherItem()},
			},
			want:    10,
			message: "Default promotions with three voucher",
		},
		{
			input: struct {
				basket model.Basket
				items  model.Items
			}{
				basket: GetBasketPromotion(repository.GetPromotionsRepository().GetPromotions()),
				items:  model.Items{getVoucherItem(), getTShirtItem(), getMugItem()},
			},
			want:    32.5,
			message: "Example 1",
		},
		{
			input: struct {
				basket model.Basket
				items  model.Items
			}{
				basket: GetBasketPromotion(repository.GetPromotionsRepository().GetPromotions()),
				items:  model.Items{getVoucherItem(), getTShirtItem(), getVoucherItem()},
			},
			want:    25,
			message: "Example 2",
		},
		{
			input: struct {
				basket model.Basket
				items  model.Items
			}{
				basket: GetBasketPromotion(repository.GetPromotionsRepository().GetPromotions()),
				items:  model.Items{getTShirtItem(), getTShirtItem(), getTShirtItem(), getVoucherItem(), getTShirtItem()},
			},
			want:    81,
			message: "Example 3",
		},
		{
			input: struct {
				basket model.Basket
				items  model.Items
			}{
				basket: GetBasketPromotion(repository.GetPromotionsRepository().GetPromotions()),
				items:  model.Items{getVoucherItem(), getTShirtItem(), getVoucherItem(), getVoucherItem(),
					getMugItem(), getTShirtItem(), getTShirtItem()},
			},
			want:    74.5,
			message: "Example 4",
		},
	}

	for i, test := range tests {
		test.input.basket.AddItems(test.input.items)
		result := test.input.basket.GetPrice()
		if result != test.want {
			t.Errorf("[Test %d] %s. Expected price '%f' result '%f'", i, test.message, test.want, result)
		}
	}
}

func getVoucherItem() *model.Item {
	return &model.Item{
		Code:       "VOUCHER",
		Name:       "Cabify Voucher",
		Price:      5,
		FinalPrice: 5,
	}
}

func getTShirtItem() *model.Item {
	return &model.Item{
		Code:       "TSHIRT",
		Name:       "Cabify T-Shirt",
		Price:      20,
		FinalPrice: 20,
	}
}

func getMugItem() *model.Item {
	return &model.Item{
		Code:       "MUG",
		Name:       "Cabify Coffee Mug",
		Price:      7.5,
		FinalPrice: 7.5,
	}
}


