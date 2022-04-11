package fenwicktree

type FenwickTree struct {
	list []int
}

func NewFenwickTree(list []int) *FenwickTree {
	return &FenwickTree{
		list: buildFenwickTree(list),
	}
}

func (ft *FenwickTree) prefixSum(i int) int {
	sum := 0
	for i != 0 {
		sum += ft.list[i]
		i -= leastSignificantBit(i)
	}
	return sum
}

func (ft *FenwickTree) Sum(from int, to int) int {
	return ft.prefixSum(to) - ft.prefixSum(from-1)
}

func (ft *FenwickTree) Add(index int, num int) {
	for index < len(ft.list) {
		ft.list[index] += num
		index += leastSignificantBit(index)
	}
}

func (ft *FenwickTree) Set(index int, val int) {
	v := ft.Sum(index, index)
	ft.Add(index, val-v)
}

func buildFenwickTree(list []int) []int {
	rl := list
	for i := 1; i < len(rl); i++ {
		j := i + leastSignificantBit(i)
		if j < len(rl) {
			rl[j] += rl[i]
		}
	}
	return rl
}

func leastSignificantBit(i int) int {
	return i & -i
}
