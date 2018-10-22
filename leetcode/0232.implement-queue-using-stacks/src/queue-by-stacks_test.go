package src

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MyQueue
	}{
		{
			name: "queue-constructor-1",
			want: MyQueue{
				Stack1: &stack{
					nums: []int{},
				},
				Stack2: &stack{
					nums: []int{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyQueue_Push(t *testing.T) {
	type fields struct {
		a *stack
		b *stack
	}
	type args struct {
		x int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "queue-push-1",
			fields: fields{
				a: &stack{
					nums: []int{},
				},
				b: &stack{
					nums: []int{},
				},
			},
			args: args{
				x: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := &MyQueue{
				Stack1: tt.fields.a,
				Stack2: tt.fields.b,
			}
			queue.Push(tt.args.x)
		})
	}
}

func TestMyQueue_Pop(t *testing.T) {
	type fields struct {
		a *stack
		b *stack
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "queue-pop-1",
			fields: fields{
				a: &stack{
					nums: []int{3},
				},
				b: &stack{
					nums: []int{2, 1},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := &MyQueue{
				Stack1: tt.fields.a,
				Stack2: tt.fields.b,
			}
			if got := queue.Pop(); got != tt.want {
				t.Errorf("MyQueue.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyQueue_Peek(t *testing.T) {
	type fields struct {
		a *stack
		b *stack
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "test-peek-1",
			fields: fields{
				a: &stack{
					nums: []int{3},
				},
				b: &stack{
					nums: []int{2, 1},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := &MyQueue{
				Stack1: tt.fields.a,
				Stack2: tt.fields.b,
			}
			if got := queue.Peek(); got != tt.want {
				t.Errorf("MyQueue.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyQueue_Empty(t *testing.T) {
	type fields struct {
		a *stack
		b *stack
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "queue-Empty-1",
			fields: fields{
				a: &stack{
					nums: []int{1},
				},
				b: &stack{
					nums: []int{},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := &MyQueue{
				Stack1: tt.fields.a,
				Stack2: tt.fields.b,
			}
			if got := queue.Empty(); got != tt.want {
				t.Errorf("MyQueue.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newStack(t *testing.T) {
	tests := []struct {
		name string
		want *stack
	}{
		{
			name: "stack-1",
			want: &stack{
				nums: []int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_push(t *testing.T) {
	type fields struct {
		nums []int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "stack-push-1",
			fields: fields{
				nums: []int{},
			},
			args: args{
				n: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				nums: tt.fields.nums,
			}
			s.push(tt.args.n)
		})
	}
}

func Test_stack_pop(t *testing.T) {
	type fields struct {
		nums []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "stack-pop-1",
			fields: fields{
				nums: []int{5},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				nums: tt.fields.nums,
			}
			if got := s.pop(); got != tt.want {
				t.Errorf("stack.pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_len(t *testing.T) {
	type fields struct {
		nums []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "stack-len-1",
			fields: fields{
				nums: []int{5},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				nums: tt.fields.nums,
			}
			if got := s.len(); got != tt.want {
				t.Errorf("stack.len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_isEmpty(t *testing.T) {
	type fields struct {
		nums []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "stack-isEmpty-1",
			fields: fields{
				nums: []int{5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				nums: tt.fields.nums,
			}
			if got := s.isEmpty(); got != tt.want {
				t.Errorf("stack.isEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
