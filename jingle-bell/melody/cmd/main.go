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

type Melody []struct {
	conf.Note
	Length float64
}

func main() {
	file := "out.bin"
	f, _ := os.Create(file)
	defer f.Close()

	melody := Melody{
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

	melody.WriteTo(f)
}

func (m Melody) WriteTo(file *os.File) {
	for _, noteWithLength := range m {
		samples := int((noteWithLength.Length * samplesPerSec) / 4)
		damping := math.Pow(end, 1.0/float64(samples))
		for i := 0; i < samples; i++ {
			sample := math.Sin((tau * noteWithLength.Note.Hertz() * float64(i)) / float64(samplesPerSec))
			sample = sample * math.Pow(damping, float64(i))
			buf := make([]byte, 4)
			binary.LittleEndian.PutUint32(buf, math.Float32bits(float32(sample)))
			file.Write(buf)
		}
	}
}
