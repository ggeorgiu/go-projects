package main

import (
	"bytes"
	"math/rand"
)

type Game struct {
	a, b *Board
	h, w int
}

func NewGame(h, w int) *Game {
	a := NewBoard(h, w)
	for i := 0; i < (h*w)/4; i++ {
		a.Set(rand.Intn(h), rand.Intn(w), true)
	}

	return &Game{
		h: h,
		w: w,
		a: a,
		b: NewBoard(h, w),
	}
}

func (g *Game) Step() {
	for i := 1; i < g.h; i++ {
		for j := 1; j < g.w; j++ {
			g.b.Set(i, j, g.a.Next(i, j))
		}
	}

	g.a, g.b = g.b, g.a
}

func (g *Game) String() string {
	var buf bytes.Buffer

	buf.WriteString(g.a.String())

	return buf.String()
}
