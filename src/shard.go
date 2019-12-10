package cache

import (
	"sync"
	"time"
)

type Shard struct {
	items    map[uint64]item
	lock     *sync.RWMutex
	size     uint
	lifeTime int64
}

func NewShard(config Config) (Shard, error) {
	shard := Shard{
		items: make(map[uint64]item, config.ShardSize),
		size:  config.ShardSize,
	}

	return shard, nil
}

func (s Shard) get(key uint64) (interface{}, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	item, found := s.items[key]
	if found {
		return item.object, true
	}
	return nil, false
}

func (s Shard) set(key uint64, entry interface{}) error {
	s.lock.RLock()
	defer s.lock.RUnlock()

	s.items[key] = item{
		object:    entry,
		endOfLife: time.Now().Unix() + s.lifeTime,
	}

	return nil

}
