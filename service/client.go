package service

import (
	"context"
	"github.com/PaddyMc/go-tron/api"
	"github.com/PaddyMc/go-tron/common/base58"
	"github.com/PaddyMc/go-tron/core"
	"google.golang.org/grpc"
	"log"
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

func (g *GrpcClient) GetBlockByNum(num int64) (*api.BlockExtention, error) {
	numMessage := new(api.NumberMessage)
	numMessage.Num = num

	result, err := g.Client.GetBlockByNum2(context.Background(), numMessage)

	return result, err
}

func (g *GrpcClient) GetAccount(address string) (*core.Account, error){
	account := new(core.Account)

	account.Address = base58.DecodeCheck(address)

	result, err := g.Client.GetAccount(context.Background(), account)

	return result, err
}

func (g *GrpcClient) GetNowBlock() (*api.BlockExtention, error){
	result, err := g.Client.GetNowBlock2(context.Background(), new(api.EmptyMessage))

	return result, err
}
