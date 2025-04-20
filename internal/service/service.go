package service

import test "OrderProject/pkg/api/test/api"

type Service struct {
	test.OrderServiceServer
}

func New() *Service {
	return &Service{}
}
