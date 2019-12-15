package aion

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestSettingAndGetting(t *testing.T) {
	for nShards := 1; nShards <= 128; nShards++ {
		config := Config{
			Lifetime:       uint(time.Hour * 24),
			MaxShardSize:   1024,
			NumberOfShards: nShards,
		}

		c, err := NewCache(config)
		if err != nil {
			log.Fatal(err)
		}

		t.Run(fmt.Sprintf("%d shards", nShards), func(t *testing.T) {
			for i := 0; i < rand.Intn(10000)+1; i++ {
				s := struct {
					x int
				}{x: rand.Intn(10000)}

				err = c.Set(string(i), s)
				assert.Nil(t, err)

				data, hit := c.Get(string(i))
				assert.True(t, hit)
				assert.Equal(t, s, data)
			}

		})

	}

}
