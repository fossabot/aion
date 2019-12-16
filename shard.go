package aion

import (
	"log"
	"sync"
)

type shard struct {
	items    map[uint64]item
	lock     sync.RWMutex
	maxSize  uint
	lifeTime int64
}

func newShard(config Config) *shard {
	return &shard{
		items:    make(map[uint64]item, config.MaxShardSize),
		maxSize:  config.MaxShardSize,
		lifeTime: config.Lifetime,
	}

}

func (s *shard) get(hashKey uint64) (interface{}, bool) {
	//s.lock.RLock()
	//defer s.lock.RUnlock()
	item, found := s.items[hashKey]

	if found {
		if item.hasExpired() {
			s.delete(hashKey)
		} else {
			return item.object, true
		}
	}

	return nil, false
}

func (s *shard) set(hashKey uint64, entry interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.items[hashKey] = newItem(entry, s.lifeTime)
	return nil

}

func (s *shard) delete(hashKey uint64) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.items, hashKey)
}

func (s *shard) len() int {
	return len(s.items)
}
