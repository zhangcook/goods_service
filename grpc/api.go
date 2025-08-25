package grpc

import (
	"github.com/zhangcook/goods_service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var ApiService struct {
	GoodsClient proto.GoodsServerClient
}

func GoodsApi(address string) proto.GoodsServerClient {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ApiService.GoodsClient = proto.NewGoodsServerClient(conn)
	return ApiService.GoodsClient
}
