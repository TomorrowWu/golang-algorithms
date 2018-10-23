package src

import "testing"

func Test_countingSort(t *testing.T) {
	type args struct {
		a []int
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "countingSort-1",
			args: args{
				a: []int{5, 4},
				n: 2,
			},
		},
		{
			name: "countingSort-2",
			args: args{
				a: []int{5, 4, 3, 2, 1},
				n: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("排序前:%v", tt.args.a)
			countingSort(tt.args.a, tt.args.n)
			t.Logf("排序后:%v", tt.args.a)
		})
	}
}
