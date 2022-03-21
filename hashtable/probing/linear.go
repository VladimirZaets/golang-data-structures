package probing

type Linear struct{}

//linear probing function H(k) + P(K, V) % N
// N - capacity
// The general formula for linear probing is a*i where a is constant.
// a should follow GCD(a, N) = 1
// GCD(3, 9) == 3 - wrong
// GCD(5, 9) == 3 - ok
func (l *Linear) Get(key interface{}, i int) int {
	return i
}
