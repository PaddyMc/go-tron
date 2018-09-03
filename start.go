/*package main

import (
	"github.com/go-tron/calculations"
	"time"
	"fmt"
)

func main() {
	calculationsClient := calculations.NewClaculationClient("localhost:50051")

	for{
		totalCirculatingSupply := calculationsClient.CalculateCirculatingSupply()
		fmt.Printf("Total Circulating Supply: %v\t\n", totalCirculatingSupply)
		time.Sleep(time.Second * 10)
	}
}*/