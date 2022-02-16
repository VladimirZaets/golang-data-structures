package main

import (
	"math/rand"
	"testing"
)

func TestAppendItemToArray(t *testing.T) {
	dynamicArray := NewDynamicArray()
	for i := 0; i < 10; i++ {
		dynamicArray.append(rand.Intn(100))
	}

	if len(dynamicArray.arr) != 10 {
		t.Errorf("Length of array is %d, the length should be 10", len(dynamicArray.arr))
	}
}

func TestGetFromArray(t *testing.T) {
	dynamicArray := NewDynamicArray()
	for i := 0; i < 2; i++ {
		dynamicArray.append(i)
	}

	val, err := dynamicArray.get(1)

	if err != nil {
		t.Error(err)
		return
	}

	if val != 1 {
		t.Errorf("Length of array is %d, the length should be 1", val)
	}
}

func TestDeleteFromArray(t *testing.T) {
	dynamicArray := NewDynamicArray()
	for i := 0; i < 6; i++ {
		dynamicArray.append(rand.Intn(100))
	}

	for i := 0; i < 6; i++ {
		dynamicArray.delete(i)
	}

	dynamicArray.delete(3)
	if dynamicArray.len != 5 && dynamicArray.last != 1 {
		t.Errorf("Length of array is %d, the length should be 5, capacity is %d, the capacity should be 5", dynamicArray.last, dynamicArray.len)
	}
}

func BenchmarkAppendToArray(b *testing.B) {
	dynamicArray := NewDynamicArray()
	for i := 0; i < 1000000; i++ {
		dynamicArray.append(i)
	}
}

func BenchmarkDeleteFromArray(b *testing.B) {
	dynamicArray := NewDynamicArray()
	iteration := 10000
	for i := 0; i < iteration; i++ {
		dynamicArray.append(i)
	}
	for i := 0; i < iteration; i++ {
		dynamicArray.delete(i)
	}
}
