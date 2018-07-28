-- 28 july 2018

CREATE TABLE Game (
	GUID		{{.GUID}},
	DisplayName	{{.String}},		-- ASCII only
);

{{Enum "Region"}}
	{{EnumCase "Japan"}}
	{{EnumCase "US"}}
	{{EnumCase "Europe"}}
	{{EnumCase "Brazil"}}
	{{EnumCase "SouthKorea"}}
	{{EnumCase "Australia"}}
	{{EnumCase "AsiaNTSC"}}
	{{EnumCase "AsiaPAL"}}
{{EndEnum}}

CREATE TABLE Release (
	GUID			{{.GUID}},
	Game			{{.GUID}} {{.ToOne "Game"}},
	Region			{{.Enum "Region"}},
	Name			{{.String}},
	NameRomanized	{{.String}},		-- Japan, Korea, Asia only
	NameSource		{{.String}},
	Publisher			{{.String}},
	ReleaseDate		{{.DateNoTime}},
	ReleaseDateAccuracy	{{.DateNoTime}},	-- each component: 1 in the most precise known digit (so 199?-??-?? means a ReleaseDate of 1990-1-1 and Accuracy of 10-0-0)
	ReleaseDateSource	{{.String}},		-- if unspecified (empty), marked as unconfirmed; accuracy should be to the year from the title screen of the game
	SerialNumber		{{.String}},
	-- TODO details for variant releases like Sega Classics? if so the asset map will need to change to allow a single ROM to map to multiple releases
	-- TODO box type for arranging image assets?
);

CREATE TABLE Asset (
	GUID		{{.GUID}},
	Size			{{.Int64}},
	MIME		{{.String}},
	Blake2b512	{{.Bytes .Const.Blake2b512Size}},
	Source		{{.String}},
	BadReason	{{.String}},		-- if unspecified, asset is considered "good"
);

{{Enum "AssetType"}}
	{{EnumCase "ROM"}}
	{{EnumCase "Cover"}}
	{{EnumCase "Spine"}}
	{{EnumCase "SpineTop"}}
	{{EnumCase "SpineRight"}}
	{{EnumCase "SpineBottom"}}
	{{EnumCase "BackCover"}}
	{{EnumCase "Manual"}}
	{{EnumCase "Other"}}
[{EndEnum}}

CREATE TABLE AssetMap (
	Release		{{.GUID}} {{.ToOne "Release"}},
	Asset		{{.GUID}} {{.ToOne "Asset"}} {{.Unique}},
	Type			{{.Enum "AssetType"}},
	Details		{[.String}},		-- primary assets of the given type MUST leave this unspecified
);
