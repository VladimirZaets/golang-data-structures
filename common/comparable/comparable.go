package comparable

type Comparable struct {
	index int
	data  interface{}
}

func NewComparable(index int, data interface{}) *Comparable {
	return &Comparable{
		index: index,
		data:  data,
	}
}

func (c *Comparable) Get() interface{} {
	return c.data
}

func (c *Comparable) GetIndex() int {
	return c.index
}
