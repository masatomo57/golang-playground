package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

const (
	samplesPerSec = 44100
	tau           = 2 * math.Pi

	end = 1.0e-1

	C3  = 130.813
	D3  = 146.832
	E3  = 164.814
	F3  = 174.614
	Fs3 = 184.997
	G3  = 195.998
	A3  = 220.000
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
	Ds5 = 622.254
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
	generateMelody(E5, 1, f)
	generateMelody(E5, 1, f)
	generateMelody(E5, 2, f)

	generateMelody(E5, 1, f)
	generateMelody(E5, 1, f)
	generateMelody(E5, 2, f)

	// Jingle all the way
	generateMelody(E5, 1, f)
	generateMelody(G5, 1, f)
	generateMelody(C5, 1.5, f)
	generateMelody(D5, 0.5, f)
	generateMelody(E5, 4, f)

	// Oh what fun it is to ride
	generateMelody(F5, 1, f)
	generateMelody(F5, 1, f)
	generateMelody(F5, 1.5, f)
	generateMelody(F5, 0.5, f)

	generateMelody(F5, 1, f)
	generateMelody(E5, 1, f)
	generateMelody(E5, 1, f)
	generateMelody(E5, 1, f)

	// In a one-horse open sleigh
	generateMelody(E5, 1, f)
	generateMelody(D5, 1, f)
	generateMelody(D5, 1, f)
	generateMelody(E5, 1, f)
	generateMelody(D5, 2, f)
	generateMelody(G5, 2, f)

	// Jingle bells, jingle bells
	generateMelody(E5, 1, f)
	generateMelody(E5, 1, f)
	generateMelody(E5, 2, f)

	generateMelody(E5, 1, f)
	generateMelody(E5, 1, f)
	generateMelody(E5, 2, f)

	// Jingle all the way
	generateMelody(E5, 1, f)
	generateMelody(G5, 1, f)
	generateMelody(C5, 1.5, f)
	generateMelody(D5, 0.5, f)
	generateMelody(E5, 4, f)

	// Oh what fun it is to ride
	generateMelody(F5, 1, f)
	generateMelody(F5, 1, f)
	generateMelody(F5, 1.5, f)
	generateMelody(F5, 0.5, f)

	generateMelody(F5, 1, f)
	generateMelody(E5, 1, f)
	generateMelody(E5, 1, f)
	generateMelody(E5, 1, f)

	// In a one-horse open sleigh
	generateMelody(G5, 1, f)
	generateMelody(G5, 1, f)
	generateMelody(F5, 1, f)
	generateMelody(D5, 1, f)
	generateMelody(C5, 4, f)

	// silence
	generateMelody(0, 4, f)
}

func generateMelody(frequency float64, soundLength float64, file *os.File) {
	samples := int((soundLength * samplesPerSec) / 4)
	fmt.Println(samples)
	damping := math.Pow(end, 1.0/float64(samples))
	for i := 0; i < samples; i++ {
		sample := math.Sin((tau * frequency * float64(i)) / float64(samplesPerSec))
		sample = sample * math.Pow(damping, float64(i))
		buf := make([]byte, 4)

		// バイト順序=LittleEndian
		binary.LittleEndian.PutUint32(buf, math.Float32bits(float32(sample)))
		file.Write(buf)
	}
}
