package aion

import (
	"testing"
	"time"
)

func Test_item_hasExpired(t *testing.T) {
	tests := []struct {
		name string
		i    item
		want bool
	}{
		{
			name: "has expired",
			i: item{
				object: struct {
					x int
				}{
					x: 1,
				},
				endOfLife: uint(time.Now().Unix()) - 2000,
			},
			want: true,
		},
		{
			name: "has not expired",
			i: item{
				object: struct {
					x int
				}{
					x: 1,
				},
				endOfLife: uint(time.Now().Unix()) + 2000,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.hasExpired(); got != tt.want {
				t.Errorf("item.hasExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}
