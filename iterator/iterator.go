package iterator

type Aggregate interface {
	NewIterator() Iterator
}

type Iterator interface {
	First()
	Next() interface{}
	IsEnd() bool
}

type NumbersIterator struct {
	numbers *Numbers
	idx     int
}

func (n *NumbersIterator) First() {
	n.idx = n.numbers.start
}

func (n *NumbersIterator) Next() interface{} {
	if !n.IsEnd() {
		tmp := n.idx
		n.idx++
		return tmp
	}
	return nil
}

func (n *NumbersIterator) IsEnd() bool {
	return n.idx > n.numbers.end
}

type Numbers struct {
	start int
	end   int
}

func NewNumbers(s int, e int) *Numbers {
	return &Numbers{
		start: s,
		end:   e,
	}
}

func (n *Numbers) NewIterator() Iterator {
	return &NumbersIterator{
		numbers: n,
		idx:     n.start,
	}
}
