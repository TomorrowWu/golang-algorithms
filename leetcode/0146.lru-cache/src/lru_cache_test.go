package src

import "testing"

func TestLRUCache_Put(t *testing.T) {
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "LRUCache-Put-1",
			args: args{
				key:   0,
				value: 0,
			},
		},
		{
			name: "LRUCache-Put-2",
			args: args{
				key:   1,
				value: 1,
			},
		},
		{
			name: "LRUCache-Put-3",
			args: args{
				key:   2,
				value: 2,
			},
		},
		{
			name: "LRUCache-Put-4",
			args: args{
				key:   3,
				value: 3,
			},
		},
		{
			name: "LRUCache-Put-5",
			args: args{
				key:   3,
				value: 33,
			},
		},
	}

	l := Constructor(1)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Put(tt.args.key, tt.args.value)
		})
	}

	t.Logf("cache Get(%d),got %d", 3, l.Get(3))
}
