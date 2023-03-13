package main

type imData struct {
	c      []rune
	curP   int // P is x
	length int
	t      *node
	ts     [][]*node
}

func newImData(length int) *imData {
	return &imData{
		c:      make([]rune, length),
		curP:   0,
		length: length,
		t:      newT(),
		ts:     make([][]*node, length),
	}
}

func (id *imData) full() bool {
	return id.curP == id.length
}

func (id *imData) empty() bool {
	return id.curP == 0
}

func (id *imData) put(r rune) {
	if id.full() {
		return
	}
	id.c[id.curP] = r
	if id.empty() {
		id.ts[id.curP] = searchT0(id.t, r)
	} else {
		id.ts[id.curP] = searchT(id.ts[id.curP-1], r)
	}
	id.curP++
}

func (id *imData) kj() []string {
	if id.empty() {
		return nil
	}
	var ret []string
	for _, n := range id.ts[id.curP-1] {
		ret = append(ret, n.kj()...)
	}
	return ret
}

func (id *imData) bs() {
	if id.empty() {
		return
	}
	id.curP--
	id.ts[id.curP] = nil
}

func subjectIm(r rune) bool {
	return ('a' <= r && r <= 'z') || r == wc
}

func subjectSel(r rune) bool {
	return '1' <= r && r <= '9'
}
