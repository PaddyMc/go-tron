package service

import (
	"context"
/*	"crypto/ecdsa"*/
	"github.com/go-tron/api"
/*	"github.com/go-tron/common/base58"
	"github.com/go-tron/common/crypto"*/
	/*"github.com/go-tron/common/hexutil"*/
	"github.com/go-tron/core"
	//"github.com/tronprotocol/go-client-api/util"
	"google.golang.org/grpc"
	"log"
//	"strconv"
//	"time"
)

type GrpcClient struct {
	Address string
	Conn    *grpc.ClientConn
	Client  api.WalletClient
}

func NewGrpcClient(address string) *GrpcClient {
	client := new(GrpcClient)
	client.Address = address
	return client
}

func (g *GrpcClient) Start() {
	var err error
	g.Conn, err = grpc.Dial(g.Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}

	g.Client = api.NewWalletClient(g.Conn)
}

func (g *GrpcClient) TotalTransaction() *api.NumberMessage {

	result, err := g.Client.TotalTransaction(context.Background(),
		new(api.EmptyMessage))

	if err != nil {
		log.Fatalf("total transaction error: %v", err)
	}

	return result
}

func (g *GrpcClient) GetBlockByNum(num int64) *core.Block {
	numMessage := new(api.NumberMessage)
	numMessage.Num = num

	result, err := g.Client.GetBlockByNum2(context.Background(), numMessage)

	if err != nil {
		log.Fatalf("get block by num error: %v", err)
	}

	return result
}

