// 4 august 2018
package fuzzytime

import (
	"time"
	"testing"
)

func TestParse(t *testing.T) {
	// TODO
}

func TestIsValid(t *testing.T) {
	for _, tc := range []struct {
		name	string
		ft		FuzzyTime
		valid		bool
	}{
		{
			name:	"valid 20?? time",
			ft:		FuzzyTime{
				Time:		time.Date(1992, 8, 15, 0, 0, 0, 0, time.UTC),
				Accuracy:		Know20XX,
			},
			valid:	true,
		},
		{
			name:	"invalid 20?? time",
			ft:		FuzzyTime{
				Time:		time.Time{},
				Accuracy:		Know20XX,
			},
			valid:	false,
		},
		{
			name:	"valid 200? time",
			ft:		FuzzyTime{
				Time:		time.Date(1992, 8, 15, 0, 0, 0, 0, time.UTC),
				Accuracy:		Know200X,
			},
			valid:	true,
		},
		{
			name:	"invalid 200? time",
			ft:		FuzzyTime{
				Time:		time.Time{},
				Accuracy:		Know200X,
			},
			valid:	false,
		},
		{
			name:	"valid 2006 time",
			ft:		FuzzyTime{
				Time:		time.Date(1992, 8, 15, 0, 0, 0, 0, time.UTC),
				Accuracy:		Know2006,
			},
			valid:	true,
		},
		{
			name:	"invalid 2006 time",
			ft:		FuzzyTime{
				Time:		time.Time{},
				Accuracy:		Know2006,
			},
			valid:	false,
		},
		{
			name:	"valid 2006-01 time",
			ft:		FuzzyTime{
				Time:		time.Date(1992, 8, 15, 0, 0, 0, 0, time.UTC),
				Accuracy:		Know200601,
			},
			valid:	true,
		},
		{
			name:	"invalid 2006-01 time",
			ft:		FuzzyTime{
				Time:		time.Time{},
				Accuracy:		Know200601,
			},
			valid:	false,
		},
		{
			name:	"valid 2006-01-02 time",
			ft:		FuzzyTime{
				Time:		time.Date(1992, 8, 15, 0, 0, 0, 0, time.UTC),
				Accuracy:		Know20060102,
			},
			valid:	true,
		},
		{
			name:	"invalid 2006-01-02 time",
			ft:		FuzzyTime{
				Time:		time.Time{},
				Accuracy:		Know20060102,
			},
			valid:	false,
		},
		{
			name:	"invalid accuracy with valid time.Time",
			ft:		FuzzyTime{
				Time:		time.Date(1992, 8, 15, 0, 0, 0, 0, time.UTC),
				Accuracy:		12345678,
			},
			valid:	false,
		},
		{
			name:	"invalid accuracy with invalid time.Time",
			ft:		FuzzyTime{
				Time:		time.Time{},
				Accuracy:		12345678,
			},
			valid:	false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			gotValid := tc.ft.IsValid()
			if gotValid != tc.valid {
				t.Errorf("IsValid(%#v) failed: got %v, want %v", tc.ft, gotValid, tc.valid)
			}
		})
	}
}

func TestNormalized(t *testing.T) {
	// TODO
}

func TestString(t *testing.T) {
	for _, tc := range []struct {
		name	string
		ft		FuzzyTime
		str		string
	}{
		{
			name:	"invalid",
			ft:		FuzzyTime{
				Time:		time.Time{},
				Accuracy:		Know20XX,
			},
			str:		"<invalid-fuzzy-time>",
		},
		{
			name:	"20?? format",
			ft:		FuzzyTime{
				Time:		time.Date(1992, 8, 15, 0, 0, 0, 0, time.UTC),
				Accuracy:		Know20XX,
			},
			str:		"19??",
		},
		{
			name:	"200? format",
			ft:		FuzzyTime{
				Time:		time.Date(1992, 8, 15, 0, 0, 0, 0, time.UTC),
				Accuracy:		Know200X,
			},
			str:		"199?",
		},
		{
			name:	"2006 format",
			ft:		FuzzyTime{
				Time:		time.Date(1992, 8, 15, 0, 0, 0, 0, time.UTC),
				Accuracy:		Know2006,
			},
			str:		"1992",
		},
		{
			name:	"2006-01 format",
			ft:		FuzzyTime{
				Time:		time.Date(1992, 12, 31, 0, 0, 0, 0, time.UTC),
				Accuracy:		Know200601,
			},
			str:		"1992-12",
		},
		{
			name:	"2006-01 format, month < 10",
			ft:		FuzzyTime{
				Time:		time.Date(1992, 8, 15, 0, 0, 0, 0, time.UTC),
				Accuracy:		Know200601,
			},
			str:		"1992-08",
		},
		{
			name:	"2006-01-02 format",
			ft:		FuzzyTime{
				Time:		time.Date(1992, 12, 31, 0, 0, 0, 0, time.UTC),
				Accuracy:		Know20060102,
			},
			str:		"1992-12-31",
		},
		{
			name:	"2006-01-02 format, month < 10",
			ft:		FuzzyTime{
				Time:		time.Date(1992, 8, 15, 0, 0, 0, 0, time.UTC),
				Accuracy:		Know20060102,
			},
			str:		"1992-08-15",
		},
		{
			name:	"2006-01-02 format, day < 10",
			ft:		FuzzyTime{
				Time:		time.Date(1992, 11, 2, 0, 0, 0, 0, time.UTC),
				Accuracy:		Know20060102,
			},
			str:		"1992-11-02",
		},
		{
			name:	"2006-01-02 format, month < 10, day < 10",
			ft:		FuzzyTime{
				Time:		time.Date(1992, 1, 1, 0, 0, 0, 0, time.UTC),
				Accuracy:		Know20060102,
			},
			str:		"1992-01-01",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.ft.String()
			if got != tc.str {
				t.Errorf("String(%#v) failed: got %q, want %q", tc.ft, got, tc.str)
			}
		})
	}
}