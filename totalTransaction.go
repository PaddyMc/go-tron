package main

import (
	"flag"
	"fmt"
	"github.com/go-tron/service"
	//"github.com/go-tron/common/hexutil"
	"log"
	"strings"
)

func main() {
	grpcAddress := flag.String("grpcAddress", "",
		"gRPC address: <IP:port> example: -grpcAddress localhost:50051")

	flag.Parse()

	if strings.EqualFold("", *grpcAddress) && len(*grpcAddress) == 0 {
		log.Fatalln("./total-transaction -grpcAddress localhost:50051")
	}

	client := service.NewGrpcClient(*grpcAddress)
	client.Start()
	defer client.Conn.Close()

	//num := client.TotalTransaction()

	//fmt.Printf("total transaction: %v\n", num)

	// Total amount of Tron 												=>		100,000,000,000



	// Block Produce Rewards, getLatestBlock * 32(reward for each block) 	=>	 		59,696,032	+

	// node Rewards getLatestBlock*16										=>			29,848,160	+

	// fee 	(from amount of coin burned per transaction)					=>		1,854,298.69	-



	// Independance day														=>		1,000,000,000	-

	// foundation freeze  (get from genesis block)							=>		33,251,807,425	-

	

	block := client.GetBlockByNum(0)

	// fmt.Printf("block: %v\n", block.GetAmmount())

	for i, v := range  block.GetTransactions() {
		//addr := hexutil.Encode(v.Address)
		value := v
/*		totalProduced := v.TotalProduced
		totalMissed := v.TotalMissed
		latestBlockNum := v.LatestBlockNum
		latestSlotNum := v.LatestSlotNum
		isJobs := v.IsJobs*/
		fmt.Printf("index: %d, rd: %s\n", i, value)
	}
}