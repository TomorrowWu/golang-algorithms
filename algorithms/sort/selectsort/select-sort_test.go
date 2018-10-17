package selectsort

import "testing"

func TestSelectSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
	}{
		{
			name: "select-sort-1",
			arr:  []int{10, 34, 19, 100, 80},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("排序前:%+v", tt.arr)
			SelectSort(tt.arr)
			t.Logf("排序后:%+v", tt.arr)
		})
	}
}
