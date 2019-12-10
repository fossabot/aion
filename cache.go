package shardcache

import ()

type Cache struct {
	shards       []shard
	close        chan struct{}
	lifetime     uint
	maxShardSize uint
}

func NewCache(config Config) (Cache, error) {
	cache := Cache{
		shards:       make([]shard, config.NumberOfShards),
		lifetime:     config.Lifetime,
		close:        make(chan struct{}),
		maxShardSize: config.MaxShardSize,
	}
	for i := 0; i < config.NumberOfShards; i++ {
		cache.shards[i] = NewShard(config)
	}

	return cache, nil
}

func (c *Cache) Close() error {
	close(c.close)
	return nil
}

func (c *Cache) getShard(hashKey uint64) *shard {
	shardID := hashKey & uint64(len(c.shards)-1)
	return &c.shards[shardID]
}

func (c *Cache) Get(hashKey uint64) (data interface{}, hit bool) {
	shard := c.getShard(hashKey)
	data, hit = shard.get(hashKey)
	return data, hit
}

func (c *Cache) Set(entry interface{}) error {
	hashKey := hash(entry)
	shard := c.getShard(hashKey)

	err := shard.set(entry)

	return err
}
