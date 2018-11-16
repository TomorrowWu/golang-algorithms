package src

import (
	"reflect"
	"testing"
)

//func TestConstructor(t *testing.T) {
//	type args struct {
//		nums []int
//	}
//	tests := []struct {
//		name string
//		args args
//		want Solution
//	}{
//		{
//			name: "1",
//			args: args{
//				nums: []int{1, 2, 3, 3, 3},
//			},
//			want: Solution{
//				m: map[int][]int{
//					1: []int{0},
//					2: []int{1},
//					3: []int{2, 3, 4},
//				},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := Constructor(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Constructor() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestSolution_Pick(t *testing.T) {
//	type fields struct {
//		m map[int][]int
//	}
//	type args struct {
//		target int
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   int
//	}{
//		{
//			name: "1",
//			fields: fields{
//				m: map[int][]int{
//					1: []int{0},
//					2: []int{1},
//					3: []int{2, 3, 4},
//				},
//			},
//			args: args{
//				target: 1,
//			},
//			want: 0,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Solution{
//				m: tt.fields.m,
//			}
//			if got := s.Pick(tt.args.target); got != tt.want {
//				t.Errorf("Solution.Pick() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestConstructor(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want Solution
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 3, 3, 3},
			},
			want: Solution{
				nums: []int{1, 2, 3, 3, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution_Pick(t *testing.T) {
	type fields struct {
		nums []int
	}
	type args struct {
		target int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "1",
			fields: fields{
				nums: []int{1, 2, 3, 3, 3},
			},
			args: args{
				target: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solution{
				nums: tt.fields.nums,
			}
			if got := s.Pick(tt.args.target); got != tt.want {
				t.Errorf("Solution.Pick() = %v, want %v", got, tt.want)
			}
		})
	}
}
