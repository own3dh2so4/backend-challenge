package model

import (
	"reflect"
	"testing"
)

func TestPromotionBasicIsRepeatable(t *testing.T) {
	promo := PromotionBasic{}
	if promo.IsRepeatable() {
		t.Error("PromotionBasic never is repeatable")
	}
}

func TestPromotionBasicIsExclusive(t *testing.T) {
	var tests = []struct{
		input PromotionBasic
		want bool
		message string
	}{
		{
			input:   PromotionBasic{
				Exclusive:     false,
			},
			want:    false,
			message: "This promotion basic is not exclusive",
		},
		{
			input:   PromotionBasic{
				Exclusive:     true,
			},
			want:    true,
			message: "This promotion basic is exclusive",
		},
	}

	for i, test := range tests {
		if test.input.IsExclusive() != test.want {
			t.Errorf("[Test %d] %s", i, test.message)
		}
	}
}

func TestPromotionBasicCanBeApplied(t *testing.T) {
	promo := PromotionBasic{}
	if promo.CanBeApplied(map[string]Items{}) {
		t.Error("PromotionBasic never is applied")
	}
}

func TestPromotionBasicApply(t *testing.T) {
	var tests = []struct {
		input struct {
			PromotionBasic
			items map[string]Items
		}
		want    map[string]Items
		message string
	} {
		{
			input: struct {
				PromotionBasic
				items map[string]Items
			}{
				PromotionBasic: PromotionBasic{},
				items:          nil,
			},
			want:    nil,
			message: "If nil is given nil is returned",
		},
		{
			input: struct {
				PromotionBasic
				items map[string]Items
			}{
				PromotionBasic: PromotionBasic{},
				items: map[string]Items{},
			},
			want:    map[string]Items{},
			message: "Must return a empty map",
		},
		{
			input: struct {
				PromotionBasic
				items map[string]Items
			}{
				PromotionBasic: PromotionBasic{},
				items: map[string]Items{"aaa" : {&Item{Name:"aaa",}}},
			},
			want:    map[string]Items{"aaa" : {&Item{Name:"aaa",}}},
			message: "Must return the same items map",
		},
	}

	for i, test := range tests {
		result := test.input.PromotionBasic.Apply(test.input.items)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("[Test %d] %s", i, test.message)
		}
	}
}

func TestPromotionItemListIsRepeatable(t *testing.T) {
	var tests = []struct{
		input PromotionItemList
		want bool
		message string
	}{
		{
			input:   PromotionItemList{
				Repeatable:     false,
			},
			want:    false,
			message: "This promotion is not repeatable",
		},
		{
			input:   PromotionItemList{
				Repeatable:     true,
			},
			want:    true,
			message: "This promotion is repeatable",
		},
	}

	for i, test := range tests {
		if test.input.IsRepeatable() != test.want {
			t.Errorf("[Test %d] %s", i, test.message)
		}
	}
}

func getPromotionItemListExample() PromotionItemList {
	return PromotionItemList{
		PromotionBasic: PromotionBasic{
			PromotionID:   "qwerty",
			PromotionType: "PromotionItemList",
			PromotionName: "2x1 VOUCHER",
			Exclusive:     false,
		},
		Needed:         map[string]int{
			"VOUCHER": 2,
		},
		Discounts:      map[string][]float64{
			"VOUCHER": {1, 0},
		},
		Repeatable:     true,
	}
}

func getVoucherItem() *Item {
	return &Item{
		Code:       "VOUCHER",
		Name:       "Cabify Voucher",
		Price:      5,
		FinalPrice: 5,
	}
}

func getVoucherItemDiscount(discount float64) *Item {
	return &Item{
		Code:       "VOUCHER",
		Name:       "Cabify Voucher",
		Price:      5,
		FinalPrice: 5 * discount,
	}
}

func getTShirtItem() *Item {
	return &Item{
		Code:       "TSHIRT",
		Name:       "Cabify T-Shirt",
		Price:      20,
		FinalPrice: 20,
	}
}

func getTShirtItemDiscount(discount float64) *Item {
	return &Item{
		Code:       "TSHIRT",
		Name:       "Cabify T-Shirt",
		Price:      20,
		FinalPrice: 20 * discount,
	}
}

