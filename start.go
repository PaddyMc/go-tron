package main

import (
	"github.com/go-tron/calculations"
)

func main() {
	calculationsClient := calculations.NewClaculationClient()
	calculationsClient.CalculateCirculatingSupply()
}