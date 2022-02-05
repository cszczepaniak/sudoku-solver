package model

type Point struct {
	Row int `json:"row"`
	Col int `json:"col"`
	Box int `json:"-"`
}

func NewPoint(r, c int) Point {
	return Point{
		Row: r,
		Col: c,
		Box: 3*(r/3) + c/3,
	}
}
