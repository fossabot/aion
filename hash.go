package aion

import (
	"hash/fnv"
)

func hash(input string) uint64 {
	h := fnv.New64()
	h.Write([]byte(input))
	return h.Sum64()
}
