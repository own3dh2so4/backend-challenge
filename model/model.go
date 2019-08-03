package model

type Item struct{
	Code string
	Name string
	Price float64
	FinalPrice float64
}

type Items []*Item

func (items Items) getFinalPrice() float64{
	price := 0.0
	for _, item := range items {
		price = price + item.FinalPrice
	}
	return price
}