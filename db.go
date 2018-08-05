// 5 august 2018
package db

import (
	"encoding/json"
)

type DB struct {
	games			[]*Game
	Releases			[]*Release
	Assets			[]*Asset
	AssetMappings		[]*AssetMapping
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
