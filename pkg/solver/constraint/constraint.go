package constraint

type Constraint interface {
	Evaluate(n int) bool
}

type sudokuConstraint struct {
	row map[int]struct{}
	col map[int]struct{}
	box map[int]struct{}
}

func (c *sudokuConstraint) Evaluate(n int) bool {
	for _, set := range []map[int]struct{}{c.row, c.col, c.box} {
		if _, ok := set[n]; ok {
			return false
		}
	}
	return true
}
