package model

type Promotion interface {
	CanBeApplied(items map[string]Items) bool
	Apply(items map[string]Items) map[string]Items
	IsRepeatable() bool
	IsExclusive() bool
	GetPromotionName() string
}

type NoPromotion struct {
}

func (NoPromotion) IsRepeatable() bool {
	return false
}

func (NoPromotion) IsExclusive() bool {
	return true
}

func (NoPromotion) GetPromotionName() string {
	return "no promo"
}

func (NoPromotion) CanBeApplied(items map[string]Items) bool {
	return true
}

func (NoPromotion) Apply(items map[string]Items) map[string]Items {
	return items
}

type PromotionBasic struct {
	PromotionID string `json:"promotion_id"`
	PromotionType string `json:"promotion_type"`
	PromotionName string `json:"promotion_name"`
	Exclusive bool `json:"exclusive"`
}

func (pm PromotionBasic) IsRepeatable() bool {
	return false
}

func (pm PromotionBasic) IsExclusive() bool {
	return pm.Exclusive
}

func (pm PromotionBasic) GetPromotionName() string {
	return pm.PromotionName
}

func (pm PromotionBasic) CanBeApplied(items map[string]Items) bool {
	return false
}

func (pm PromotionBasic) Apply(items map[string]Items) map[string]Items {
	return items
}

type PromotionItemList struct {
	PromotionBasic
	Needed map[string]int  `json:"needed"`
	Discounts map[string][]float64 `json:"discounts"`
	Repeatable bool `json:"repeatable"`
}

func (pil PromotionItemList) CanBeApplied(items map[string]Items) bool {
	for itemNeedID, num := range pil.Needed {
		if items[itemNeedID] == nil || len(items[itemNeedID]) < num {
			return false
		}
	}
	return true
}

func (pil PromotionItemList) Apply(items map[string]Items) map[string]Items {
	if !pil.CanBeApplied(items) {
		return items
	}
	itemsResult := map[string]Items{}
	for itemID, itemSlice := range items {
		if pil.Needed[itemID] != 0 {
			newItemKey := itemID + "_" + pil.PromotionID
			num := pil.Needed[itemID]
			itemsResult[newItemKey] = itemSlice[:num]
			if len(itemSlice) > num {
				itemsResult[itemID] = itemSlice[num:]
			}
			for i, item := range itemsResult[newItemKey] {
				item.FinalPrice = item.FinalPrice * pil.Discounts[itemID][i]
			}
		} else {
			itemsResult[itemID] = itemSlice
		}
	}
	return itemsResult
}

func (pil PromotionItemList) IsRepeatable() bool {
	return pil.Repeatable
}

type PromotionSameItem struct {
	PromotionBasic
	Needed string  `json:"needed"`
	Min int `json:"min"`
	Max int `json:"max"`
	Discount float64 `json:"discount"`
}

func (psi PromotionSameItem) CanBeApplied(items map[string]Items) bool {
	return items[psi.Needed] != nil && len(items[psi.Needed]) >= psi.Min
}

func (psi PromotionSameItem) Apply(items map[string]Items) map[string]Items {
	if !psi.CanBeApplied(items) {
		return items
	}
	itemsResult := map[string]Items{}
	numItems := len(items[psi.Needed])
	newItemKey := psi.Needed + "_" + psi.PromotionID
	if psi.Max != 0 && numItems > psi.Max {
		itemsResult[newItemKey] = items[psi.Needed][:psi.Max]
		itemsResult[newItemKey] = items[psi.Needed][psi.Max:]
	} else {
		itemsResult[newItemKey] = items[psi.Needed]
	}
	for _, item := range itemsResult[newItemKey] {
		item.FinalPrice = item.FinalPrice * psi.Discount
	}
	for itemID, itemSlice := range items {
		if itemID != psi.Needed {
			itemsResult[itemID] = itemSlice
		}
	}

	return itemsResult
}

type Promotions []Promotion

func (p Promotions) Apply(items map[string]Items) (float64, map[string]Items) {
	for _, promo := range p {
		items = promo.Apply(items)
	}
	price := 0.0
	for _, itemSlice := range items {
		for _, item := range itemSlice {
			price = price + item.FinalPrice
		}
	}

	return price, items
}
