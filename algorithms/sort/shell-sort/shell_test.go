package shell

import "testing"

func Test_shellSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "shell-1",
			args: args{
				a: []int{55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("排序前:%+v", tt.args)
			shellSort(tt.args.a)
			t.Logf("排序后:%+v", tt.args)
		})
	}
}
