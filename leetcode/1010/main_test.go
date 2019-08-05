package main

import (
	"testing"
)

func Test_numPairsDivisibleBy60(t *testing.T) {
	type args struct {
		time []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test-1",
			args: args{
				time: []int{30, 20, 150, 100, 40},
			},
			want: 3,
		},
		{
			name: "Test-1",
			args: args{
				time: []int{60, 60, 60},
			},
			want: 3,
		},
		{
			name: "Test-1",
			args: args{
				time: []int{15, 63, 451, 213, 37, 209, 343, 319},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numPairsDivisibleBy602(tt.args.time); got != tt.want {
				t.Errorf("numPairsDivisibleBy60() = %v, want %v", got, tt.want)
			}
		})
	}
}
