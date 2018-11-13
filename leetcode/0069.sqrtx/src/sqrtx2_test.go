package src

import "testing"

func Test_mySqrt3(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "one",
			args: args{
				x: 1,
			},
			want: 1,
		},
		{
			name: "two",
			args: args{
				x: 2,
			},
			want: 1.414213,
		},
		{
			name: "three",
			args: args{
				x: 10,
			},
			want: 3.162277,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt3(tt.args.x); got != tt.want {
				t.Errorf("mySqrt3() = %v, want %v", got, tt.want)
			}
		})
	}
}
