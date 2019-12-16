package aion

import (
	"time"
)

type item struct {
	object    interface{}
	endOfLife int64
}

func newItem(object interface{}, timeToLive int64) item {
	return item{
		object:    object,
		endOfLife: time.Now().Unix() + timeToLive,
	}
}

func (i item) hasExpired() bool {
	return time.Now().Unix() > i.endOfLife
}
