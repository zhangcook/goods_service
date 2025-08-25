package grpc

import (
	"flag"
	"github.com/zhangcook/goods_service/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	proto.UnimplementedGoodsServerServer
}

func GoodsService(address string) Server {
	flag.Parse()
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterGoodsServerServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return Server{}
}