func TestPromotionItemListCanBeApplied(t *testing.T) {
	tests := []struct {
		input struct {
			PromotionItemList
			items map[string]Items
		}
		want    bool
		message string
	}{
		{
			input: struct {
				PromotionItemList
				items map[string]Items
			}{
				PromotionItemList: getPromotionItemListExample(),
				items: map[string]Items{
					"VOUCHER": {getVoucherItem()},
				},
			},
			want: false,
			message: "This promotion is not applicable",
		},
		{
			input: struct {
				PromotionItemList
				items map[string]Items
			}{
				PromotionItemList: getPromotionItemListExample(),
				items:             map[string]Items{
					"VOUCHER": {getVoucherItem(), getVoucherItem()},
				},
			},
			want:    true,
			message: "This promo is applicable",
		},
		{
			input: struct {
				PromotionItemList
				items map[string]Items
			}{
				PromotionItemList: getPromotionItemListExample(),
				items:             map[string]Items{
					"VOUCHER": {getVoucherItem(), getVoucherItem(), getVoucherItem()},
				},
			},
			want:    true,
			message: "This promo is applicable",
		},
		{
			input: struct {
				PromotionItemList
				items map[string]Items
			}{
				PromotionItemList: getPromotionItemListExample(),
				items:             map[string]Items{
					"VOUCHER": {getVoucherItem(), getVoucherItem(), getVoucherItem(), getVoucherItem()},
					"TSHIRT": {getTShirtItem()},
				},
			},
			want:    true,
			message: "This promo is applicable",
		},
	}

	for i, test := range tests {
		result := test.input.PromotionItemList.CanBeApplied(test.input.items)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("[Test %d] %s", i, test.message)
		}
	}
}

func TestPromotionItemListApply(t *testing.T) {
	tests := []struct {
		input struct {
			PromotionItemList
			items map[string]Items
		}
		want    map[string]Items
		message string
	}{
		{
			input: struct {
				PromotionItemList
				items map[string]Items
			}{
				PromotionItemList: getPromotionItemListExample(),
				items: map[string]Items{
					"VOUCHER": {getVoucherItem()},
				},
			},
			want: map[string]Items{
				"VOUCHER": {getVoucherItem()},
			},
			message: "This promotion is not applicable",
		},
		{
			input: struct {
				PromotionItemList
				items map[string]Items
			}{
				PromotionItemList: getPromotionItemListExample(),
				items:             map[string]Items{
					"VOUCHER": {getVoucherItem(), getVoucherItem()},
				},
			},
			want:    map[string]Items{
				"VOUCHER_qwerty": {getVoucherItem(), getVoucherItemDiscount(0)},
			},
			message: "Must return one voucher free",
		},
		{
			input: struct {
				PromotionItemList
				items map[string]Items
			}{
				PromotionItemList: getPromotionItemListExample(),
				items:             map[string]Items{
					"VOUCHER": {getVoucherItem(), getVoucherItem(), getVoucherItem()},
				},
			},
			want:    map[string]Items{
				"VOUCHER_qwerty": {getVoucherItem(), getVoucherItemDiscount(0)},
				"VOUCHER": {getVoucherItem()},
			},
			message: "Must return one voucher free",
		},
		{
			input: struct {
				PromotionItemList
				items map[string]Items
			}{
				PromotionItemList: getPromotionItemListExample(),
				items:             map[string]Items{
					"VOUCHER": {getVoucherItem(), getVoucherItem(), getVoucherItem(), getVoucherItem()},
					"TSHIRT": {getTShirtItem()},
				},
			},
			want:    map[string]Items{
				"VOUCHER_qwerty": {getVoucherItem(), getVoucherItemDiscount(0)},
				"VOUCHER": {getVoucherItem(), getVoucherItem()},
				"TSHIRT": {getTShirtItem()},
			},
			message: "Must return one voucher free",
		},
	}

	for i, test := range tests {
		result := test.input.PromotionItemList.Apply(test.input.items)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("[Test %d] %s", i, test.message)
		}
	}
}

func getPromotionSameItemExample() PromotionSameItem {
	return PromotionSameItem{
		PromotionBasic: PromotionBasic{
			PromotionID:   "qweqweqwe",
			PromotionType: "PromotionSameItem",
			PromotionName: "TSHIRT",
			Exclusive:     false,
		},
		Needed:         "TSHIRT",
		Min:            3,
		Max:            0,
		Discount:       0.95,
	}
}

