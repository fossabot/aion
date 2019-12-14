package aion

import (
	"github.com/chronark/charon/logger"
	"sync"
	"time"
)

type shard struct {
	items    map[uint64]item
	lock     sync.RWMutex
	maxSize  uint
	lifeTime uint
}

func newShard(config Config) *shard {
	return &shard{
		items:    make(map[uint64]item, config.MaxShardSize),
		maxSize:  config.MaxShardSize,
		lifeTime: config.Lifetime,
	}

}

func (s *shard) get(hashKey uint64) (interface{}, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	item, found := s.items[hashKey]

	if found {
		if !item.hasExpired() {
			return item.object, true
		} else {
			err := s.delete(hashKey)
			if err != nil {
				logger.Error(err)
			}
		}
	}

	return nil, false
}

func (s *shard) set(hashKey uint64, entry interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.items[hashKey] = item{
		object:    entry,
		endOfLife: uint(time.Now().Unix()) + s.lifeTime,
	}
	return nil

}

func (s *shard) delete(hashKey uint64) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.items, hashKey)
	return nil
}

func (s *shard) len() int {
	return len(s.items)
}
