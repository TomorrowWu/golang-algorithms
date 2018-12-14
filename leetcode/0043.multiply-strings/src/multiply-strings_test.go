package src

import "testing"

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				num1: "2",
				num2: "3",
			},
			want: "6",
		},
		{
			name: "test2",
			args: args{
				num1: "123",
				num2: "456",
			},
			want: "56088",
		},
		{
			name: "test3",
			args: args{
				num1: "999",
				num2: "999",
			},
			want: "998001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
