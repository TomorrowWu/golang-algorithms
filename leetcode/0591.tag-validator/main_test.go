package main

import "testing"

func TestIsValid(t *testing.T) {
	datas := []struct {
		code     string
		expected bool
	}{
		{"", true},
	}

	for _, d := range datas {
		isV := isValid(d.code)
		t.Logf("expected() expect %t , actual %t", d.expected, isV)
	}
}
