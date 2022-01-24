package main

import (
	"fmt"

	"github.com/hasegit/bitex/adapter"
	"github.com/hasegit/bitex/domain"
	"github.com/hasegit/bitex/infrastructure"
)

type JsonRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
	Id      int    `json:"id"`
}

type Params struct {
	Channel string `json:"channel"`
}

func main() {

	client := infrastructure.NewBitFlyerClient()
	defer client.Close()

	repository := adapter.NewBitFlyerRepository(client)
	c := make(chan domain.MarketData)
	go repository.ReadExecutions(c)
	for {
		fmt.Println(<-c)
	}

}
