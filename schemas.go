// 28 july 2018
package mdlist

import (
	"golang.org/x/crypto/blake2b"
)

struct Game {
	ID			ID
	DisplayName	string		// ASCII only
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
	ID					ID
	Game				*Game
	Region				Region
	Name				string
	NameRomanized		string	// Japan, Korea, Asia only
	NameSource			string
	Publisher				string
	ReleaseDate			time.Time
	ReleaseDateAccuracy	time.Time		// each component: 1 in the most precise known digit (so 199?-??-?? means a ReleaseDate of 1990-1-1 and Accuracy of 10-0-0)
	ReleaseDateSource		string		// if unspecified (empty), marked as unconfirmed; accuracy should be to the year from the title screen of the game
	SerialNumber			string
	// TODO details for variant releases like Sega Classics? if so the asset map will need to change to allow a single ROM to map to multiple releases
	// TODO box type for arranging image assets?
}

type Asset struct {
	ID			ID
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
	Asset		*Asset		// must be unique
	Type			AssetType
	Details		string		// primary assets of the given type MUST leave this unspecified
}
