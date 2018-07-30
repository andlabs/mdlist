// 28 july 2018
package main

import (
	"fmt"
	"io"
	"encoding/json"
	"sync"
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
	Name		string		// ASCII only
	Year			int			// earliest year of actual release or TODO in the case of an unreleased game
	// TODO YearAccuracy?
	Type			GameType
}

type GameDatabase struct {
	mu		sync.RWMutex
	games	map[ID]*Game
	pool		*IDPool
}

func NewGameDatabase() *GameDatabase {
	g := new(GameDatabase)
	g.games = make(map[ID]*Game)
	g.pool = NewIDPool()
	return g
}

func ReadGameDatabase(r io.Reader) (*GameDatabase, error) {
	g := new(GameDatabase)
	err := json.NewDecoder(r).Decode(&(g.games))
	if err != nil {
		return nil, err
	}
	g.pool = NewIDPool()
	for id, _ := range g.games {
		g.pool.Mark(id)
	}
	return g, nil
}

func (g *GameDatabase) Add(name string, year int, ty GameType) *Game {
	g.mu.Lock()
	defer g.mu.Unlock()
	game := &Game{
		Name:	name,
		Year:		year,
		Type:	ty,
	}
	game.ID = g.pool.Next()
	g.games[game.ID] = game
	return game
}

func (g *GameDatabase) Lookup(id ID) *Game {
	g.mu.RLock()
	defer g.mu.RUnlock()
	ret, _ := g.games[id]		// comma-ok syntax to avoid creating nil entries
	return ret
}

func (g *GameDatabase) ForEach(f func(game *Game)) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	for _, game := range g.games {
		f(game)
	}
}

func (g *GameDatabase) Write(w io.Writer) error {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return json.NewEncoder(w).Encode(g.games)
}

func main() {
	g := NewGameDatabase()
	g.Add("Mega Drive", 1988, MegaDriveHardware)
	g.Add("Space Harrier II", 1988, MegaDriveGame)
	g.Add("Super Thunder Blade", 1988, MegaDriveGame)
	g.Add("Altered Beast", 1988, MegaDriveGame)
	g.Add("Osomatsu-kun: Hachamecha Gekijou", 1988, MegaDriveGame)
	g.Add("Alex Kidd in the Enchanted Castle", 1989, MegaDriveGame)
	g.Add("Phantasy Star II", 1989, MegaDriveGame)
	g.ForEach(func(game *Game) {
		fmt.Printf("%d %-50s %v\n", game.Year, game.Name, game.Type)
	})
}
