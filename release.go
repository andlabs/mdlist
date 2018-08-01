// 1 august 2018
package mdlist

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
	ReleaseDate		FuzzyDate
	ReleaseDateSource	string
	BuildDate			FuzzyTime
}
