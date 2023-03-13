package main

import (
	"github.com/nsf/termbox-go"
	"log"
	"time"
)

func flushRetry() (err error) {
	for i := 1; i <= 2; i++ {
		if err = termbox.Flush(); err == nil {
			return nil
		} else {
			time.Sleep(time.Duration(100_000_000 * i))
		}
	}
	return
}

func setLine(x0, x1, y int, r rune) {
	if x1 == 0 {
		x1, _ = termbox.Size()
	}
	for i := x0; i < x1; i++ {
		termbox.SetChar(i, y, r)
	}
}

// y0  input abcd
// y1  num-kanji
// y2  code
// y3  <hr>
// y4+ buf
var gbi = newBufInfo(4)

func bufTo(r rune) bool {
	if r == gbi.curBuf {
		return false
	}
	termbox.SetChar(0, gbi.curY, ' ')
	gbi.set(r)
	termbox.SetChar(0, gbi.curY, '*')
	return true
}

func initUI() error {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	X, Y := termbox.Size()
	setLine(0, X, 3, '=')
	bufCnt := 0
	for j, r := 4, 'A'; j < Y && r < bufBound; j, r = j+1, r+1 {
		termbox.SetChar(1, j, r)
		bufCnt++
	}
	gbi.initX(bufCnt)
	termbox.SetChar(0, gbi.curY, '*')
	return flushRetry()
}

var codeOn = false

func toggleCode() bool {
	if codeOn {
		codeOn = false
		setLine(0, 0, 2, ' ')
	} else {
		codeOn = true
		pgShowCode()
	}
	return true
}

var gex = newExiter()

func showKj(kj string, x int) int {
	for _, r := range kj {
		termbox.SetChar(x, 1, r)
		x++
	}
	return x + 2
	//return -1
}

var gpg = newPager(eLen)

const num0 rune = '1'

func pgKj() {
	xPos := 0
	idx := num0
	for _, kj := range gpg.getKj() {
		termbox.SetChar(xPos, 1, idx)
		idx++
		xPos++
		termbox.SetChar(xPos, 1, kj)
		xPos += 2
		termbox.SetChar(xPos, 1, ' ')
		xPos++
		termbox.SetChar(xPos, 1, ' ')
		xPos++
	}
	setLine(xPos, 0, 1, ' ')
	if codeOn {
		pgShowCode()
	}
}

func pgShowCode() {
	xPos := 0
	i := 0
	for _, code := range gpg.getCode() {
		termbox.SetChar(xPos, 2, rune(code))
		xPos++
		if i == 3 {
			termbox.SetChar(xPos, 2, ' ')
			xPos++
			i = 0
		} else {
			i++
		}
	}
	setLine(xPos, 0, 2, ' ')
}

var gim = newImData(eLen)

func imInput(r rune) bool {
	if gim.full() {
		return false
	}

	termbox.SetChar(gim.curP, 0, r)
	gim.put(r)
	gpg.put(gim.kj())
	pgKj()
	return true
}

func imBs() bool {
	if gim.empty() {
		return false
	}

	gim.bs()
	termbox.SetChar(gim.curP, 0, ' ')
	gpg.pop()
	if gim.empty() {
		setLine(0, 0, 1, ' ')
		setLine(0, 0, 2, ' ')
	} else {
		pgKj()
	}
	return true
}

func imClear() bool {
	if gim.empty() {
		return false
	}

	for !gim.empty() {
		gim.bs()
	}
	setLine(0, eLen, 0, ' ')
	setLine(0, 0, 1, ' ')
	setLine(0, 0, 2, ' ')
	gpg.clear()
	return true
}

func imSel(idx int) bool {
	r := gpg.get1Kj(idx)
	if r == 0 {
		return false
	}
	termbox.SetChar(gbi.useX(), gbi.curY, r)
	return true
}

func bufClear() bool {
	if gbi.empty() {
		return false
	}
	setLine(bufX0, 0, gbi.curY, ' ')
	gbi.clear()
	return true
}

func pgBack() bool {
	if !gpg.back() {
		return false
	}
	pgKj()
	return true
}

func pgForth() bool {
	if !gpg.forth() {
		return false
	}
	pgKj()
	return true
}

func handleCh(r rune) bool {
	if subjectEx(r) {
		gex.handle(r)
		return false
	}

	if subjectBuf(r) {
		return bufTo(r)
	}

	if subjectIm(r) {
		return imInput(r)
	}

	if subjectSel(r) {
		return imSel(int(r - num0))
	}

	switch r {
	case '?':
		return toggleCode()
	case '\'':
		return imClear()
	case '"':
		return bufClear()
	case ',':
		return pgBack()
	case '.':
		return pgForth()
	}

	return false
}

func evlp() {
	for {
		flush := false
		switch e := termbox.PollEvent(); e.Type {
		case termbox.EventKey:
			if e.Ch != 0 {
				flush = handleCh(e.Ch)
			} else {
				switch e.Key {
				case termbox.KeyBackspace, termbox.KeyBackspace2:
					flush = imBs()
				case termbox.KeySpace:
					flush = imSel(0)
				}
			}
		}
		if gex.bye {
			return
		}
		if flush {
			flushRetry()
		}
	}
}

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	if err := initUI(); err != nil {
		log.Printf("initUI err %v\n", err)
		return
	}
	evlp()
}
