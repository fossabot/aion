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
		endOfLife: uint(time.Now().Add(lifeTime * time.Second)
	}
}

func (i item) hasExpired() bool {
	return uint(time.Now()) > i.endOfLife
}
