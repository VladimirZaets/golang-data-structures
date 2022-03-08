package main

type UnionFind struct {
	ids           []int //Root node id, if ids[i] = i the node is root
	sz            []int
	size          int
	numComponents int
}

func NewUnionFind(size int) *UnionFind {
	sz := make([]int, size)
	ids := make([]int, size)
	for i := 0; i < size-1; i++ {
		ids[i] = i
		sz[i] = i
	}

	return &UnionFind{
		ids:           ids,
		sz:            sz,
		size:          size,
		numComponents: size,
	}
}

func (uf *UnionFind) Find(p int) int {
	root := p
	// Find the root node. If "p" is not root find the root for the node
	for root != uf.ids[root] {
		root = uf.ids[root]
	}

	for p != root {
		next := uf.ids[p]
		uf.ids[p] = root
		p = next
	}

	return root
}

func (uf *UnionFind) Connected(p int, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UnionFind) ComponentSize(p int) int {
	return uf.sz[uf.Find(p)]
}

func (uf *UnionFind) Size() int {
	return uf.size
}

func (uf *UnionFind) Components() int {
	return uf.numComponents
}

func (uf *UnionFind) Unify(p int, q int) {
	root1 := uf.Find(p)
	root2 := uf.Find(q)

	if root1 == root2 {
		return
	}

	if uf.sz[root1] < uf.sz[root2] {
		uf.sz[root2] += uf.sz[root1]
		uf.ids[root1] = root2
	} else {
		uf.sz[root1] += uf.sz[root2]
		uf.ids[root2] = root1
	}

	uf.numComponents--
}
