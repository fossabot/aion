package aion

import (
	"time"
)

type item struct {
	object    interface{}
	endOfLife uint
}

func newItem(object interface{}, lifeTime uint) item {
	return item{
		object:    object,
		endOfLife: uint(time.Now().Unix()) + lifeTime,
	}
}

func (i item) hasExpired() bool {
	return uint(time.Now().Unix()) > i.endOfLife
}
