package chronos

import (
	"fmt"
	"github.com/chronark/charon/logger"
	"hash/fnv"
)

func PublicHash(input interface{}) uint64 {
	return hash(input)
}

func hash(input interface{}) uint64 {
	h := fnv.New64()
	_, err := h.Write([]byte(fmt.Sprintf("%v", input)))
	if err != nil {
		logger.Error(err)
	}
	return h.Sum64()
}
