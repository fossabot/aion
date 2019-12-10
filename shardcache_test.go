package shardcache

import (
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"testing"
)

func TestSettingAndGetting(t *testing.T) {

	c, err := NewCache(DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 1000000; i++ {
		s := struct {
			x int
		}{x: rand.Intn(1000000)}

		err = c.Set(s)
		if err != nil {
			log.Fatal(err)
		}
		h := hash(s)

		data, _ := c.Get(h)

		assert.Equal(t, s, data)
	}

}
