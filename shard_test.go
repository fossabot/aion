package aion

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func Test_newShard(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
		want *shard
	}{
		{

			name: "default shard",
			args: args{config: DefaultConfig()},
			want: newShard(DefaultConfig()),
		}, {

			name: "custom shard",
			args: args{config: Config{
				Lifetime:       2,
				MaxShardSize:   4,
				NumberOfShards: 1,
			}},
			want: newShard(Config{
				Lifetime:       2,
				MaxShardSize:   4,
				NumberOfShards: 1,
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newShard(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newShard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shard_get(t *testing.T) {
	data := struct {
		x int
	}{
		x: 1,
	}
	type args struct {
		hashKey uint64
	}
	tests := []struct {
		name     string
		s        *shard
		args     args
		wantData interface{}
		wantHit  bool
	}{
		{
			name:     "no hit",
			s:        newShard(DefaultConfig()),
			args:     args{hashKey: hash("empty")},
			wantData: nil,
			wantHit:  false,
		}, {
			name: "hit",
			s: func() *shard {
				shard := newShard(DefaultConfig())
				err := shard.set(hash("hit"), data)
				if err != nil {
					log.Fatal(err)
				}
				return shard
			}(),
			args:     args{hashKey: hash("hit")},
			wantData: data,
			wantHit:  true,
		}, {
			name: "expired hit",
			s: func() *shard {
				shard := newShard(Config{
					Lifetime:       0,
					MaxShardSize:   1024,
					NumberOfShards: 16,
				})
				err := shard.set(hash("expired_hit"), data)
				if err != nil {
					log.Fatal(err)
				}
				return shard
			}(),
			args:     args{hashKey: hash("expired_hit")},
			wantData: nil,
			wantHit:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.get(tt.args.hashKey)
			if !reflect.DeepEqual(got, tt.wantData) {
				t.Errorf("shard.get() gotData = %v, want %v", got, tt.wantData)
			}
			if got1 != tt.wantHit {
				t.Errorf("shard.get() gotHit = %v, want %v", got1, tt.wantHit)
			}
		})
	}
}

func Test_shard_set(t *testing.T) {

	type args struct {
		hashKey uint64
		entry   interface{}
	}
	tests := []struct {
		name    string
		s       *shard
		args    args
		wantErr bool
	}{
		{
			name: "success",
			s:    newShard(DefaultConfig()),
			args: args{
				hashKey: hash("HELLOKEY"),
				entry: struct {
					x int
				}{
					x: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.set(tt.args.hashKey, tt.args.entry); (err != nil) != tt.wantErr {
				t.Errorf("shard.set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_shard_delete(t *testing.T) {
	emptyShard := newShard(DefaultConfig())
	fullShard := newShard(DefaultConfig())
	err := fullShard.set(hash("key"), struct {
		x int
	}{
		x: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	type args struct {
		hashKey uint64
	}
	tests := []struct {
		name       string
		s          *shard
		args       args
		wantLength int
	}{
		{
			name: "can remove",
			s:    fullShard,
			args: args{
				hashKey: hash("key"),
			},
			wantLength: 0,
		},
		{
			name: "empty shard",
			s:    emptyShard,
			args: args{
				hashKey: hash("key"),
			},
			wantLength: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.delete(tt.args.hashKey)
			if tt.s.len() != tt.wantLength {
				t.Errorf("shard.delete() shardLength = %d, wantLength %d", tt.s.len(), tt.wantLength)
			}
		})
	}
}

func Test_shard_len(t *testing.T) {
	shard := newShard(DefaultConfig())
	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("length: %d", i), func(t *testing.T) {
			if got := shard.len(); got != i {
				t.Errorf("Shard.Len() = %v, want %v", got, i)
			}
		})
		err := shard.set(hash(string(i)), struct{ x int }{x: 1})
		if err != nil {
			log.Fatal(err)
		}
	}
}
