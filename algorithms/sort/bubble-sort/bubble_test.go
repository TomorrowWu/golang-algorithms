package bubble

import "testing"

func Test_bubbleSort(t *testing.T) {
	tests := []struct {
		name string
		a    []int
	}{
		{
			name: "bubble-sort-1",
			a:    []int{3, 2, 4, 5, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("排序前:%+v", tt.a)
			bubbleSort(tt.a)
			t.Logf("排序后:%+v", tt.a)
		})
	}
}
