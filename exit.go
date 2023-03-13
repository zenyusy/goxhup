package main

import "time"

type exiter struct {
	t     time.Time
	bye   bool
	thres time.Duration
}

func newExiter() *exiter {
	return &exiter{t: time.Now(), thres: time.Duration(400_000_000)}
}

func (e *exiter) q1() {
	if !e.bye {
		e.t = time.Now()
	}
}

func (e *exiter) q2() {
	if !e.bye {
		d := time.Now().Sub(e.t)
		if d > 0 && d < e.thres {
			e.bye = true
		}
	}
}

func (e *exiter) handle(r rune) {
	switch r {
	case '[':
		e.q1()
	case ']':
		e.q2()
	}
}

func subjectEx(r rune) bool {
	return r == '[' || r == ']'
}
