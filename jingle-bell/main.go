package main

import (
	"encoding/binary"
	"math"
	"os"

	"github.com/masatomo57/golang-oreore-comparable/jingle-bell/accompaniment"
	"github.com/masatomo57/golang-oreore-comparable/jingle-bell/conf"
	"github.com/masatomo57/golang-oreore-comparable/jingle-bell/melody"
)

func main() {
	m := melody.Melody{
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

	melodySamples := m.GenerateSamples()

	accompanimentSamples := a.GenerateSamples()

	file := "out.bin"
	f, _ := os.Create(file)
	defer f.Close()
	samples := min(len(melodySamples), len(accompanimentSamples))
	for i := 0; i < samples; i++ {
		sample := 0.5*melodySamples[i] + 0.5*accompanimentSamples[i]
		buf := make([]byte, 4)

		binary.LittleEndian.PutUint32(buf, math.Float32bits(float32(sample)))
		f.Write(buf)
	}
}
