package src

import (
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "quick-sort-1",
			args: args{
				a: []int{3, 2, 1, 5, 6, 4},
			},
		},
		{
			name: "quick-sort-2",
			args: args{
				a: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("排序前:%+v", tt.args.a)
			quickSort(tt.args.a)
			t.Logf("排序后:%+v", tt.args.a)
		})
	}
}

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "findKthLargest-1",
			args: args{
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			want: 5,
		},
		{
			name: "findKthLargest-2",
			args: args{
				nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
				k:    4,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
