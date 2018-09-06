package main

import "testing"

func Test_maxPoints(t *testing.T) {
	t.Logf("最多的点数:%d\n", maxPoints2([]Point{
		{
			X: 1,
			Y: 1,
		},
		{
			X: 2,
			Y: 2,
		},
		{
			X: 3,
			Y: 3,
		},
	})) //3

	t.Logf("最多的点数:%d\n", maxPoints2([]Point{
		{
			X: 1,
			Y: 1,
		},
		{
			X: 3,
			Y: 2,
		},
		{
			X: 5,
			Y: 3,
		},
		{
			X: 4,
			Y: 1,
		},
		{
			X: 2,
			Y: 3,
		},
		{
			X: 1,
			Y: 4,
		},
	})) //4

	t.Logf("最多的点数:%d\n", maxPoints2([]Point{
		{
			X: 0,
			Y: 0,
		},
		{
			X: 0,
			Y: 0,
		},
	})) //2

	t.Logf("最多的点数:%d\n", maxPoints2([]Point{
		{
			X: 0,
			Y: -12,
		},
		{
			X: 5,
			Y: 2,
		},
		{
			X: 2,
			Y: 5,
		},
		{
			X: 0,
			Y: -5,
		},
		{
			X: 1,
			Y: 5,
		},
		{
			X: 2,
			Y: -2,
		},
		{
			X: 5,
			Y: -4,
		},
		{
			X: 3,
			Y: 4,
		},
		{
			X: -2,
			Y: 4,
		},
		{
			X: -1,
			Y: 4,
		},
		{
			X: 0,
			Y: -5,
		},
		{
			X: 0,
			Y: -8,
		},
		{
			X: -2,
			Y: -1,
		},
		{
			X: 0,
			Y: -11,
		},
		{
			X: 0,
			Y: -9,
		},
	})) //6
}

func Test_mapKey(t *testing.T) {
	t.Logf("两key相等=%t", Point{
		X: 0,
		Y: 0,
	} == Point{
		X: 0,
		Y: 0,
	})
}
