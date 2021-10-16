package solver

type point struct {
	row int
	col int
	box int
}

func newPoint(r, c int) point {
	return point{
		row: r,
		col: c,
		box: 3*(r/3) + c/3,
	}
}

type puzzleCache struct {
	rows  []pointCache
	cols  []pointCache
	boxes []pointCache
}

func newPuzzleCache() *puzzleCache {
	pc := &puzzleCache{
		rows:  make([]pointCache, Dimension),
		cols:  make([]pointCache, Dimension),
		boxes: make([]pointCache, Dimension),
	}
	for i := 0; i < Dimension; i++ {
		pc.rows[i] = make(pointCache)
		pc.cols[i] = make(pointCache)
		pc.boxes[i] = make(pointCache)
	}
	return pc
}

func (pc *puzzleCache) add(r, c, n int) {
	if n == 0 {
		return
	}
	pt := newPoint(r, c)
	pc.rows[pt.row].add(pt, n)
	pc.cols[pt.col].add(pt, n)
	pc.boxes[pt.box].add(pt, n)
}

func (pc *puzzleCache) remove(r, c, n int) {
	pt := newPoint(r, c)
	pc.rows[pt.row].remove(pt, n)
	pc.cols[pt.col].remove(pt, n)
	pc.boxes[pt.box].remove(pt, n)
}

func (pc *puzzleCache) isValidEntry(r, c, n int) bool {
	pt := newPoint(r, c)
	return pc.rows[pt.row].isValidEntry(n) &&
		pc.cols[pt.col].isValidEntry(n) &&
		pc.boxes[pt.box].isValidEntry(n)
}

func (pc *puzzleCache) validateDuplicates() []*InvalidSquareError {
	var errSet map[point]*InvalidSquareError
	for i := 0; i < Dimension; i++ {
		for _, err := range pc.rows[i].getInvalidEntries() {
			if errSet == nil {
				errSet = make(map[point]*InvalidSquareError)
			}
			errSet[newPoint(err.Row, err.Col)] = err
		}
		for _, err := range pc.cols[i].getInvalidEntries() {
			if errSet == nil {
				errSet = make(map[point]*InvalidSquareError)
			}
			errSet[newPoint(err.Row, err.Col)] = err
		}
		for _, err := range pc.boxes[i].getInvalidEntries() {
			if errSet == nil {
				errSet = make(map[point]*InvalidSquareError)
			}
			errSet[newPoint(err.Row, err.Col)] = err
		}
	}
	errs := make([]*InvalidSquareError, 0, len(errSet))
	for _, err := range errSet {
		errs = append(errs, err)
	}
	return errs
}

type pointCache map[int]map[point]struct{}

func (pc pointCache) add(pt point, n int) {
	ptSet, ok := pc[n]
	if !ok {
		ptSet = make(map[point]struct{})
		pc[n] = ptSet
	}
	ptSet[pt] = struct{}{}
}

func (pc pointCache) remove(pt point, n int) {
	delete(pc[n], pt)
}

func (pc pointCache) isValidEntry(n int) bool {
	return len(pc[n]) == 0
}

func (pc pointCache) getInvalidEntries() []*InvalidSquareError {
	var errs []*InvalidSquareError
	for _, pts := range pc {
		if len(pts) <= 1 {
			// valid
			continue
		}
		for pt := range pts {
			errs = append(errs, newInvalidSquareError(pt.row, pt.col, duplicateNumber))
		}
	}
	return errs
}
