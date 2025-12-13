package score

import (
	"github.com/masatomo57/golang-playground/music/accompaniment"
	"github.com/masatomo57/golang-playground/music/melody"
)

type Score struct {
	Melody melody.Melody
	Accompaniment accompaniment.Accompaniment
}

var Scores = map[string]Score{
	"jingle_bell": {
		Melody: JingleBellMelody,
		Accompaniment: JingleBellAccompaniment,
	},
}
