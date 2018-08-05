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
	d.AddRelease(g,
		"Space Harrier II", "", "",
		"Sega", "", db.US,
		false, false, false,
		mustParse("1989"), "", fuzzytime.FuzzyTime{})

	g, _ = d.AddGame(db.GameGame, db.MegaDrive)
	d.AddRelease(g,
		"スーパーサンダーブレード", "Super Thunder Blade", "https://sega.jp/fb/segahard/md/soft.html",
		"Sega", "G-4003", db.Japan,
		false, false, false,
		mustParse("1988-10-29"), "https://sega.jp/fb/segahard/md/soft.html", fuzzytime.FuzzyTime{})

	g, _ = d.AddGame(db.GameGame, db.MegaDrive)
	d.AddRelease(g,
		"獣王記", "Juuouki", "https://sega.jp/fb/segahard/md/soft.html",
		"Sega", "G-4001", db.Japan,
		false, false, false,
		mustParse("1988-11-27"), "https://sega.jp/fb/segahard/md/soft.html", fuzzytime.FuzzyTime{})
	d.AddRelease(g,
		"Altered Beast", "", "",
		"Sega", "", db.US,
		false, false, false,
		mustParse("1989"), "", fuzzytime.FuzzyTime{})
	d.AddRelease(g,
		"Altered Beast", "", "",
		"Sega", "", db.Europe,
		false, false, false,
		mustParse("19??"), "", fuzzytime.FuzzyTime{})

	g, _ = d.AddGame(db.GameGame, db.MegaDrive)
	d.AddRelease(g,
		"ワールドカップサッカー", "World Cup Soccer", "https://sega.jp/fb/segahard/md/soft.html",
		"Sega", "G-4009", db.Japan,
		false, false, false,
		mustParse("1989-07-29"), "https://sega.jp/fb/segahard/md/soft.html", fuzzytime.FuzzyTime{})
	d.AddRelease(g,
		"World Championship Soccer", "", "",
		"Sega", "", db.US,
		false, false, false,
		mustParse("1989"), "", fuzzytime.FuzzyTime{})

	d.Write(os.Stdout)
}
