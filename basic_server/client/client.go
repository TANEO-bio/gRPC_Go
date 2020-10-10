package main

import (
	"context"
	"fmt"
	"log"

	pb "../pb"
	"google.golang.org/grpc"
)

func main() {
	//sampleなのでwithInsecure
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewMyServiceClient(conn)
	message := &pb.Request{Seed: 2}
	res, err := client.GetRandomNumber(context.TODO(), message)
	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
