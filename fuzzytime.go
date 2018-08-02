// 2 august 2018
package mdlist

import (
	"time"
)

type FuzzyTime struct {
	Time		time.TIme
	Accuracy	time.Time
}

type timeComponents struct {
	year		int
	month	time.Month
	day		int
	hour		int
	minute	int
	second	int
	ns		int
	loc		*time.Location
}

func fromTimeComponents(tc timeComponents) time.Time {
	return time.Date(tc.year, tc.month, tc.day,
		tc.hour, tc.minute, tc.second,
		tc.ns, tc.loc)
}

func toTimeComponents(t time.Time) timeComponents {
	tc := timeComponents{}
	tc.year, tc.month, tc.day = t.Date()
	tc.hour, tc.minute, tc.second = t.Clock()
	tc.ns = t.Nanosecond()
	tc.loc = t.Location()
	return tc
}

func normalizeTimePiece(cont bool, val int, acc int) (int, bool) {
	if !cont {
		return val, false
	}
	if acc == 0 {
		return 0, true
	}
	mul := 1
	for acc != 1 {
		val /= 10
		mul *= 10
		acc /= 10
	}
	return val * mul, false
}

func normalizeTimeMonth(cont bool, val time.Month, acc time.Month) (time.Month, bool) {
	norm, cont := normalizeTimeMonth(cont, int(val), int(acc))
	return time.Month(norm), cont
}

func (f FuzzyTime) Normalized() FuzzyTime {
	tc := toTimeComponents(f.Time)
	ac := toTimeComponents(f.Accuracy)
	cont := true
	tc.ns, cont = normalizeTimePiece(cont, tc.ns, ac.ns)
	tc.second, cont = normalizeTimePiece(cont, tc.second, ac.second)
	tc.minute, cont = normalizeTimePiece(cont, tc.minute, ac.minute)
	tc.hour, cont = normalizeTimePiece(cont, tc.hour, ac.hour)
	tc.day, cont = normalizeTimePiece(cont, tc.day, ac.day)
	tc.month, cont = normalizeTimeMonth(cont, tc.month, ac.month)
	tc.year, cont = normalizeTimePiece(cont, tc.year, ac.year)
	return fromTimeComponents(tc)
}

func (f FuzzyTime) String() string {
	f = f.Normalized()
}
