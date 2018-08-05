// 1 august 2018
package db

import (
	"github.com/andlabs/mdlist/fuzzytime"
)

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
	ID				ID
	Game			ID
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
