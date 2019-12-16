package aion

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"testing"
)

func TestNewCache(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name    string
		args    args
		want    *Cache
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCache(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCache_Close(t *testing.T) {

	cache, err := NewCache(DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}

	if err := cache.Close(); err != nil {
		t.Errorf("Cache.Close() error = %v, wantErr %v", err, nil)
	}
}

func TestCache_getShard(t *testing.T) {
	type args struct {
		hashKey uint64
	}
	tests := []struct {
		name string
		c    *Cache
		args args
		want *shard
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.getShard(tt.args.hashKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cache.getShard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCache_Get(t *testing.T) {

	key := "RANDOMKEY"
	data := struct {
		x int
	}{x: rand.Intn(10000)}

	type args struct {
		key string
	}
	tests := []struct {
		name     string
		c        *Cache
		args     args
		wantData interface{}
		wantHit  bool
	}{
		{
			name: "miss",
			c: func() *Cache {
				cache, err := NewCache(DefaultConfig())
				if err != nil {
					log.Fatal(err)
				}
				return cache
			}(),
			args:     args{key: key},
			wantData: nil,
			wantHit:  false,
		},
		{
			name: "hit",
			c: func() *Cache {
				cache, err := NewCache(DefaultConfig())
				if err != nil {
					log.Fatal(err)
				}
				err = cache.Set(key, data)
				if err != nil {
					log.Fatal(err)
				}
				return cache
			}(),
			args:     args{key: key},
			wantData: data,
			wantHit:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, gotHit := tt.c.Get(tt.args.key)
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("Cache.Get() gotData = %v, want %v", gotData, tt.wantData)
			}
			if gotHit != tt.wantHit {
				t.Errorf("Cache.Get() gotHit = %v, want %v", gotHit, tt.wantHit)
			}
		})
	}
}

func TestCache_Set(t *testing.T) {
	cache, err := NewCache(DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}

	type args struct {
		key   string
		entry interface{}
	}
	tests := []struct {
		name    string
		c       *Cache
		args    args
		wantErr bool
	}{
		{
			name: "empty key",
			c:    cache,
			args: args{
				key: "",
				entry: struct {
					x int
				}{
					x: rand.Intn(10000),
				},
			},
			wantErr: true,
		},
		{
			name: "success",
			c:    cache,
			args: args{
				key: "HELLOKEY",
				entry: struct {
					x int
				}{
					x: rand.Intn(10000),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Set(tt.args.key, tt.args.entry); (err != nil) != tt.wantErr {
				t.Errorf("Cache.Set() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func TestCache_Len(t *testing.T) {
	cache, err := NewCache(DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("length: %d", i), func(t *testing.T) {
			if got := cache.Len(); got != i {
				t.Errorf("Cache.Len() = %v, want %v", got, i)
			}
		})
		err := cache.Set(string(i), struct{ x int }{x: 1})
		if err != nil {
			log.Fatal(err)
		}
	}
}
