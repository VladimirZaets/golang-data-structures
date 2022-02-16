package main

import (
	"fmt"
)

type DynamicArray struct {
	arr  []interface{}
	len  int
	last int
}

func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		arr:  make([]interface{}, 5),
		len:  5,
		last: 0,
	}
}

func (darr *DynamicArray) get(index int) (interface{}, error) {
	if darr.last-1 < index {
		return nil, fmt.Errorf("invalid array index %d (out of bounds for %d-element array)", index, darr.last)
	}
	return darr.arr[index], nil
}

func (darr *DynamicArray) append(element interface{}) {
	if darr.last < darr.len {
		darr.arr[darr.last] = element
		darr.last++
	} else {
		newArr := make([]interface{}, darr.len*2)
		copy(newArr, darr.arr)
		darr.len = len(newArr)
		darr.arr = newArr
		newArr = nil
		darr.arr[darr.last] = element
		darr.last++
	}
}

func (darr *DynamicArray) delete(index int) error {
	if darr.last < index {
		return fmt.Errorf("invalid array index %d (out of bounds for %d-element array)", index, len(darr.arr))
	}

	if darr.last-1 <= darr.len/2 {
		darr.len = darr.len / 2
		darr.arr = fillWithoutIndex(make([]interface{}, darr.len), darr, index)
	} else {
		fillWithoutIndex(darr.arr, darr, index)
	}

	return nil
}

func fillWithoutIndex(dst []interface{}, src *DynamicArray, index int) []interface{} {
	src.last = 0
	for i, val := range src.arr {
		if val == nil {
			break
		}
		if i != index {
			dst[src.last] = val
			src.last++
		}
	}
	return dst
}
