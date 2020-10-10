package main

import (
	"log"
	"net"

	pb "../pb"
	"../service"

	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	Service := &service.MyService{}
	// 実行したい実処理をseverに登録する
	pb.RegisterMyServiceServer(server, Service)
	server.Serve(listenPort)
}
