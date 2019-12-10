package cache

import ()

type Cache struct {
	shards []shard
	close chan struct{}
	lifetime  uint
	maxShardSize uint
}


func NewCache(config Config) (Cache, error) {
	cache := Cache{
		shards: make([]shard, config.NumberOfShards),
		lifetime: config.lifetime,
		close: make(chan struct{})
	}
	return cache, nil
}

func (c Cache) Close() error{
	close(c.close)
	return nil
}