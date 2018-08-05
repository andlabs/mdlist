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
	id				ID
	game			ID
	title				string	// TODO Name?
	romanized		string
	titleSource		string
	publisher			string
	serialNumber		string
	region			Region
	nonphysical		bool		// TODO NonPhysical? Digital?
	unlicensed		bool
	prototype			bool
	releaseDate		fuzzytime.FuzzyTime
	releaseDateSource	string
	buildDate			fuzzytime.FuzzyTime
	// TODO details for variant releases like Sega Classics? if so the asset map will need to change to allow a single ROM to map to multiple releases
	// TODO box type for arranging image assets?
}

func (r *Release) ID() ID {
	return r.id
}

func (r *Release) GameID() ID {
	return r.game
}

func (r *Release) Title() string {
	return r.title
}

func (r *Release) RomanizedTitle() string {
	return r.romanized
}

func (r *Release) TitleSource() string {
	return r.titleSource
}

func (r *Release) Publisher() string {
	return r.publisher
}

func (r *Release) SerialNumber() string {
	return r.serialNumber
}

func (r *Release) Region() Region {
	return r.region
}

func (r *Release) Nonphysical() bool {
	return r.nonphysical
}

func (r *Release) Unlicensed() bool {
	return r.unlicensed
}

func (r *Release) Prototype() bool {
	return r.prototype
}

func (r *Release) ReleaseDate() fuzzytime.FuzzyTime {
	return r.releeaseDate
}

func (r *Release) ReleaseDateSource() string {
	return r.releaseDateSource
}

func (r *Release) BuildDate() fuzzytime.FuzzyTime {
	return r.buildDate
}

var dbReleaseOps = interface {
	AddRelease(game ID, title string, romanized string, titleSource string, publisher string, serialNumber string, region Region, nonphysical bool, unlicensed bool, prototype bool, releaseDate fuzzytime.FuzzyTime, releaseDateSource string, buildDate fuzzytime.FuzzyTime) (*Release, error)
	EnumReleases(f func(r *Release) error) error
} = &DB{}

type releaseJSON struct {
	ID				ID
	Game			ID
	Title				string
	Romanized		string
	TitleSource		string
	Publisher			string
	SerialNumber		string
	Region			Region
	Nonphysical		bool
	Unlicensed		bool
	Prototype			bool
	ReleaseDate		string
	ReleaseDateSource	string
	BuildDate			string
}

func (r *Region) MarshalJSON() ([]byte, error) {
	return json.Marshal(releaseJSON{
		ID:				r.id,
		Game:			r.game,
		Title:				r.title,
		Romanized:		r.romanized,
		TitleSource:		r.titleSource,
		Publisher:			r.publisher,
		SerialNumber:		r.serialNumber,
		Region:			r.region,
		Nonphysical:		r.nonphysical,
		Unlicensed:		r.unlicensed,
		Prototype:			r.prototype,
		ReleaseDate:		r.releaseDate.String(),
		ReleaseDateSource:	r.releaseDateSource,
		BuildDate:			r.buildDate.String(),
	})
}

func (r *Release) UnmarshalJSON(b []byte) error {
	var x releaseJSON
	err := json.Unmarshal(b, &x)
	if err != nil {
		return err
	}
	r.id = x.ID
	r.game = x.Game
	r.title = x.Title
	r.romanized = x.Romanized
	r.titleSource = x.TitleSource
	r.publisher = x.Publisher
	r.serialNumber = x.SerialNumber
	r.region = x.Region
	r.nonphysical = x.Nonphysical
	r.unlicensed = x.Unlicensed
	r.prototype = x.Prototype
	r.releaseDate, err = fuzzytime.Parse(x.ReleaseDate)
	if err != nil {
		return err
	}
	r.releaseDateSource = x.ReleaseDateSource
	r.buildDate, err = fuzzytime.Parse(x.BuildDate)
	return err
}