func TestPromotionSameItemCanBeApplied(t *testing.T) {
	tests := []struct {
		input struct {
			PromotionSameItem
			items map[string]Items
		}
		want    bool
		message string
	}{
		{
			input: struct {
				PromotionSameItem
				items map[string]Items
			}{
				PromotionSameItem: getPromotionSameItemExample(),
				items: map[string]Items{
					"TSHIRT": {getTShirtItem()},
				},
			},
			want: false,
			message: "This promotion is not applicable",
		},
		{
			input: struct {
				PromotionSameItem
				items map[string]Items
			}{
				PromotionSameItem: getPromotionSameItemExample(),
				items:             map[string]Items{
					"TSHIRT": {getTShirtItem()},
				},
			},
			want:    false,
			message: "This promo is not applicable",
		},
		{
			input: struct {
				PromotionSameItem
				items map[string]Items
			}{
				PromotionSameItem: getPromotionSameItemExample(),
				items:             map[string]Items{
					"TSHIRT": {getTShirtItem(), getTShirtItem(), getTShirtItem()},
				},
			},
			want:    true,
			message: "This promo is applicable",
		},
		{
			input: struct {
				PromotionSameItem
				items map[string]Items
			}{
				PromotionSameItem: getPromotionSameItemExample(),
				items:             map[string]Items{
					"TSHIRT": {getTShirtItem(), getTShirtItem(), getTShirtItem(), getTShirtItem()},
				},
			},
			want:    true,
			message: "This promo is applicable",
		},
	}

	for i, test := range tests {
		result := test.input.PromotionSameItem.CanBeApplied(test.input.items)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("[Test %d] %s", i, test.message)
		}
	}
}

func TestPromotionSameItemApply(t *testing.T) {
	tests := []struct {
		input struct {
			PromotionSameItem
			items map[string]Items
		}
		want    map[string]Items
		message string
	}{
		{
			input: struct {
				PromotionSameItem
				items map[string]Items
			}{
				PromotionSameItem: getPromotionSameItemExample(),
				items: map[string]Items{
					"TSHIRT": {getTShirtItem()},
				},
			},
			want: map[string]Items{
				"TSHIRT": {getTShirtItem()},
			},
			message: "This promotion is not applicable",
		},
		{
			input: struct {
				PromotionSameItem
				items map[string]Items
			}{
				PromotionSameItem: getPromotionSameItemExample(),
				items:             map[string]Items{
					"TSHIRT": {getTShirtItem(), getTShirtItem()},
				},
			},
			want:    map[string]Items{
				"TSHIRT": {getTShirtItem(), getTShirtItem()},
			},
			message: "This promo is not applicable",
		},
		{
			input: struct {
				PromotionSameItem
				items map[string]Items
			}{
				PromotionSameItem: getPromotionSameItemExample(),
				items:             map[string]Items{
					"TSHIRT": {getTShirtItem(), getTShirtItem(), getTShirtItem()},
				},
			},
			want:    map[string]Items{
				"TSHIRT_qweqweqwe": {
					getTShirtItemDiscount(0.95),
					getTShirtItemDiscount(0.95),
					getTShirtItemDiscount(0.95),
				},
			},
			message: "This promo is applicable",
		},
		{
			input: struct {
				PromotionSameItem
				items map[string]Items
			}{
				PromotionSameItem: getPromotionSameItemExample(),
				items:             map[string]Items{
					"TSHIRT": {getTShirtItem(), getTShirtItem(), getTShirtItem(), getTShirtItem()},
				},
			},
			want:    map[string]Items{
				"TSHIRT_qweqweqwe": {
					getTShirtItemDiscount(0.95),
					getTShirtItemDiscount(0.95),
					getTShirtItemDiscount(0.95),
					getTShirtItemDiscount(0.95),
				},
			},
			message: "This promo is applicable",
		},
		{
			input: struct {
				PromotionSameItem
				items map[string]Items
			}{
				PromotionSameItem: getPromotionSameItemExample(),
				items:             map[string]Items{
					"TSHIRT": {getTShirtItem(), getTShirtItem(), getTShirtItem(), getTShirtItem()},
					"VOUCHER": {getVoucherItem(), getVoucherItem()},
				},
			},
			want:    map[string]Items{
				"TSHIRT_qweqweqwe": {
					getTShirtItemDiscount(0.95),
					getTShirtItemDiscount(0.95),
					getTShirtItemDiscount(0.95),
					getTShirtItemDiscount(0.95),
				},
				"VOUCHER": {getVoucherItem(), getVoucherItem()},
			},
			message: "This promo is applicable",
		},
	}

	for i, test := range tests {
		result := test.input.PromotionSameItem.Apply(test.input.items)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("[Test %d] %s", i, test.message)
		}
	}
}