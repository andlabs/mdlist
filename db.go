// 5 august 2018
package db

import (
	"sync"
	"encoding/json"
)

type ID int

type DB struct {
	mu				sync.RWMutex
	games			[]*Game
	releases			[]*Release
	assets			[]*Asset
	assetMappings		[]*AssetMapping
}

func Read(r io.Reader) (*DB, error) {
	db := new(DB)
	err := json.NewDecoder(r).Decode(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (db *DB) Write(w io.Writer) error {
	return json.NewEncoder(w).Encode(db)
}

func (db *DB) AddGame(ty GameType, platform Platform) (ID, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	g := &Game{
		ty:			ty,
		platform:		platform,
	}
	db.games = append(db.games, g)
	return ID(len(db.games) - 1), nil
}

func (db *DB) EnumGames(f func(id ID, g *Game) error) error {
	db.mu.RLock()
	defer db.mu.RUnlock()
	for i, g := range db.games {
		err := f(ID(i), g)
		if err != nil {
			return err
		}
	}
	return nil
}
