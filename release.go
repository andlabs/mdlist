// 1 august 2018
package mdlist

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
	Type		GameType
	Platform	Platform
}

type Region int
const (
	Japan Region = iota
	US
	Europe
	Brazil
	SouthKorea
	Australia
	AsiaNTSC
	AsiaPAL
)

type Release struct {
	Title				string	// TODO Name?
	Romanized		string
	TitleSource		string
	Publisher			string
	SerialNumber		string
	Region			Region
	Nonphysical		bool		// TODO NonPhysical? Digital?
	Unlicensed		bool
	Prototype			bool
	ReleaseDate		FuzzyTime
	ReleaseDateSource	string
	BuildDate			FuzzyTime
}
