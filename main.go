package main

import (
	"flag"
	"fmt"
	"prueba_cabify/client"
	"prueba_cabify/server"
)

func main() {
	defer logPanicError()
	fmt.Println("Starting...")


	var runMode string
	flag.StringVar(&runMode, "mode", "client", "App mode client|server")
	flag.Parse()

	if runMode == "server" {
		fmt.Println("Stating server")
		server.Run()
	} else {
		fmt.Println("Stating client")
		client.Run()
	}

}

func logPanicError() {
	if r := recover(); r != nil {
		fmt.Println(fmt.Errorf("%+v", r))
	}
}