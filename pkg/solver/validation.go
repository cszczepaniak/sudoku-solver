package solver

import (
	"fmt"
)

type InvalidBoardError struct {
	Errors []*InvalidSquareError
}

func (ibe *InvalidBoardError) Error() string {
	return fmt.Sprintf(`invalid board: %d invalid squares`, len(ibe.Errors))
}

type InvalidSquareError struct {
	Row int
	Col int
	Msg string
}

func newInvalidSquareError(r, c int, reason invalidReason) *InvalidSquareError {
	return &InvalidSquareError{
		Row: r,
		Col: c,
		Msg: reasonToMsg[reason],
	}
}

func (ise *InvalidSquareError) Error() string {
	return fmt.Sprintf(`invalid square at (%d, %d): %s`, ise.Row, ise.Col, ise.Msg)
}

type invalidReason int

const (
	_ invalidReason = iota

	duplicateNumber
	outOfRange
)

var reasonToMsg = map[invalidReason]string{
	duplicateNumber: `duplicate number in row, column, or box`,
	outOfRange:      `number out of range`,
}
