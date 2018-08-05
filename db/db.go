// 5 august 2018
package db

import (
	"io"
	"sync"
	"encoding/json"
	"sort"

	"github.com/andlabs/mdlist/fuzzytime"
)

type ID int

type DB struct {
	mu				sync.RWMutex
	games			map[ID]*Game
	nextGame			ID
	releases			map[ID]*Release
	nextRelease		ID
//	assets			[]*Asset
//	assetMappings		[]*AssetMapping
}

type dbJSON struct {
	Games			[]*Game
	Releases			[]*Release
//	Assets			[]*Asset
//	AssetMappings		[]*AssetMapping
}

func Read(r io.Reader) (*DB, error) {
	x := dbJSON{}
	err := json.NewDecoder(r).Decode(&x)
	if err != nil {
		return nil, err
	}
	db := &DB{
		games:		make(map[ID]*Game, len(x.Games)),
		releases:		make(map[ID]*Release, len(x.Releases)),
	}
	for _, g := range x.Games {
		db.games[g.ID()] = g
	}
	for _, r := range x.Releases {
		db.releases[r.ID()] = r
	}
	return db, nil
}

func (db *DB) Write(w io.Writer) error {
	db.mu.RLock()
	defer db.mu.RUnlock()

	x := dbJSON{
		Games:		make([]*Game, 0, len(db.games)),
		Releases:		make([]*Release, 0, len(db.releases)),
	}
	for _, g := range db.games {
		x.Games = append(x.Games, g)
	}
	for _, r := range db.releases {
		x.Releases = append(x.Releases, r)
	}
	// and sort to make sure Write+Read+Write is idempotent
	sort.Slice(x.Games, func(i, j int) bool {
		return x.Games[i].ID() < x.Games[j].ID()
	})
	sort.Slice(x.Releases, func(i, j int) bool {
		return x.Releases[i].ID() < x.Releases[j].ID()
	})
	return json.NewEncoder(w).Encode(x)
}

func (db *DB) AddGame(ty GameType, platform Platform) (*Game, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	gameID := db.nextGame
	for {
		if _, ok := db.games[gameID]; !ok {
			break
		}
		gameID++
	}
	db.nextGame = gameID + 1
	g := &Game{
		id:			gameID,
		ty:			ty,
		platform:		platform,
	}
	db.games[gameID] = g
	return g, nil
}

func (db *DB) EnumGames(f func(g *Game) error) error {
	db.mu.RLock()
	defer db.mu.RUnlock()

	for _, g := range db.games {
		err := f(g)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *DB) AddRelease(game *Game, title string, romanized string, titleSource string, publisher string, serialNumber string, region Region, nonphysical bool, unlicensed bool, prototype bool, releaseDate fuzzytime.FuzzyTime, releaseDateSource string, buildDate fuzzytime.FuzzyTime) (*Release, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	panic("TODO")
}

func (db *DB) EnumReleases(f func(r *Release) error) error {
	db.mu.RLock()
	defer db.mu.RUnlock()

	for _, r := range db.releases {
		err := f(r)
		if err != nil {
			return err
		}
	}
	return nil
}
