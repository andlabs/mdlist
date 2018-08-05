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
	ty		GameType
	platform	Platform
}

func (g *Game) Type() GameType {
	return g.ty
}

func (g *Game) Platform() Platform {
	return g.platform
}

var dbGameOps interface {
	AddGame(ty GameType, platform Platform) (ID, error)
	EnumGames(f func(id ID, g *Game) error) error
} = &DB{}

func (g *Game) MarshalJSON() ([]byte, error) {
	x := struct {
		Type		GameType
		Platform	Platform
	}{
		Type:	g.ty,
		Platform:	g.platform,
	}
	return json.Marshal(x)
}

func (g *Game) UnmarshalJSON(b []byte) error {
	var x struct {
		Type		GameType
		Platform	Platform
	}
	err := json.Unmarshal(b, &x)
	if err != nil {
		return err
	}
	g.ty = x.Type
	g.platform = x.Platform
	return nil
}
