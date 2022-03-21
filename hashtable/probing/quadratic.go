package probing

type Quadratic struct{}

//Quadratic probing function H(k) + P(K, V) % N
// N - capacity
// Probing function
// 1) P(x) = x*x  - keep the table size a prime number > 3 and also maxLoadFactor <= 1/2
// 2) P(x) = (x^2 + x)/2  - keep the table size a power of two
// 3) P(x) = (-1^x)*x^2
func (q *Quadratic) Get(key interface{}, i int) int {
	return (i*i + i) / 2
}
