package aion

import (
	"errors"
)

type Cache struct {
	shards       []*shard
	close        chan struct{}
	lifetime     int64
	maxShardSize uint
}

func NewCache(config Config) (*Cache, error) {
	cache := &Cache{
		shards:       make([]*shard, config.NumberOfShards),
		lifetime:     config.Lifetime,
		close:        make(chan struct{}),
		maxShardSize: config.MaxShardSize,
	}
	for i := 0; i < config.NumberOfShards; i++ {
		cache.shards[i] = newShard(config)
	}

	return cache, nil
}

func (c *Cache) Close() error {
	close(c.close)
	return nil
}

func (c *Cache) getShard(hashKey uint64) *shard {
	shardID := hashKey & uint64(len(c.shards)-1)
	return c.shards[shardID]
}

func (c *Cache) Get(key string) (data interface{}, hit bool) {
	hashKey := hash(key)
	shard := c.getShard(hashKey)
	return shard.get(hashKey)
}

func (c *Cache) Set(key string, entry interface{}) error {
	if len(key) == 0 {
		return errors.New("key must not be empty")
	}
	hashKey := hash(key)
	shard := c.getShard(hashKey)
	return shard.set(hashKey, entry)
}

func (c *Cache) Len() int {
	var len int
	for _, shard := range c.shards {
		len += shard.len()
	}

	return len
}
