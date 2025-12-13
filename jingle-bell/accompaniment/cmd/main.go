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
	generate(conf.C3, conf.E3, conf.G3, conf.None, 4, f) // C
	generate(conf.C3, conf.E3, conf.G3, conf.None, 4, f) // C

	// Jingle all the way
	generate(conf.C3, conf.E3, conf.G3, conf.None, 4, f) // C
	generate(conf.C3, conf.E3, conf.G3, conf.None, 4, f) // C

	// Oh what fun it is to ride
	generate(conf.F3, conf.A3, conf.C4, conf.None, 4, f) // F
	generate(conf.C3, conf.E3, conf.G3, conf.None, 4, f) // C

	// In a one-horse open sleigh
	generate(conf.C3, conf.D3, conf.Fs3, conf.A3, 4, f) // D7
	generate(conf.B2, conf.D3, conf.G3, conf.None, 4, f)        // G

	// Jingle bells, jingle bells
	generate(conf.C3, conf.E3, conf.G3, conf.None, 4, f) // C
	generate(conf.C3, conf.E3, conf.G3, conf.None, 4, f) // C

	// Jingle all the way
	generate(conf.C3, conf.E3, conf.G3, conf.None, 4, f) // C
	generate(conf.C3, conf.E3, conf.G3, conf.None, 4, f) // C

	// Oh what fun it is to ride
	generate(conf.F3, conf.A3, conf.C4, conf.None, 4, f) // F
	generate(conf.C3, conf.E3, conf.G3, conf.None, 4, f) // C

	// In a one-horse open sleigh
	generate(conf.B2, conf.F3, conf.G3, conf.None, 4, f) // G7
	generate(conf.C3, conf.E3, conf.G3, conf.None, 4, f) // C
}

func generate(note1, note2, note3, note4 conf.Note, soundLength float64, file *os.File) {
	samples := int((soundLength * samplesPerSec) / 4)
	fmt.Println(samples)
	damping := math.Pow(end, 1.0/float64(samples))
	for i := 0; i < samples; i++ {
		sample := 0.25*math.Sin((tau*note1.Hertz()*float64(i))/samplesPerSec) +
			0.25*math.Sin((tau*note2.Hertz()*float64(i))/samplesPerSec) +
			0.25*math.Sin((tau*note3.Hertz()*float64(i))/samplesPerSec) +
			0.25*math.Sin((tau*note4.Hertz()*float64(i))/samplesPerSec)

		sample = sample * math.Pow(damping, float64(i))
		buf := make([]byte, 4)

		// バイト順序=LittleEndian
		binary.LittleEndian.PutUint32(buf, math.Float32bits(float32(sample)))
		file.Write(buf)
	}
}
