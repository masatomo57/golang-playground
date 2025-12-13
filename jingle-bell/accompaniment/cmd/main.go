package main

import (
	"encoding/binary"
	"math"
	"os"

	"github.com/masatomo57/golang-oreore-comparable/jingle-bell/conf"
)

const (
	samplesPerSec = 44000
	tau           = 2 * math.Pi

	end = 1.0e-1
)

func main() {
	file := "out.bin"
	f, _ := os.Create(file)
	defer f.Close()

	// Jingle bells, jingle bells
	generate(conf.ChordC, 4, f) // C
	generate(conf.ChordC, 4, f) // C

	// Jingle all the way
	generate(conf.ChordC, 4, f) // C
	generate(conf.ChordC, 4, f) // C

	// Oh what fun it is to ride
	generate(conf.ChordF, 4, f) // F
	generate(conf.ChordC, 4, f) // C

	// In a one-horse open sleigh
	generate(conf.ChordD7, 4, f) // D7
	generate(conf.ChordG, 4, f)  // G

	// Jingle bells, jingle bells
	generate(conf.ChordC, 4, f) // C
	generate(conf.ChordC, 4, f) // C

	// Jingle all the way
	generate(conf.ChordC, 4, f) // C
	generate(conf.ChordC, 4, f) // C

	// Oh what fun it is to ride
	generate(conf.ChordF, 4, f) // F
	generate(conf.ChordC, 4, f) // C

	// In a one-horse open sleigh
	generate(conf.ChordG7, 4, f) // G7
	generate(conf.ChordC, 4, f)  // C
}

func generate(chord conf.Chord, soundLength float64, file *os.File) {
	samples := int((soundLength * samplesPerSec) / 4)
	damping := math.Pow(end, 1.0/float64(samples))
	for i := 0; i < samples; i++ {
		notesNum := len(chord.Notes)
		vol := 1.0 / float64(notesNum)
		var sample float64
		for j := 0; j < notesNum; j++ {
			sample += vol * math.Sin((tau*chord.Notes[j].Hertz()*float64(i))/samplesPerSec)
		}
		sample = sample * math.Pow(damping, float64(i))
		buf := make([]byte, 4)

		// バイト順序=LittleEndian
		binary.LittleEndian.PutUint32(buf, math.Float32bits(float32(sample)))
		file.Write(buf)
	}
}
