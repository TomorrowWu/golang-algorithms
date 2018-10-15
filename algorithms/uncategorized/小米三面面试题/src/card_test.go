package src

import "testing"

func TestPositiveSolution(t *testing.T) {

	testCase := []struct {
		n         int
		expectArr []int
	}{
		{
			n:         5,
			expectArr: []int{1, 5, 2, 4, 3},
		},
	}

	for _, c := range testCase {
		result := positiveSolution(c.n)
		t.Logf("expect result:%v, actual result:%v", c.expectArr, result)
	}
}
