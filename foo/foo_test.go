package foo

import (
	"math"
	"testing"
)


func Test_add(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "0+0",
			a:    0,
			b:    0,
			want: 0,
		},
		{
			name: "10+13",
			a:    10,
			b:    13,
			want: 23,
		},
		{
			name: "MaxInt32",
			a:    math.MaxInt32,
			b:    math.MaxInt32,
			want: 4294967294,
		},
		{
			name: "MaxInt32",
			a:    math.MaxInt64,
			b:    math.MaxInt64,
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.a, tt.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
