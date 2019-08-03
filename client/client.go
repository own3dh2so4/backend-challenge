package client

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"os"
	"prueba_cabify/proto/basket_service"
	"strings"
)

const (
	address     = "server:9999"
)

func Run() error {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return err
	}

	defer conn.Close()
	c := basket_service.NewBasketServiceClient(conn)
	play(c)
	return nil;

}

func play(c basket_service.BasketServiceClient) {
	command :=""
	reader := bufio.NewReader(os.Stdin)
	var actualBasketID int64
	for command != "exit" {
		fmt.Println("What do you want?")
		text, _ := reader.ReadString('\n')
		command = strings.Replace(text, "\n", "", -1)
		switch command {
		case "createBasket":
			fmt.Println("create basket")
			response, err := c.CreateBasket(context.TODO(), &basket_service.CreateBasketRequest{})
			if err != nil {
				panic(err)
			}
			actualBasketID = response.BasketId;
			fmt.Printf("Basket %d created \n", actualBasketID)
		case "addProduct":
			fmt.Println("What product do you want?")
			text, _ := reader.ReadString('\n')
			request := &basket_service.AddProductRequest{
				BasketId:             actualBasketID,
				ProductId:            strings.Replace(text, "\n", "", -1),
			}
			_, err := c.AddProduct(context.TODO(), request)
			if err != nil {
				fmt.Printf("Error: %s \n", err.Error())
			}
		case "getTotalAmount":
			fmt.Println("get total amount")
			request := &basket_service.GetTotalAmountRequest{
				BasketId: actualBasketID,
			}
			response, err := c.GetTotalAmount(context.TODO(), request)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Total: %.2f \n", response.Price)
		case "removeBasket":
			fmt.Println("remove basket")
			request := &basket_service.RemoveBasketRequest{
				BasketId: actualBasketID,
			}
			_, err := c.RemoveBasket(context.TODO(), request)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Basket %d deleted \n", actualBasketID)
		case "exit":
			fmt.Println("bye")
			break
		default:
			fmt.Printf("invalid command '%s' \n", command)
		}
	}
}