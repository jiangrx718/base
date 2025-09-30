package demo

import (
	"base/gopkg/gorms"
)

type Dao struct {
	*gorms.BaseDao
}

func NewDao() *Dao {
	return &Dao{
		BaseDao: gorms.NewBaseDao(),
	}
}
