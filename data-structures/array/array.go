package array

import (
	"errors"
	"fmt"
)

/**
1) 数组的插入、删除、按照下标随机访问操作；
2）数组中的数据是int类型的；
*/

// Array is a define for Array
type Array struct {
	data   []int
	length uint
}

// NewArray inits an array
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

// Len returns the length of array
func (a *Array) Len() uint {
	return a.length
}

// isIndexOutOfRange determines whether an index overrides an array
func (a *Array) isIndexOutOfRange(index uint) bool {
	return index >= a.length
}

// Find returns the element of index in array
func (a *Array) Find(index uint) (int, error) {
	if a.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return a.data[index], nil
}

// Insert inserts numeric values on index
func (a *Array) Insert(index uint, v int) error {
	if a.Len() == uint(cap(a.data)) {
		return errors.New("full array")
	}
	//直接调用Insert函数,index必须是已有下标
	if index != a.length && a.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}
	for i := a.length; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = v
	a.length++
	return nil
}

// InsertToTail inserts v on tail
func (a *Array) InsertToTail(v int) error {
	return a.Insert(a.Len(), v)
}

// Delete deletes element at index
func (a *Array) Delete(index uint) (int, error) {
	if a.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	v := a.data[index]
	for i := index; i < a.Len()-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
	return v, nil
}

// Print prints the array
func (a *Array) Print() {
	var format string
	for i := uint(0); i < a.Len(); i++ {
		format += fmt.Sprintf("|%+v", a.data[i])
	}
	fmt.Println(format)
}
