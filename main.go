package main

import (
	"github.com/chronark/cache/src/cache"
"log"
)


func main(){
	cache, err := cache.NewCache(cache.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}

}
