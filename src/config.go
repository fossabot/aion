package cache

import (
	"time"
)

type Config struct {
	Lifetime  uint
	MaxShardSize uint
	NumberOfShards uint

}

func DefaultConfig() Config {
	return Config{
		Lifetime:  uint(time.Hour * 24),
		MaxShardSize: 1024,
	NumberOfShards: 4,
	}

}
