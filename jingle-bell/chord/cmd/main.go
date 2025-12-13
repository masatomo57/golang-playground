package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

const (
	samplesPerSec = 44000
	tau           = 2 * math.Pi

	end = 1.0e-1

	B2 = 123.471

	C3  = 130.813
	D3  = 146.832
	E3  = 164.814
	F3  = 174.614
	Fs3 = 184.997
	G3  = 195.998
	A3  = 220.000
	As3 = 233.082
	B3  = 246.942

	C4  = 261.626
	Cs4 = 277.183
	D4  = 293.665
	E4  = 329.628
	F4  = 349.228
	Fs4 = 369.994
	G4  = 391.995
	A4  = 440.000
	B4  = 493.883

	C5  = 523.251
	Cs5 = 554.365
	D5  = 587.330
	E5  = 659.255
	F5  = 698.456
	Fs5 = 739.989
	G5  = 783.991
	A5  = 880.000
	B5  = 987.767
)

func main() {
	file := "out.bin"
	f, _ := os.Create(file)
	defer f.Close()

	// Jingle bells, jingle bells
	generate(C3, E3, G3, 0, 4, f) // C
	generate(C3, E3, G3, 0, 4, f) // C

	// Jingle all the way
	generate(C3, E3, G3, 0, 4, f) // C
	generate(C3, E3, G3, 0, 4, f) // C

	// Oh what fun it is to ride
	generate(F3, A3, C4, 0, 4, f) // F
	generate(C3, E3, G3, 0, 4, f) // C

	// In a one-horse open sleigh
	generate(C3, D3, Fs3, A3, 4, f) // D7
	generate(B2, D3, G3, 0, 4, f)   // G

	// Jingle bells, jingle bells
	generate(C3, E3, G3, 0, 4, f) // C
	generate(C3, E3, G3, 0, 4, f) // C

	// Jingle all the way
	generate(C3, E3, G3, 0, 4, f) // C
	generate(C3, E3, G3, 0, 4, f) // C

	// Oh what fun it is to ride
	generate(F3, A3, C4, 0, 4, f) // F
	generate(C3, E3, G3, 0, 4, f) // C

	// In a one-horse open sleigh
	generate(B2, F3, G3, 0, 4, f) // G7
	generate(C3, E3, G3, 0, 4, f) // C

	// silence
	generate(0, 0, 0, 0, 4, f)
}

func generate(frequency1, frequency2, frequency3, frequency4 float64, soundLength float64, file *os.File) {
	samples := int((soundLength * samplesPerSec) / 4)
	fmt.Println(samples)
	damping := math.Pow(end, 1.0/float64(samples))
	for i := 0; i < samples; i++ {
		sample := 0.25*math.Sin((tau*frequency1*float64(i))/samplesPerSec) +
			0.25*math.Sin((tau*frequency2*float64(i))/samplesPerSec) +
			0.25*math.Sin((tau*frequency3*float64(i))/samplesPerSec) +
			0.25*math.Sin((tau*frequency4*float64(i))/samplesPerSec)

		sample = sample * math.Pow(damping, float64(i))
		buf := make([]byte, 4)

		// バイト順序=LittleEndian
		binary.LittleEndian.PutUint32(buf, math.Float32bits(float32(sample)))
		file.Write(buf)
	}
}
