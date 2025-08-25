package main

import (
	"github.com/zhangcook/goods_service/grpc"
	"github.com/zhangcook/goods_service/proto"
	"log"
)

type Grpc interface {
	RegisterApi() proto.GoodsServerClient
	RegisterRpc() grpc.Server
}

type GrpcClient struct {
	Address string
}

func NewGrpcClient(address string) GrpcClient {

	return GrpcClient{Address: address}
}

func (g *GrpcClient) RegisterApi() proto.GoodsServerClient {
	log.Println("Api启动")
	return grpc.GoodsApi(g.Address)
}

func (g *GrpcClient) RegisterRpc() grpc.Server {
	log.Println("Rpc启动")
	return grpc.GoodsService(g.Address)
}
