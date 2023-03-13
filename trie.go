package main

const eLen int = 4
const eLst int = eLen - 1
const wc rune = '`'

type node struct {
	c []*node
	w []string
}

func newN() *node {
	return &node{c: make([]*node, 26)}
}

func newT() *node {
	ret := newN()
	for _, w := range rawin {
		ret.insert(w, 0)
	}
	return ret
}

func (n *node) insert(w string, l int) {
	if l == eLen {
		n.w = append(n.w, w)
		return
	}
	i := w[l] - 'a'
	child := n.c[i]
	if child == nil {
		if l == eLst {
			child = &node{}
		} else {
			child = newN()
		}
		n.c[i] = child
	}
	child.insert(w, l+1)
}

func (n *node) kj() []string {
	if n == nil {
		return nil
	}
	if len(n.w) != 0 {
		return n.w
	} else {
		var ret []string
		for _, c := range n.c {
			if c != nil {
				ret = append(ret, c.kj()...)
			}
		}
		return ret
	}
}

func searchT0(n *node, r rune) []*node {
	if r == wc {
		return n.c
	}
	return []*node{n.c[r-'a']}
}

func searchT(n []*node, r rune) []*node {
	var ret []*node
	if r == wc {
		for _, e := range n {
			if e != nil {
				ret = append(ret, e.c...)
				// when n is small, faster than
				// `for ec in e.c if ec != nil`
			}
		}
	} else {
		idx := r - 'a'
		for _, e := range n {
			if e != nil && e.c[idx] != nil {
				ret = append(ret, e.c[idx])
			}
		}
	}
	return ret
}
