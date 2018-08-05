// 1 august 2018
package mdlist

import (
	"github.com/andlabs/mdlist/fuzzytime"
	"golang.org/x/crypto/blake2b"
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
	Game			*Game
	Title				string	// TODO Name?
	Romanized		string
	TitleSource		string
	Publisher			string
	SerialNumber		string
	Region			Region
	Nonphysical		bool		// TODO NonPhysical? Digital?
	Unlicensed		bool
	Prototype			bool
	ReleaseDate		fuzzytime.FuzzyTime
	ReleaseDateSource	string
	BuildDate			fuzzytime.FuzzyTime
	// TODO details for variant releases like Sega Classics? if so the asset map will need to change to allow a single ROM to map to multiple releases
	// TODO box type for arranging image assets?
}

type Asset struct {
	Size			int64
	MIME		string
	Blake2b512	[blake2b.Size]byte
	Source		string
	BadReason	string	// if unspecified, asset is considered "good"
}

type AssetType int
const (
	Other AssetType = iota
	ROM
	Cover
	Spine
	SpineTop
	SpineRight
	SpineBottom
	BackCover
	Media
	MediaLeft
	MediaTop
	MediaRight
	MediaBottom
	MediaBack
	PCB
	PCBBack
	Manual
)

type AssetMapping struct {
	Release		*Release
	Asset		*Asset
	Type			AssetType
	Details		string		// primary assets of the given type MUST leave this unspecified
}

type DB struct {
	Games			[]*Game
	Releases			[]*Release
	Assets			[]*Asset
	AssetMappings		[]*AssetMapping
}
