package cache

import (
	"sync"
	"time"
)

type shard struct {
	items    map[uint64]item
	lock     *sync.RWMutex
	maxSize  uint
	lifeTime uint
}

func NewShard(config Config) (shard, error) {
	shard := shard{
		items:    make(map[uint64]item, config.MaxShardSize),
		maxSize:  config.MaxShardSize,
		lifeTime: config.Lifetime,
	}

	return shard, nil
}

func (s shard) get(key uint64) (interface{}, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	item, found := s.items[key]
	if found {
		return item.object, true
	}
	return nil, false
}

func (s shard) set(key uint64, entry interface{}) error {
	s.lock.RLock()
	defer s.lock.RUnlock()

	s.items[key] = item{
		object:    entry,
		endOfLife: uint(time.Now().Unix()) + s.lifeTime,
	}

	return nil

}
