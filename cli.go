// 29 july 2018
package main

import (
	"fmt"
	"os"
)

type command struct {
	Name		string
	Usage		string
	Description	string
	Fn			interface{}
}

func newGameDB(gameDBName string) {
	db := NewGameDatabase()
	f, err := os.OpenFile(gameDBName, os.O_WRONLY | os.O_CREATE | os.O_EXCL, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating game DB: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	err = db.Write(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing game DB: %v\n", err)
		os.Exit(1)
	}
}

var cmdNewGameDB = command{
	Name:		"newgamedb",
	Usage:		"newgamedb",
	Description:	"Creates a blank game database at the specified file, which must not exist.",
	Fn:			newGameDB,
}

var commands = []command{
	cmdNewGameDB,
	cmdAddGame,
	cmdListGames,
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s gamedb cmd [args...]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "known commands:\n")
	for _, c := range commands {
		fmt.Fprintf(os.Stderr, "  %s\n", c.Usage)
		fmt.Fprintf(os.Stderr, "    \t%s\n", c.Description)
	}
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 {
		usage()
	}
	gamedbname := os.Args[1]
	command := os.Args[2]
	for _, c := range commands {
		if c.Name == command {
			// TODO
			return
		}
	}
	fmt.Fprintf(os.Stderr, "error: unknown command %q\n", command)
	usage()
}
