package main

import (
	"encoding/binary"
	"fmt"
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
	generateMelody(conf.E5, 1, f)
	generateMelody(conf.E5, 1, f)
	generateMelody(conf.E5, 2, f)

	generateMelody(conf.E5, 1, f)
	generateMelody(conf.E5, 1, f)
	generateMelody(conf.E5, 2, f)

	// Jingle all the way
	generateMelody(conf.E5, 1, f)
	generateMelody(conf.G5, 1, f)
	generateMelody(conf.C5, 1.5, f)
	generateMelody(conf.D5, 0.5, f)
	generateMelody(conf.E5, 4, f)

	// Oh what fun it is to ride
	generateMelody(conf.F5, 1, f)
	generateMelody(conf.F5, 1, f)
	generateMelody(conf.F5, 1.5, f)
	generateMelody(conf.F5, 0.5, f)

	generateMelody(conf.F5, 1, f)
	generateMelody(conf.E5, 1, f)
	generateMelody(conf.E5, 1, f)
	generateMelody(conf.E5, 1, f)

	// In a one-horse open sleigh
	generateMelody(conf.E5, 1, f)
	generateMelody(conf.D5, 1, f)
	generateMelody(conf.D5, 1, f)
	generateMelody(conf.E5, 1, f)
	generateMelody(conf.D5, 2, f)
	generateMelody(conf.G5, 2, f)

	// Jingle bells, jingle bells
	generateMelody(conf.E5, 1, f)
	generateMelody(conf.E5, 1, f)
	generateMelody(conf.E5, 2, f)

	generateMelody(conf.E5, 1, f)
	generateMelody(conf.E5, 1, f)
	generateMelody(conf.E5, 2, f)

	// Jingle all the way
	generateMelody(conf.E5, 1, f)
	generateMelody(conf.G5, 1, f)
	generateMelody(conf.C5, 1.5, f)
	generateMelody(conf.D5, 0.5, f)
	generateMelody(conf.E5, 4, f)

	// Oh what fun it is to ride
	generateMelody(conf.F5, 1, f)
	generateMelody(conf.F5, 1, f)
	generateMelody(conf.F5, 1.5, f)
	generateMelody(conf.F5, 0.5, f)

	generateMelody(conf.F5, 1, f)
	generateMelody(conf.E5, 1, f)
	generateMelody(conf.E5, 1, f)
	generateMelody(conf.E5, 1, f)

	// In a one-horse open sleigh
	generateMelody(conf.G5, 1, f)
	generateMelody(conf.G5, 1, f)
	generateMelody(conf.F5, 1, f)
	generateMelody(conf.D5, 1, f)
	generateMelody(conf.C5, 4, f)
}

func generateMelody(note conf.Note, soundLength float64, file *os.File) {
	samples := int((soundLength * samplesPerSec) / 4)
	fmt.Println(samples)
	damping := math.Pow(end, 1.0/float64(samples))
	for i := 0; i < samples; i++ {
		sample := math.Sin((tau * note.Hertz() * float64(i)) / float64(samplesPerSec))
		sample = sample * math.Pow(damping, float64(i))
		buf := make([]byte, 4)

		// バイト順序=LittleEndian
		binary.LittleEndian.PutUint32(buf, math.Float32bits(float32(sample)))
		file.Write(buf)
	}
}
