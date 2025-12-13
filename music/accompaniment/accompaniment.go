package accompaniment

import (
	"encoding/binary"
	"math"
	"os"

	"github.com/masatomo57/golang-oreore-comparable/music/conf"
)

type Accompaniment []struct {
	conf.Chord
	Length float64
}

// GenerateSamples generates audio samples as []float32 without writing to file
func (a Accompaniment) GenerateSamples() []float32 {
	var samples []float32
	for _, chordWithLength := range a {
		sampleCount := int((chordWithLength.Length * conf.SamplesPerSec) / 4)
		damping := math.Pow(conf.End, 1.0/float64(sampleCount))
		for i := 0; i < sampleCount; i++ {
			notesNum := len(chordWithLength.Chord.Notes)
			vol := 1.0 / float64(notesNum)
			var sample float64
			for j := 0; j < notesNum; j++ {
				sample += vol * math.Sin((2*math.Pi*chordWithLength.Chord.Notes[j].Hertz()*float64(i))/ float64(conf.SamplesPerSec))
			}
			sample = sample * math.Pow(damping, float64(i))
			samples = append(samples, float32(sample))
		}
	}
	return samples
}

// WriteToFile writes the accompaniment samples to a file
func (a Accompaniment) WriteToFile(file *os.File) {
	for _, chordWithLength := range a {
		samples := int((chordWithLength.Length * conf.SamplesPerSec) / 4)
		damping := math.Pow(conf.End, 1.0/float64(samples))
		for i := 0; i < samples; i++ {
			notesNum := len(chordWithLength.Chord.Notes)
			vol := 1.0 / float64(notesNum)
			var sample float64
			for j := 0; j < notesNum; j++ {
				sample += vol * math.Sin((2*math.Pi*chordWithLength.Chord.Notes[j].Hertz()*float64(i))/ float64(conf.SamplesPerSec))
			}
			sample = sample * math.Pow(damping, float64(i))
			buf := make([]byte, 4)
			binary.LittleEndian.PutUint32(buf, math.Float32bits(float32(sample)))
			file.Write(buf)
		}
	}
}
