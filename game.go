// 1 august 2018
package db

import (
	"encoding/json"
)

type GameType int
const (
	Hardware		GameType = iota
	Game
	Accessory
)

type Platform int
const (
	MegaDrive Platform = iota
	MegaCD
	Mega32X
	Pico
	Arcade
	Teradrive
)

type Game struct {
	id		ID
	ty		GameType
	platform	Platform
}

func (g *Game) ID() ID {
	return g.id
}

func (g *Game) Type() GameType {
	return g.ty
}

func (g *Game) Platform() Platform {
	return g.platform
}

var dbGameOps interface {
	AddGame(ty GameType, platform Platform) (*Game, error)
	EnumGames(f func(g *Game) error) error
} = &DB{}

type gameJSON struct {
	ID		ID
	Type		GameType
	Platform	Platform
}

func (g *Game) MarshalJSON() ([]byte, error) {
	return json.Marshal(gameJSON{
		ID:		g.id,
		Type:	g.ty,
		Platform:	g.platform,
	})
}

func (g *Game) UnmarshalJSON(b []byte) error {
	var x gameJSON
	err := json.Unmarshal(b, &x)
	if err != nil {
		return err
	}
	g.id = x.ID
	g.ty = x.Type
	g.platform = x.Platform
	return nil
}
