package aion

import (
	"github.com/chronark/charon/logger"
	"hash/fnv"
)

func hash(input string) uint64 {
	h := fnv.New64()
	_, err := h.Write([]byte(input))
	if err != nil {
		logger.Error(err)
	}
	return h.Sum64()
}
