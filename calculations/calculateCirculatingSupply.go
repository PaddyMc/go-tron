package calculations

import (
//	"fmt"
	"github.com/PaddyMc/go-tron/service"
	"log"
	"strings"
)

type CalculationClient struct {
	grpcAddress string
}

func NewClaculationClient(address string) *CalculationClient {
	calc := new(CalculationClient)
	calc.grpcAddress = address
	return calc
}

func (c *CalculationClient) CalculateCirculatingSupply() int64{

	if strings.EqualFold("", c.grpcAddress) && len(c.grpcAddress) == 0 {
		log.Fatalln("Please specify an address e.g 'localhost:50051'")
	}

	client := service.NewGrpcClient(c.grpcAddress)
	client.Start()
	defer client.Conn.Close()

	totalAmountOfTron 	:= 100000000000
	independanceDayBurn := 1000000000

	block := client.GetNowBlock()
	blockHeight := block.GetBlockHeader().RawData.Number
	blockRewards := blockHeight * 32
	nodeRewards := blockHeight * 16

	account := client.GetAccount("TLsV52sRDL79HXGGm9yzwKibb6BeruhUzy")
	startFeeBurnedNum := float64(9223372036854.775808)
	feeBurnedNum := int64(startFeeBurnedNum) + account.Balance/1000000

	foundationFreeze := 33251807425
	
	totalCirculatingSupply := int64(totalAmountOfTron) + int64(blockRewards) + int64(nodeRewards) - int64(independanceDayBurn) - int64(feeBurnedNum) - int64(foundationFreeze)

	return totalCirculatingSupply

	// fmt.Printf("grpcAddress: \t\t%s\n", c.grpcAddress)
	// fmt.Printf("blockHeight: \t\t%v\n", blockHeight)
	// fmt.Printf("+ blockRewards: \t%v\n", blockRewards)
	// fmt.Printf("+ nodeRewards: \t\t%v\n", nodeRewards)
	// fmt.Printf("- feesBurned: \t\t%v\n", feeBurnedNum)
	// fmt.Printf("- independanceDayBurn: \t\t%v\n", independanceDayBurn)
	// fmt.Printf("- foundationFreeze: \t\t%v\n", foundationFreeze)
	// fmt.Printf("\ntotalCirculatingSupply: %v\n\n\n\n", totalCirculatingSupply)

}