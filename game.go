// 1 august 2018
package db

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
	InsertGame(ty GameType, platform Platform) (ID, error)
	LookupGame(id ID) (*Game, error)
	DeleteGame(id ID) error
	EnumGames(f func(*Game) error) error
} = &DB{}
