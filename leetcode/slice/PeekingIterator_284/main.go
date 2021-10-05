package main

func main() {

}

type Iterator struct {
}

func (i *Iterator) hasNext() bool {
	return true
}

func (i *Iterator) next() int {
	return 1
}

type PeekingIterator struct {
	iter *Iterator
	data *int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter: iter}
}

func (p *PeekingIterator) hasNext() bool {
	return p.data != nil || p.iter.hasNext()
}

func (p *PeekingIterator) next() (ans int) {
	if p.data != nil {
		ans = *p.data
		p.data = nil
		return ans
	}
	return p.iter.next()
}

func (p *PeekingIterator) peek() int {
	if p.data != nil {
		return *p.data
	}
	c := p.iter.next()
	p.data = &c
	return *p.data
}
