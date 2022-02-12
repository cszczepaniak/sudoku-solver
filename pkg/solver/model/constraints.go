package model

type KillerCage struct {
	Target int
	Cells  map[Point]struct{}
}

func NewKillerCage(target int, points ...Point) KillerCage {
	res := KillerCage{
		Target: target,
		Cells:  make(map[Point]struct{}),
	}
	for _, p := range points {
		res.Cells[p] = struct{}{}
	}
	return res
}
