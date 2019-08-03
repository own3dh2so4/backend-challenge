package repository

import (
	"encoding/json"
	"prueba_cabify/model"
	"sync"
)


type PromotionRepository interface {
    GetPromotions() []model.Promotions
}

var promotionRepository PromotionRepository
var oncePromotionRepository *sync.Once

func init()  {
	oncePromotionRepository = &sync.Once{}
}

func GetPromotionsRepository() PromotionRepository {
	oncePromotionRepository.Do(
		func() {
			promotionRepository = hardCodePromotionRepository{}
		})
	return promotionRepository
}


type hardCodePromotionRepository struct {
}

func (hardCodePromotionRepository) GetPromotions() []model.Promotions {
	return promotionsPermutations(parsePromotions(getPromotionsJSON()))
}

func getPromotionsJSON() []byte  {
    return []byte(`
{
    "promotions": {
        "promotionItemList": [
            {
                "promotion_id": "qwerty",
                "promotion_type": "PromotionItemList",
                "promotion_name": "2x1 VOUCHER",
                "exclusive": false,
                "needed": {
                    "VOUCHER": 2
                },
                "discounts": {
                    "VOUCHER": [1,0]
                },
                "repeatable": true
            }
        ],
        "promotionSameItem": [
            {
                "promotion_id": "qweqweqwe",
                "promotion_type": "PromotionSameItem",
                "promotion_name": "TSHIRT",
                "exclusive": false,
                "needed": "TSHIRT",
                "min": 3,
                "max": 0,
                "discount": 0.95
            }
        ]
    }
}
    `)
}

type promotions struct {
    Promotions struct{
    	PromotionItemList []model.PromotionItemList `json:"promotionItemList"`
    	PromotionSameItem []model.PromotionSameItem `json:"promotionSameItem"`
	} `json:"promotions"`
}

func parsePromotions(promotionsJSON []byte) []model.Promotion {
	promotions := promotions{}
	err := json.Unmarshal(promotionsJSON, &promotions)
	if err != nil {
		panic("Invalid promotion JSON!")
	}
	var promotionSlice []model.Promotion
	for _, promo := range promotions.Promotions.PromotionItemList {
		promotionSlice = append(promotionSlice, promo)
	}
	for _, promo := range promotions.Promotions.PromotionSameItem {
		promotionSlice = append(promotionSlice, promo)
	}
	return promotionSlice
}

func promotionsPermutations(promotions []model.Promotion) []model.Promotions {
	var promos []model.Promotions
	promosAux := map[string][]model.Promotion{}
	promosAux["combo"] = []model.Promotion{}
	for _, promo := range promotions {
		if promo.IsExclusive() {
			promosAux[promo.GetPromotionName()] = []model.Promotion{promo}
		} else {
			promosAux["combo"] = append(promosAux["combo"], promo)
		}
	}
	for _, promoSlice := range promosAux {
		promos = append(promos, promoSlice)
	}
	return promos
}


