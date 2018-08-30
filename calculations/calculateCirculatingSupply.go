package calculations

import (
	"flag"
	"fmt"
	"github.com/go-tron/service"
	//"github.com/go-tron/common/hexutil"
	"log"
	//"math"
	"strings"
)

type CalculationClient struct {
}

func NewClaculationClient() *CalculationClient {
	client := new(CalculationClient)
	return client
}

func (g *CalculationClient) CalculateCirculatingSupply() {
	grpcAddress := flag.String("grpcAddress", "",
		"gRPC address: <IP:port> example: -grpcAddress localhost:50051")

	flag.Parse()

	if strings.EqualFold("", *grpcAddress) && len(*grpcAddress) == 0 {
		log.Fatalln("./total-transaction -grpcAddress localhost:50051")
	}

	client := service.NewGrpcClient(*grpcAddress)
	client.Start()
	defer client.Conn.Close()


	// // Total amount of Tron 												=>		100,000,000,000
	totalAmountOfTron 	:= 100000000000


	// // Independance day														=>		1,000,000,000	-
	independanceDayBurn := 1000000000


	block := client.GetNowBlock()
	blockHeight := block.GetBlockHeader().RawData.Number
	
	// Block Produce Rewards, getLatestBlock * 32(reward for each block) 	=>	 		59,696,032	+
	blockProduceRewards := blockHeight * 32

	// node Rewards getLatestBlock*16										=>			29,848,160	+
	nodeRewards := blockHeight * 16


	fmt.Printf("blockProduceRewards: \t%v\n", blockProduceRewards)
	fmt.Printf("nodeRewards: \t\t%v\n", nodeRewards)

	//fmt.Printf("TRXSupply: %v\n", int64(totalAmountOfTron) - int64(independanceDayBurn) + blockProduceRewards + nodeRewards)



	// foundation freeze  (get from genesis block)							=>		33,251,807,425	-

	// block := client.GetBlockByNum(0)
	// fmt.Printf("block: %v\n", block.GetTransactions())

	// for i, v := range  block.GetTransactions() {
	// 	//addr := hexutil.Decode(v.Address)
	// 	value := hexutil.Decode(v.Transaction.RawData.Contract[0].Parameter.Value)
	// 	fmt.Printf("index: %d, rd: %s\n", i, value)
	// }


	// Fee Burned 															=>		1,854,298.69	-
	account := client.GetAccount("TLsV52sRDL79HXGGm9yzwKibb6BeruhUzy")
	startFeeBurnedNum := float64(9223372036854.775808)
	//feesBurned := startFeeBurnedNum
	feeBurnedNum := int64(startFeeBurnedNum) + account.Balance/1000000

	fmt.Printf("feesBurned: \t\t%v\n", feeBurnedNum)

	foundationFee := 33251807425
	fmt.Printf("foundationFee: \t\t%v\n", foundationFee)

	totalCirculatingSupply := int64(totalAmountOfTron) + int64(blockProduceRewards) + int64(nodeRewards) - int64(independanceDayBurn) - int64(feeBurnedNum) - int64(foundationFee)

	fmt.Printf("\ntotalCirculatingSupply: %v\n", totalCirculatingSupply)
	// fmt.Printf("startFeeBurnedNum: %v\n", startFeeBurnedNum)
	// fmt.Printf("account.Balance: %v\n", account.Balance)


	//fmt.Printf("account: %v\n", account.Balance)

}