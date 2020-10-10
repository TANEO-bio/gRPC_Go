package service

import (
	"context"
	"errors"
	"math/rand"

	pb "../pb"
)

type MyService struct {
}

func (s *MyService) GetRandomNumber(ctx context.Context, message *pb.Request) (*pb.Response, error) {
	if message.Seed >= 0 {
		rand.Seed(message.Seed)
		return &pb.Response{Random: rand.Float32()}, nil
	}
	return nil, errors.New("Error!")
}
