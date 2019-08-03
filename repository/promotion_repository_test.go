package repository

import (
	"prueba_cabify/model"
	"reflect"
	"testing"
)

func TestGetPromotions(t *testing.T) {
	promoRepo := GetPromotionsRepository()
	promotions := promoRepo.GetPromotions()
	want := []model.Promotions {
		{
			model.PromotionItemList{
				PromotionBasic: model.PromotionBasic{
					PromotionID:   "qwerty",
					PromotionType: "PromotionItemList",
					PromotionName: "2x1 VOUCHER",
					Exclusive:     false,
				},
				Needed: map[string]int{
					"VOUCHER": 2,
				},
				Discounts: map[string][]float64{
					"VOUCHER": {1, 0},
				},
				Repeatable: true,
			},
			model.PromotionSameItem{
				PromotionBasic: model.PromotionBasic{
					PromotionID:   "qweqweqwe",
					PromotionType: "PromotionSameItem",
					PromotionName: "TSHIRT",
					Exclusive:     false,
				},
				Needed:   "TSHIRT",
				Min:      3,
				Max:      0,
				Discount: 0.95,
			},
		},
	}
	if !reflect.DeepEqual(promotions, want) {
		t.Errorf("promotions expected %+v, received %+v", want, promotions)
	}
}