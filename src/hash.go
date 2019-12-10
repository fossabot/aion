package cache

import (
	"hash/fnv"
)

func hash(input []byte) uint64 {
	h := fnv.New64()
	h.Write(input)
	return h.Sum64()
}
