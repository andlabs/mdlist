// 28 july 2018
package main

import (
	"fmt"
)

type GameType int
const (
	MegaDriveHardware GameType = iota
	MegaDriveGame
	MegaCDHardware
	MegaCDGame
	Mega32XHardware
	Mega32XGame
	PicoHardware
	PicoGame
	ArcadeHardware
	ArcadeGame
	Accessory
	nGameTypes
)

var gameTypeUserStrings = map[GameType]string{
	MegaDriveHardware:	"mdhw",
	MegaDriveGame:		"md",
	MegaCDHardware:		"mcdhw",
	MegaCDGame:			"mcd",
	Mega32XHardware:		"32xhw",
	Mega32XGame:		"32x",
	PicoHardware:			"picohw",
	PicoGame:			"pico",
	ArcadeHardware:		"arcadehw",
	ArcadeGame:			"arcade",
	Accessory:			"accessory",
}

var gameTypeStrings = map[GameType]string{
	MegaDriveHardware:	"Mega Drive hardware",
	MegaDriveGame:		"Mega Drive game",
	MegaCDHardware:		"Mega CD hardware",
	MegaCDGame:			"Mega CD game",
	Mega32XHardware:		"32X hardware",
	Mega32XGame:		"32X game",
	PicoHardware:			"Pico hardware",
	PicoGame:			"Pico game",
	ArcadeHardware:		"arcade hardware",
	ArcadeGame:			"arcade game",
	Accessory:			"accessory",
}

func init() {
	if len(gameTypeUserStrings) != int(nGameTypes) {
		// TODO replace panic with something else
		panic("internal inconsistency: not all game types accounted for in gameTypeUserStrings")
	}
	if len(gameTypeStrings) != int(nGameTypes) {
		// TODO replace panic with something else
		panic("internal inconsistency: not all game types accounted for in gameTypeStrings")
	}
}

func GameTypeFromUser(user string) (g GameType, ok bool) {
	for g, s := range gameTypeUserStrings {
		if s == user {
			return g, true
		}
	}
	return 0, false
}

func GameTypeUserHelp() string {
	s := ""
	for g := GameType(0); g < nGameTypes; g++ {
		s += fmt.Sprintf("%10s - %s\n", gameTypeUserStrings[g], gameTypeStrings[g])
	}
	return s
}

func (g GameType) String() string {
	if s, ok := gameTypeStrings[g]; ok {
		return s
	}
	return fmt.Sprintf("<unknown game type %d>", int(g))
}

type Game struct {
	ID			ID
	DisplayName	string		// ASCII only
	Type			GameType
}

func main() {
	fmt.Println(GameTypeUserHelp())
}
