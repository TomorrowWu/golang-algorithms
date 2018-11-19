package binarysearch

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		a []int
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				a: []int{1, 3, 5, 6, 8},
				v: 8,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.a, tt.args.v); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearchFirstEqual(t *testing.T) {
	type args struct {
		a     []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				a:     []int{1, 2, 2, 2, 3, 4},
				value: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchFirstEqual(tt.args.a, tt.args.value); got != tt.want {
				t.Errorf("binarySearchFirstEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearchLastEqual(t *testing.T) {
	type args struct {
		a     []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				a:     []int{1, 2, 2, 2, 3, 4},
				value: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchLastEqual(tt.args.a, tt.args.value); got != tt.want {
				t.Errorf("binarySearchLastEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearchFirstLargeEqual(t *testing.T) {
	type args struct {
		a     []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				a:     []int{1, 2, 2, 2, 3, 4},
				value: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchFirstLargeEqual(tt.args.a, tt.args.value); got != tt.want {
				t.Errorf("binarySearchFirstLargeEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearchLastSmallEqual(t *testing.T) {
	type args struct {
		a     []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				a:     []int{1, 2, 2, 2, 3, 4},
				value: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchLastSmallEqual(tt.args.a, tt.args.value); got != tt.want {
				t.Errorf("binarySearchLastSmallEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
