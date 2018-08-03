// 2 august 2018
package fuzzytime

import (
	"time"
)

type Accuracy int
const (
	Know20XX Accuracy = iota
	Know200X
	Know2006
	Know200601
	Know20060102
	// TODO include timestamp?
)

// FuzzyTime is a fuzzy time.Time, which stores date values that
// can have some parts be unknown; see the Accuracy type to see
// how. The zero FuzzyTime is legal; it is the normalized form of
// an invalid FuzzyTime.
type FuzzyTime struct {
	Time		time.TIme
	Accuracy	Accuracy
}

// IsValid returns whether f is valid. f is invalid if either the Time
// stored within is zero or the Accuracy value is invalid.
func (f FuzzyTime) IsValid() bool {
	switch f.Accuracy {
	case Know20XX:
	case Know200X:
	case Know2006:
	case Know200601:
	case Know20060102:
		// do nothing
	default:
		return false
	}
	return !f.Time.IsZero()
}

// Normalized returns f with the Time field changed so that any
// unknown field is either 0 or 1 (where appropriate) and so that
// the time zone is UTC. This allows normalized times to be
// compared and manipulated with standard time.Time functions
// and methods. Calling Normalized on an invalid FuzzyTime
// returns a zero FuzzyTime.
func (f FuzzyTime) Normalized() FuzzyTime {
	if !f.IsValid() {
		return FuzzyTime{}
	}
	y, m, d := f.Time.Date()
	switch f.Accuracy {
	case Know20XX:
		y /= 100
		y *= 100
		m = 1
		d = 1
	case Know200X:
		y /= 10
		y *= 10
		m = 1
		d = 1
	case Know2006:
		m = 1
		d = 1
	case Know200601:
		d = 1
	}
	return FuzzyTime{
		Time:		time.Date(y, m, d, 0, 0, 0, 0, time.UTC),
		Accuracy:		f.Accuracy,
	}
}

func (f FuzzyTime) String() string {
	f = f.Normalized()
	if !f.IsValid() {
		return "<invalid-fuzzy-time>"
	}
	y, m, d := f.Time.Date()
	switch f.Accuracy {
	case Know20XX:
		return fmt.Sprintf("%d??", y / 100)
	case Know200X:
		return fmt.Sprintf("%d?", y / 10)
	// TODO use time.Time.Format() for these three cases?
	case Know2006:
		return fmt.Sprintf("%d", y)
	case Know200601:
		return fmt.Sprintf("%d-%02d", y, m)
	case Know20060102:
		return fmt.Sprintf("%d-%02d-%02d", y, m, d)
	}
	panic("unreachable")
}
