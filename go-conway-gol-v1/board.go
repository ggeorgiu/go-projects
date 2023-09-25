package main

import "bytes"

type Board struct {
	h int
	w int
	s [][]bool
}

func NewBoard(h, w int) *Board {
	s := make([][]bool, h)
	for i := range s {
		s[i] = make([]bool, w)
	}

	return &Board{
		h: h,
		w: w,
		s: s,
	}
}

func (b *Board) Set(x, y int, f bool) {
	b.s[x][y] = f
}

func (b *Board) IsAlive(x, y int) bool {
	if b.isMargin(x, y) {
		return false
	}

	return b.s[x][y]
}

func (b *Board) Next(x, y int) bool {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if b.isInBounds(x+i, y+j) && b.IsAlive(x+i, y+j) {
				alive++
			}
		}
	}

	return alive == 3 || alive == 2 && b.IsAlive(x, y)
}

func (b *Board) isMargin(i int, j int) bool {
	return i == 0 || j == 0 || i == b.h || j == b.w
}

func (b *Board) isInBounds(i int, j int) bool {
	return i > 0 && i < b.h && j > 0 && j < b.w
}

func (b *Board) getMargin(i, j int) string {
	hl := "\u2550"  // horizontal line
	vl := "\u2551"  // vertical line
	ulc := "\u2554" // upper left corner
	urc := "\u2557" // upper right corner
	llc := "\u255A" // lower left corner
	lrc := "\u255D" // lower right corner

	switch {
	case i == 0 && j == 0:
		return ulc
	case i == 0 && j == b.w:
		return urc
	case i == b.h && j == 0:
		return llc
	case i == b.h && j == b.w:
		return lrc
	case i > 0 && i < b.h && (j == 0 || j == b.w):
		return vl
	case j > 0 && j < b.w && (i == 0 || i == b.h):
		return hl
	}

	return ""
}

func (b *Board) String() string {
	var buf bytes.Buffer
	for i := 0; i <= b.h; i++ {
		for j := 0; j <= b.w; j++ {
			var c string
			switch {
			case b.isMargin(i, j):
				c = b.getMargin(i, j)
			case b.IsAlive(i, j):
				c = "*"
			default:
				c = " "
			}

			buf.WriteString(c)
		}
		buf.WriteByte('\n')
	}

	return buf.String()
}
