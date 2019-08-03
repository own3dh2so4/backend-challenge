package server

import (
	"context"
	"fmt"
	"prueba_cabify/proto/basket_service"
	"prueba_cabify/repository"
	"prueba_cabify/service"
)

type BasketServer struct {

}

func (BasketServer) CreateBasket(context.Context, *basket_service.CreateBasketRequest) (*basket_service.CreateBasketResponse, error) {
	fmt.Println("Create basket")
	basket := service.GetBasketPromotion()
	err := repository.GetBasketRepository().AddBasket(basket)
	if err != nil {
		return nil, err
	}
	return &basket_service.CreateBasketResponse{
		BasketId:             int64(basket.GetID()),
	}, nil
}

func (BasketServer) AddProduct(ctx context.Context, request *basket_service.AddProductRequest) (*basket_service.AddProductResponse, error) {
	fmt.Println("Add product")
	basket, err := repository.GetBasketRepository().GetBasket(int(request.BasketId))
	if err != nil {
		return nil, err
	}
	item, err := repository.GetItemsRepository().GetItem(request.ProductId)
	if err != nil {
		return nil, err
	}
	basket.AddItem(item)
	return &basket_service.AddProductResponse{}, nil
}

func (BasketServer) GetTotalAmount(ctx context.Context, request *basket_service.GetTotalAmountRequest) (*basket_service.GetTotalAmountResponse, error) {
	fmt.Println("Get total")
	basket, err := repository.GetBasketRepository().GetBasket(int(request.BasketId))
	if err != nil {
		return nil, err
	}
	return &basket_service.GetTotalAmountResponse{
		Price: float32(basket.GetPrice()),
	}, nil
}

func (BasketServer) RemoveBasket(ctx context.Context, request *basket_service.RemoveBasketRequest) (*basket_service.RemoveBasketResponse, error) {
	fmt.Println("Remove basket")
	repository.GetBasketRepository().DeleteBasket(int(request.BasketId))
	return &basket_service.RemoveBasketResponse{}, nil
}