package aion

import (
	"time"
)

type Config struct {
	Lifetime       uint `json:"lifetime"`
	MaxShardSize   uint `json:"max_shard_size"`
	NumberOfShards int  `json:"number_of_shards"`
}

func DefaultConfig() Config {
	return Config{
		Lifetime:       uint(time.Hour * 24),
		MaxShardSize:   1024,
		NumberOfShards: 16,
	}

}
