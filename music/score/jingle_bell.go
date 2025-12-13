package score

import (
	"github.com/masatomo57/golang-playground/music/accompaniment"
	"github.com/masatomo57/golang-playground/music/conf"
	"github.com/masatomo57/golang-playground/music/melody"
)

var JingleBellMelody = melody.Melody{
	// Jingle bells, jingle bells
	{Note: conf.E5, Length: 1},
	{Note: conf.E5, Length: 1},
	{Note: conf.E5, Length: 2},
	{Note: conf.E5, Length: 1},
	{Note: conf.E5, Length: 1},
	{Note: conf.E5, Length: 2},
	// Jingle all the way
	{Note: conf.E5, Length: 1},
	{Note: conf.G5, Length: 1},
	{Note: conf.C5, Length: 1.5},
	{Note: conf.D5, Length: 0.5},
	{Note: conf.E5, Length: 4},
	// Oh what fun it is to ride
	{Note: conf.F5, Length: 1},
	{Note: conf.F5, Length: 1},
	{Note: conf.F5, Length: 1.5},
	{Note: conf.F5, Length: 0.5},
	{Note: conf.F5, Length: 1},
	{Note: conf.E5, Length: 1},
	{Note: conf.E5, Length: 1},
	{Note: conf.E5, Length: 1},
	// In a one-horse open sleigh
	{Note: conf.E5, Length: 1},
	{Note: conf.D5, Length: 1},
	{Note: conf.D5, Length: 1},
	{Note: conf.E5, Length: 1},
	{Note: conf.D5, Length: 2},
	{Note: conf.G5, Length: 2},
	// Jingle bells, jingle bells
	{Note: conf.E5, Length: 1},
	{Note: conf.E5, Length: 1},
	{Note: conf.E5, Length: 2},
	{Note: conf.E5, Length: 1},
	{Note: conf.E5, Length: 1},
	{Note: conf.E5, Length: 2},
	// Jingle all the way
	{Note: conf.E5, Length: 1},
	{Note: conf.G5, Length: 1},
	{Note: conf.C5, Length: 1.5},
	{Note: conf.D5, Length: 0.5},
	{Note: conf.E5, Length: 4},
	// Oh what fun it is to ride
	{Note: conf.F5, Length: 1},
	{Note: conf.F5, Length: 1},
	{Note: conf.F5, Length: 1.5},
	{Note: conf.F5, Length: 0.5},
	{Note: conf.F5, Length: 1},
	{Note: conf.E5, Length: 1},
	{Note: conf.E5, Length: 1},
	{Note: conf.E5, Length: 1},
	// In a one-horse open sleigh
	{Note: conf.G5, Length: 1},
	{Note: conf.G5, Length: 1},
	{Note: conf.F5, Length: 1},
	{Note: conf.D5, Length: 1},
	{Note: conf.C5, Length: 4},
}

var JingleBellAccompaniment = accompaniment.Accompaniment{
	// Jingle bells, jingle bells
	{Chord: conf.ChordC, Length: 4},
	{Chord: conf.ChordC, Length: 4},
	// Jingle all the way
	{Chord: conf.ChordC, Length: 4},
	{Chord: conf.ChordC, Length: 4},
	// Oh what fun it is to ride
	{Chord: conf.ChordF, Length: 4},
	{Chord: conf.ChordC, Length: 4},
	// In a one-horse open sleigh
	{Chord: conf.ChordD7, Length: 4},
	{Chord: conf.ChordG, Length: 4},
	// Jingle bells, jingle bells
	{Chord: conf.ChordC, Length: 4},
	{Chord: conf.ChordC, Length: 4},
	// Jingle all the way
	{Chord: conf.ChordC, Length: 4},
	{Chord: conf.ChordC, Length: 4},
	// Oh what fun it is to ride
	{Chord: conf.ChordF, Length: 4},
	{Chord: conf.ChordC, Length: 4},
	// In a one-horse open sleigh
	{Chord: conf.ChordG7, Length: 4},
	{Chord: conf.ChordC, Length: 4},
}
