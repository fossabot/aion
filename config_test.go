package aion

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDefaultConfig(t *testing.T) {
	tests := []struct {
		name string
		want Config
	}{
		{
			name: "Default config",
			want: Config{
				Lifetime:       uint(time.Hour * 24),
				MaxShardSize:   1024,
				NumberOfShards: 16,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DefaultConfig())
		})
	}
}
