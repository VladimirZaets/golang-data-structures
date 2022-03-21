package probing

import "hash/fnv"

type DoubleHashing struct{}

//Double hashing probing function H(k) + P(K, V) % N
// N - capacity
// Probing function
// x * h2(key)
// delta = h2(key) mod N
// if delta == 0 need to set delta to 1
func (dh *DoubleHashing) Get(key interface{}, i int) int {
	return i * int(hash(key.(string)))
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
