package solver

import "github.com/cszczepaniak/sudoku-solver/pkg/solver/model"

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
	pt := model.NewPoint(r, c)
	pc.rows[pt.Row()].add(pt, n)
	pc.cols[pt.Col()].add(pt, n)
	pc.boxes[pt.Box()].add(pt, n)
}

func (pc *puzzleCache) remove(r, c, n int) {
	pt := model.NewPoint(r, c)
	pc.rows[pt.Row()].remove(pt, n)
	pc.cols[pt.Col()].remove(pt, n)
	pc.boxes[pt.Box()].remove(pt, n)
}

func (pc *puzzleCache) validateDuplicates() []*InvalidSquareError {
	var errSet map[model.Point]*InvalidSquareError
	for i := 0; i < Dimension; i++ {
		for _, err := range pc.rows[i].getInvalidEntries() {
			if errSet == nil {
				errSet = make(map[model.Point]*InvalidSquareError)
			}
			errSet[model.NewPoint(err.Row, err.Col)] = err
		}
		for _, err := range pc.cols[i].getInvalidEntries() {
			if errSet == nil {
				errSet = make(map[model.Point]*InvalidSquareError)
			}
			errSet[model.NewPoint(err.Row, err.Col)] = err
		}
		for _, err := range pc.boxes[i].getInvalidEntries() {
			if errSet == nil {
				errSet = make(map[model.Point]*InvalidSquareError)
			}
			errSet[model.NewPoint(err.Row, err.Col)] = err
		}
	}
	errs := make([]*InvalidSquareError, 0, len(errSet))
	for _, err := range errSet {
		errs = append(errs, err)
	}
	return errs
}

type pointCache map[int]map[model.Point]struct{}

func (pc pointCache) add(pt model.Point, n int) {
	ptSet, ok := pc[n]
	if !ok {
		ptSet = make(map[model.Point]struct{})
		pc[n] = ptSet
	}
	ptSet[pt] = struct{}{}
}

func (pc pointCache) remove(pt model.Point, n int) {
	delete(pc[n], pt)
}

func (pc pointCache) getInvalidEntries() []*InvalidSquareError {
	var errs []*InvalidSquareError
	for _, pts := range pc {
		if len(pts) <= 1 {
			// valid
			continue
		}
		for pt := range pts {
			errs = append(errs, newInvalidSquareError(pt.Row(), pt.Col(), duplicateNumber))
		}
	}
	return errs
}
