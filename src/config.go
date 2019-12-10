package cache

import (
	"time"
)

type Config struct {
	Lifetime  uint
	ShardSize uint
}

func DefaultConfig() Config {
	return Config{
		Lifetime:  uint(time.Hour * 24),
		ShardSize: 1024,
	}

}
