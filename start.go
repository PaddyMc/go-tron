package main

import (
	"github.com/go-tron/calculations"
	"time"
)

func main() {

	calculationsClient := calculations.NewClaculationClient("localhost:50051")
	calculationsClient.CalculateCirculatingSupply()

	for{
		calculationsClient.CalculateCirculatingSupply()
		time.Sleep(time.Second * 10)
	}
}