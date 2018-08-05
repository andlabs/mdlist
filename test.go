// 5 august 2018
package main

import (
	"os"

	"github.com/andlabs/mdlist/db"
	"github.com/andlabs/mdlist/fuzzytime"
)

func mustParse(s string) fuzzytime.FuzzyTime {
	f, _ := fuzzytime.Parse(s)
	return f
}

func main() {
	d := db.New()

	g, _ := d.AddGame(db.GameGame, db.MegaDrive)
	d.AddRelease(g,
		"スペースハリアーII", "Space Harrier II", "https://sega.jp/fb/segahard/md/soft.html",
		"Sega", "G-4002", db.Japan,
		false, false, false,
		mustParse("1988-10-29"), "https://sega.jp/fb/segahard/md/soft.html", fuzzytime.FuzzyTime{})

	d.Write(os.Stdout)
}
