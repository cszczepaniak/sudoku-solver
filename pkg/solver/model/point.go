package model

type Point struct {
	row int
	col int
	box int
}

func NewPoint(r, c int) Point {
	return Point{
		row: r,
		col: c,
		box: 3*(r/3) + c/3,
	}
}

func (p Point) Row() int {
	return p.row
}

func (p Point) Col() int {
	return p.col
}

func (p Point) Box() int {
	return p.box
}

func (p Point) LinearIndex() int {
	return 9*p.Row() + p.Col()
}
