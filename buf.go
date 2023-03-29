package main

// 0=* 1=A 2=space init@3
const bufX0 int = 3

const bufBound rune = 'X'

type bufInfo struct {
	delta  int
	curBuf rune
	curY   int
	x      []int
}

func newBufInfo(startingY int) *bufInfo {
	ret := &bufInfo{
		curBuf: 'A',
		curY:   startingY,
	}
	ret.delta = int(ret.curBuf) - ret.curY
	return ret
}

func (bi *bufInfo) set(r rune) {
	bi.curBuf = r
	bi.curY = int(r) - bi.delta
}

func (bi *bufInfo) initX(t int) {
	bi.x = make([]int, t)
	for i := range bi.x {
		bi.x[i] = bufX0
	}
}

func (bi *bufInfo) empty() bool {
	return bi.x[bi.curY] == bufX0
}

func (bi *bufInfo) clear() {
	bi.x[bi.curY] = bufX0
}

func (bi *bufInfo) useX() int {
	ret := bi.x[bi.curY]
	bi.x[bi.curY] = ret + 2
	return ret
}

func (bi *bufInfo) bsX() (int, bool) {
	p := bi.x[bi.curY]
	if p > 1+bufX0 { // >= bufX0+2 : non-empty
		p -= 2
		bi.x[bi.curY] = p
		return p, true
	} else {
		return 0, false
	}
}

func (bi *bufInfo) y2b(y int) rune {
	return rune(y + bi.delta)
}

func subjectBuf(r rune) bool {
	return 'A' <= r && r < bufBound
}
