package hashtable

import (
	"fmt"

	"github.com/VladimirZaets/godatastructures/doublelinkedlist"
)

type HashTableItem struct {
	key   interface{}
	value interface{}
}

type HashTableSeparateChaining struct {
	capacity      int
	maxLoadFactor float64
	threshold     int
	size          int
	list          []*doublelinkedlist.DoubleLinkedList
}

func NewHashTableSeparateChainint(capacity int) *HashTableSeparateChaining {
	maxLoadFactor := 0.75
	return &HashTableSeparateChaining{
		capacity:      capacity,
		maxLoadFactor: maxLoadFactor,
		threshold:     int(maxLoadFactor * float64(capacity)),
		size:          0,
		list:          make([]*doublelinkedlist.DoubleLinkedList, capacity),
	}
}

//Horner's method to generate a hash O(n)
func (ht *HashTableSeparateChaining) hashKey(k interface{}) int {
	key := fmt.Sprintf("%s", k)
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}
	return h
}

func (ht *HashTableSeparateChaining) getIndex(k interface{}) int {
	return ht.hashKey(k) % ht.capacity
}

func (ht *HashTableSeparateChaining) Set(k interface{}, value interface{}) {
	index := ht.getIndex(k)
	if ht.list[index] == nil {
		ht.list[index] = doublelinkedlist.NewDoubleLinkedList()
	} else if find := ht.findItemInList(ht.list[index], k); find != nil {
		find.value = value
		return
	}
	ht.list[index].AddToTail(&HashTableItem{key: k, value: value})
	ht.size++
	if ht.size > ht.threshold {
		ht.resizeList()
	}
}

func (ht *HashTableSeparateChaining) findItemInList(list *doublelinkedlist.DoubleLinkedList, key interface{}) *HashTableItem {
	node := ht.findNodeInList(list, key)
	if node != nil {
		return node.Get().(*HashTableItem)
	}
	return nil
}

func (ht *HashTableSeparateChaining) findNodeInList(list *doublelinkedlist.DoubleLinkedList, key interface{}) *doublelinkedlist.Node {
	node := list.GetFromHead()
	for node != nil {
		hashTableItem := node.Get().(*HashTableItem)
		if hashTableItem.key == key {
			return node
		}
		node = node.Next()
	}
	return nil
}

func (ht *HashTableSeparateChaining) resizeList() {
	ht.capacity *= 2
	ht.threshold = int(ht.maxLoadFactor * float64(ht.capacity))
	scaledList := make([]*doublelinkedlist.DoubleLinkedList, ht.capacity)
	for i, value := range ht.list {
		if value == nil {
			continue
		}
		node := value.GetFromHead()
		for node != nil {
			hashTableItem := node.Get().(*HashTableItem)
			index := ht.getIndex(hashTableItem.key)
			if scaledList[index] == nil {
				scaledList[index] = doublelinkedlist.NewDoubleLinkedList()
			}
			scaledList[index].AddToTail(hashTableItem)
			node = node.Next()
		}
		value = nil
		ht.list[i] = nil
	}
	ht.list = scaledList
}

func (ht *HashTableSeparateChaining) Remove(k interface{}) error {
	index := ht.getIndex(k)
	ll := ht.list[index]
	node := ht.findNodeInList(ll, k)
	if node == nil {
		return fmt.Errorf("element dont exists in hash table")
	}
	ll.Remove(node)
	ht.size--
	return nil
}

func (ht *HashTableSeparateChaining) Get(k interface{}) interface{} {
	index := ht.getIndex(k)
	hashTableItem := ht.findItemInList(ht.list[index], k)
	if hashTableItem != nil {
		return hashTableItem.value
	}
	return nil
}
