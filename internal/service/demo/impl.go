package demo

import (
	"base/internal/dao"
	"base/internal/dao/demo"
)

type Service struct {
	demoDao dao.Demo
}

func NewService() *Service {
	return &Service{
		demoDao: demo.NewDao(),
	}
}
