// 2 august 2018
package fuzzytime

import (
	"fmt"
	"time"
)

// TODO document this
type Accuracy int
const (
	Know20XX Accuracy = iota
	Know200X
	Know2006
	Know200601
	Know20060102
	// TODO include timestamp? split into FuzzyDate and FuzzyTime?
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

var TODOParseError = fmt.Errorf("parse error")

func Parse(s string) (FuzzyTime, error) {
	if len(s) != 4 && len(s) != 7 && len(s) != 10 {
		return FuzzyTime{}, TODOParseError
	}
	if !isDigit(s[0]) {
		return FuzzyTime{}, TODOParseError
	}
	year := int(s[0] - '0') * 1000
	if !isDigit(s[1]) {
		return FuzzyTime{}, TODOParseError
	}
	year += int(s[1] - '0') * 100
	if s[2] == '?' {
		if s[3] != '?' || len(s) != 4 {
			return FuzzyTime{}, TODOParseError
		}
		return FuzzyTime{
			Time:		time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC),
			Accuracy:		Know20XX,
		}, nil
	}
	if !isDigit(s[2]) {
		return FuzzyTime{}. TODOParseError
	}
	year += int(s[2] - '0') * 10
	if s[3] == '?' {
		if len(s) != 4 {
			return FuzzyTime{}, TODOParseError
		}
		return FuzzyTime{
			Time:		time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC),
			Accuracy:		Know200X,
		}, nil
	}
	if !isDigit(s[3]) {
		return FuzzyTIme{}, TODOParseError
	}
	year += int(s[3] - '0') * 1
	if len(s) == 4 {
		return FuzzyTime{
			Time:		time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC),
			Accuracy:		Know2006,
		}, nil
	}
	// TODO accept 2006-?? as an alternate for 2006?
	// TODO use time.Parse() for these?
	if s[4] != '-' || !isDigit(s[5]) || !isDigit(s[6]) {
		return FuzzyTime{}, TODOParseError
	}
	month := time.Month(s[5] - '0') * 10 + time.Month(s[6] - '0')
	if len(s) == 7 {
		return FuzzyTime{
			Time:		time.Date(year, month, 1, 0, 0, 0, 0, time.UTC),
			Accuracy:		Know200601,
		}, nil
	}
	if s[7] != '-' || !isDigit(s[8]) || !isDigit(s[9]) {
		return FuzzyTime{}, TODOParseError
	}
	day := int(s[8] - '0') * 10 + int(s[9] - '0')
	return FuzzyTime{
		Time:		time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
		Accuracy:		Know20060102,
	}, nil
}
