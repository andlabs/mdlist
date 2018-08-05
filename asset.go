// 1 august 2018
package db

import (
	"golang.org/x/crypto/blake2b"
)

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
	Release		ID
	Asset		ID
	Type			AssetType
	Details		string		// primary assets of the given type MUST leave this unspecified
}
