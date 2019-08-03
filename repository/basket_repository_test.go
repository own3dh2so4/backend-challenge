package repository

import (
	"prueba_cabify/model"
	"reflect"
	"testing"
)

func TestAddBasket(t *testing.T) {
	basketRepo := GetBasketRepository()
	basket := model.GetNewBasket()
	err := basketRepo.AddBasket(basket)
	if err != nil {
		t.Errorf("basket '%d' should be added", basket.GetID())
	}
	basketResult, err := basketRepo.GetBasket(basket.GetID())
	if err != nil {
		t.Errorf("basket '%d' must exists", basket.GetID())
	} else if !reflect.DeepEqual(basket, basketResult) {
		t.Errorf("basket wanted %+v, basket received %+v", basket, basketResult)
	}
}

func TestAddBasketMultipleTimes(t *testing.T) {
	basketRepo := GetBasketRepository()
	basket := model.GetNewBasket()
	err := basketRepo.AddBasket(basket)
	if err != nil {
		t.Errorf("basket '%d' should be added", basket.GetID())
	}
	err = basketRepo.AddBasket(basket)
	if err == nil {
		t.Errorf("basket '%d' can not be added again", basket.GetID())
	}
}

func TestGetBasketRepositoryDoesNotExist(t *testing.T) {
	basketRepo := GetBasketRepository()
	basket := model.GetNewBasket()
	basketResult, err := basketRepo.GetBasket(basket.GetID())
	if err == nil {
		t.Errorf("basket '%d' must not exists", basket.GetID())
	}
	if basketResult != nil {
		t.Error("basket result must be nil")
	}
}