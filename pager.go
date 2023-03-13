package main

type pager struct {
	content [][]string
	curP    int // content[curP]
	x       int // content[][x]
}

// string->[]rune(str)[-1] for kj ([:-1] for code)
// pre-split the 2 parts saves half retrieval time
// but init splitting takes lot more time...

func onlyKj(s string) rune {
	return []rune(s)[4]
}

const psize int = 8

func newPager(length int) *pager {
	return &pager{
		content: make([][]string, length),
	}
}

func (p *pager) put(s []string) {
	p.content[p.curP] = s
	p.x = 0

	p.curP++
}

func (p *pager) getDataCt() []string {
	if p.curP == 0 {
		return nil
	}
	return p.content[p.curP-1]
}

func (p *pager) getKj() []rune {
	ct := p.getDataCt()
	L := len(ct)
	if L == 0 {
		return nil
	}
	if p.x+psize <= L {
		ret := make([]rune, psize)
		for i, j := p.x, 0; j < psize; i, j = i+1, j+1 {
			ret[j] = onlyKj(ct[i])
		}
		return ret
	} else {
		ret := make([]rune, L-p.x)
		for i, j := L-1, len(ret)-1; j >= 0; i, j = i-1, j-1 {
			ret[j] = onlyKj(ct[i])
		}
		return ret
	}
}

func (p *pager) getCode() []byte {
	ct := p.getDataCt()
	L := len(ct)
	if L == 0 {
		return nil
	}
	if smax := p.x + psize; smax <= L {
		ret := make([]byte, psize*4)
		for i, j := p.x, 0; i < smax; i++ {
			for k := 0; k < 4; k++ {
				ret[j] = ct[i][k]
				j++
			}
		}
		return ret
	} else {
		ret := make([]byte, (L-p.x)*4)
		for i, j := p.x, 0; i < L; i++ {
			for k := 0; k < 4; k++ {
				ret[j] = ct[i][k]
				j++
			}
		}
		return ret
	}
}

func (p *pager) get1Kj(i int) rune {
	ct := p.getDataCt()
	L := len(ct)
	if L == 0 || p.x+i >= L {
		return 0
	}
	return onlyKj(ct[p.x+i])
}

func (p *pager) forth() bool {
	if p.curP == 0 {
		return false
	}
	if p.x+psize < len(p.content[p.curP-1]) {
		p.x += psize
		return true
	} else {
		return false
	}
}

func (p *pager) back() bool {
	if p.curP == 0 {
		return false
	}
	if p.x >= psize {
		p.x -= psize
		return true
	} else {
		return false
	}
}

func (p *pager) pop() {
	if p.curP == 0 {
		return
	}
	p.curP--
	p.content[p.curP] = nil
	p.x = 0
}

func (p *pager) clear() {
	if p.curP == 0 {
		return
	}
	p.curP--
	for p.curP >= 0 {
		p.content[p.curP] = nil
		p.curP--
	}
	p.curP = 0
	p.x = 0
}
