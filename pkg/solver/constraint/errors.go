package constraint

import "github.com/cszczepaniak/sudoku-solver/pkg/solver/model"

type ValidationError struct {
	Points  []model.Point `json:"points,omitempty"`
	Message string        `json:"message,omitempty"`
}

func (v *ValidationError) Error() string {
	return v.Message
}

type AggregateValidationError struct {
	points map[model.Point]struct{}
	msg    string
}

func NewAggregateValidationError() *AggregateValidationError {
	return &AggregateValidationError{
		points: make(map[model.Point]struct{}),
	}
}

func (e *AggregateValidationError) Add(err error) error {
	if err == nil {
		return nil
	}
	verr, ok := err.(*ValidationError)
	if !ok {
		return err
	}
	if len(e.points) == 0 {
		e.msg = err.Error()
	} else if e.msg != err.Error() {
		e.msg = `multiple validation errors occurred`
	}
	for _, pt := range verr.Points {
		e.points[pt] = struct{}{}
	}
	return nil
}

func (e *AggregateValidationError) Error() string {
	return e.msg
}

func (e *AggregateValidationError) ToValidationError() error {
	if e == nil || len(e.points) == 0 {
		// one of go's quirks: returning a nil pointer here will cause err != nil for the caller,
		// but that's avoided if we return nil as a literal here instead
		return nil
	}
	verr := &ValidationError{
		Message: e.Error(),
	}
	for pt := range e.points {
		verr.Points = append(verr.Points, pt)
	}
	return verr
}
