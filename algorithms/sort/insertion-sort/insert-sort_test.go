package insertionsort

import "testing"

func TestInsertSort(t *testing.T) {
	tests := []struct {
		name  string
		array []int
	}{
		{
			name:  "sort1",
			array: []int{23, 0, 12, 56, 34},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("排序前:%+v", tt.array)
			InsertSort(tt.array)
			t.Logf("排序后:%+v", tt.array)
		})
	}
}
