// 4 august 2018
package fuzzytime

import (
	"time"
	"testing"
)

func TestPackage(t *testing.T) {
	for _, tc := range []struct {
		str			string
		invalid		bool
		year			int
		month		time.Month
		day			int
		accuracy		Accuracy
	}{
		{
			str:			"",
			invalid:		true,
		},
		{
			str:			"abc",
			invalid:		true,
		},
		{
			str:			"????",
			invalid:		true,
		},
		{
			str:			"19??",
			year:			1900,
			month:		1,
			day:			1,
			accuracy:		Know20XX,
		},
		{
			str:			"19XX",
			invalid:		true,
		},
		{
			str:			"19?2",
			invalid:		true,
		},
		{
			str:			"19??-08",
			invalid:		true,
		},
		{
			str:			"199?",
			year:			1990,
			month:		1,
			day:			1,
			accuracy:		Know200X,
		},
		{
			str:			"199X",
			invalid:		true,
		},
		{
			str:			"199?-08",
			invalid:		true,
		},
		{
			str:			"1992",
			year:			1992,
			month:		1,
			day:			1,
			accuracy:		Know2006,
		},
		{
			str:			"1992-08",		// month < 10
			year:			1992,
			month:		8,
			day:			1,
			accuracy:		Know200601,
		},
		{
			str:			"1992-8",
			invalid:		true,
		},
		{
			str:			"1992/08",
			invalid:		true,
		},
		{
			str:			"1992/8",
			invalid:		true,
		},
		{
			str:			"1992.08",
			invalid:		true,
		},
		{
			str:			"1992.8",
			invalid:		true,
		},
		{
			str:			"1992年08月",
			invalid:		true,
		},
		{
			str:			"1992年8月",
			invalid:		true,
		},
		{
			str:			"1992/08",
			invalid:		true,
		},
		{
			str:			"08/1992",
			invalid:		true,
		},
		{
			str:			"92/8",
			invalid:		true,
		},
		{
			str:			"8/92",
			invalid:		true,
		},
		{
			str:			"1992-12",
			year:			1992,
			month:		12,
			day:			1,
			accuracy:		Know200601,
		},
		{
			str:			"1992-01-01",		// both < 10
			year:			1992,
			month:		1,
			day:			1,
			accuracy:		Know20060102,
		},
		{
			str:			"1992-1-01",
			invalid:		true,
		},
		{
			str:			"1992-01-1",
			invalid:		true,
		},
		{
			str:			"1992-1-1",
			invalid:		true,
		},
		{
			str:			"1992/01/01",
			invalid:		true,
		},
		{
			str:			"1992/1/01",
			invalid:		true,
		},
		{
			str:			"1992/01/1",
			invalid:		true,
		},
		{
			str:			"1992/1/1",
			invalid:		true,
		},
		{
			str:			"1992.01.01",
			invalid:		true,
		},
		{
			str:			"1992.1.01",
			invalid:		true,
		},
		{
			str:			"1992.01.1",
			invalid:		true,
		},
		{
			str:			"1992.1.1",
			invalid:		true,
		},
		{
			str:			"1992年01月01日",
			invalid:		true,
		},
		{
			str:			"1992年1月01日",
			invalid:		true,
		},
		{
			str:			"1992年01月1日",
			invalid:		true,
		},
		{
			str:			"1992年1月1日",
			invalid:		true,
		},
		{
			str:			"1992/01/01",
			invalid:		true,
		},
		{
			str:			"01/01/1992",
			invalid:		true,
		},
		{
			str:			"92/1/1",
			invalid:		true,
		},
		{
			str:			"1/1/92",
			invalid:		true,
		},
		{
			str:			"1992-08-15",		// month < 10
			year:			1992,
			month:		8,
			day:			15,
			accuracy:		Know20060102,
		},
		{
			str:			"1992-8-15",
			invalid:		true,
		},
		{
			str:			"1992-08.15",
			invalid:		true,
		},
		{
			str:			"1992-08-a5",
			invalid:		true,
		},
		{
			str:			"1992-08-1b",
			invalid:		true,
		},
		{
			str:			"1992-11-02",		// day < 10
			year:			1992,
			month:		11,
			day:			2,
			accuracy:		Know20060102,
		},
		{
			str:			"1992-11-2",
			invalid:		true,
		},
		{
			str:			"1992-12-31",
			year:			1992,
			month:		12,
			day:			31,
			accuracy:		Know20060102,
		},
	} {
		t.Run(tc.str, func(t *testing.T) {
			ft, err := Parse(tc.str)
			if tc.invalid {
				if err == nil {
					t.Fatalf("Parse(): got no error, want error")
				}
				return
			}
			if err != nil {
				t.Fatalf("Parse(): got error %v, want no error", err)
			}
			y, m, d := ft.Date()
			if y != tc.year || m != tc.month || d != tc.day {
				t.Errorf("ft.Date(): got %d-%02d-%02d, want %d-%02d-%02d", y, m, d, tc.year, tc.month, tc.day)
			}
			a := ft.Accuracy()
			if a != tc.accuracy {
				t.Errorf("ft.Accuracy(): got %v, want %v", a, tc.accuracy)
			}
			str := ft.String()
			if str != tc.str {
				t.Errorf("ft.String(): got %q, want %q", str, tc.str)
			}
		})
	}
}
