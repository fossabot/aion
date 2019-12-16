package aion

import (
	"time"
)

type Config struct {
	Lifetime       int64 `json:"lifetime"`
	MaxShardSize   uint  `json:"max_shard_size"`
	NumberOfShards int   `json:"number_of_shards"`
}

func DefaultConfig() Config {
	return Config{
		Lifetime:       int64(time.Hour * 24),
		MaxShardSize:   1024,
		NumberOfShards: 16,
	}

}
