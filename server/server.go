package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"prueba_cabify/proto/basket_service"
)

func Run() error {
	list, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption

	s := grpc.NewServer(opts...)

	var basketsrv = BasketServer{}
	basket_service.RegisterBasketServiceServer(s, &basketsrv)

	reflection.Register(s)
	err = s.Serve(list)
	return err
}