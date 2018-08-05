// 2 august 2018
package fuzzytime

// TODO document all the exported functions and types

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
	// TODO include timestamp?
)

// FuzzyTime is a fuzzy time.Time, which stores date values that
// can have some parts be unknown; see the Accuracy type to see
// how. The zero FuzzyTime is invalid and produces undefined
// behavior (TODO).
type FuzzyTime struct {
	time		time.Time
	accuracy	Accuracy
}

var TODOParseError = fmt.Errorf("parse error")

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

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
			time:			time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC),
			accuracy:		Know20XX,
		}, nil
	}
	if !isDigit(s[2]) {
		return FuzzyTime{}, TODOParseError
	}
	year += int(s[2] - '0') * 10
	if s[3] == '?' {
		if len(s) != 4 {
			return FuzzyTime{}, TODOParseError
		}
		return FuzzyTime{
			time:			time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC),
			accuracy:		Know200X,
		}, nil
	}
	if !isDigit(s[3]) {
		return FuzzyTime{}, TODOParseError
	}
	year += int(s[3] - '0') * 1
	if len(s) == 4 {
		return FuzzyTime{
			time:			time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC),
			accuracy:		Know2006,
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
			time:			time.Date(year, month, 1, 0, 0, 0, 0, time.UTC),
			accuracy:		Know200601,
		}, nil
	}
	if s[7] != '-' || !isDigit(s[8]) || !isDigit(s[9]) {
		return FuzzyTime{}, TODOParseError
	}
	day := int(s[8] - '0') * 10 + int(s[9] - '0')
	return FuzzyTime{
		time:			time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
		accuracy:		Know20060102,
	}, nil
}

// Date returns the year, month, and day stored in f.
// Digits of the year not provided by the accuracy of f are 0.
// Months and days not provided by the accuracy of f are 1.
func (f FuzzyTime) Date() (year int, month time.Month, day int) {
	return f.time.Date()
}

// Accuracy returns the accuracy of f.
func (f FuzzyTime) Accuracy() Accuracy {
	return f.accuracy
}

func (f FuzzyTime) String() string {
	y, m, d := f.time.Date()
	switch f.accuracy {
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
