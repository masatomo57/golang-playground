package main

import (
	"encoding/binary"
	"math"
	"os"

	"github.com/masatomo57/golang-oreore-comparable/jingle-bell/conf"
)

type Accompaniment []struct {
	conf.Chord
	Length float64
}

func main() {
	file := "out.bin"
	f, _ := os.Create(file)
	defer f.Close()

	accompaniment := Accompaniment{
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

	accompaniment.WriteTo(f)
}

func (a Accompaniment) WriteTo(file *os.File) {
	for _, chordWithLength := range a {
		samples := int((chordWithLength.Length * conf.SamplesPerSec) / 4)
		damping := math.Pow(conf.End, 1.0/float64(samples))
		for i := 0; i < samples; i++ {
			notesNum := len(chordWithLength.Chord.Notes)
			vol := 1.0 / float64(notesNum)
			var sample float64
			for j := 0; j < notesNum; j++ {
				sample += vol * math.Sin((2*math.Pi*chordWithLength.Chord.Notes[j].Hertz()*float64(i))/conf.SamplesPerSec)
			}
			sample = sample * math.Pow(damping, float64(i))
			buf := make([]byte, 4)
			binary.LittleEndian.PutUint32(buf, math.Float32bits(float32(sample)))
			file.Write(buf)
		}
	}
}
