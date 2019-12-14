package chronos

import (
	"time"
)

type item struct {
	object    interface{}
	endOfLife uint
}

func (i item) hasExpired() bool {
	return uint(time.Now().Unix()) > i.endOfLife
}
