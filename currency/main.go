package main

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"working/currency/protos"
	"working/currency/server"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	cs := server.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, cs)
	fmt.Println("Hello, World!", gs, log)
}
