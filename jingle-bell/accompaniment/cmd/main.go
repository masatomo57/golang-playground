package main

import (
	"os"

	"github.com/masatomo57/golang-playground/jingle-bell/accompaniment"
	"github.com/masatomo57/golang-playground/jingle-bell/conf"
)

func main() {
	file := "out.bin"
	f, _ := os.Create(file)
	defer f.Close()

	a := accompaniment.Accompaniment{
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

	a.WriteToFile(f)
}
