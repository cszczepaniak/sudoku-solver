package solver

type Point struct {
	Row int
	Col int
	Box int
}

func newPoint(r, c int) Point {
	return Point{
		Row: r,
		Col: c,
		Box: 3*(r/3) + c/3,
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
	pc.rows[pt.Row].add(pt, n)
	pc.cols[pt.Col].add(pt, n)
	pc.boxes[pt.Box].add(pt, n)
}

func (pc *puzzleCache) remove(r, c, n int) {
	pt := newPoint(r, c)
	pc.rows[pt.Row].remove(pt, n)
	pc.cols[pt.Col].remove(pt, n)
	pc.boxes[pt.Box].remove(pt, n)
}

func (pc *puzzleCache) isValidEntry(r, c, n int) bool {
	pt := newPoint(r, c)
	return pc.rows[pt.Row].isValidEntry(n) &&
		pc.cols[pt.Col].isValidEntry(n) &&
		pc.boxes[pt.Box].isValidEntry(n)
}

func (pc *puzzleCache) validateDuplicates() []*InvalidSquareError {
	var errSet map[Point]*InvalidSquareError
	for i := 0; i < Dimension; i++ {
		for _, err := range pc.rows[i].getInvalidEntries() {
			if errSet == nil {
				errSet = make(map[Point]*InvalidSquareError)
			}
			errSet[newPoint(err.Row, err.Col)] = err
		}
		for _, err := range pc.cols[i].getInvalidEntries() {
			if errSet == nil {
				errSet = make(map[Point]*InvalidSquareError)
			}
			errSet[newPoint(err.Row, err.Col)] = err
		}
		for _, err := range pc.boxes[i].getInvalidEntries() {
			if errSet == nil {
				errSet = make(map[Point]*InvalidSquareError)
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

type pointCache map[int]map[Point]struct{}

func (pc pointCache) add(pt Point, n int) {
	ptSet, ok := pc[n]
	if !ok {
		ptSet = make(map[Point]struct{})
		pc[n] = ptSet
	}
	ptSet[pt] = struct{}{}
}

func (pc pointCache) remove(pt Point, n int) {
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
			errs = append(errs, newInvalidSquareError(pt.Row, pt.Col, duplicateNumber))
		}
	}
	return errs
}
