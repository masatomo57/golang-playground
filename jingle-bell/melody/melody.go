package melody

import (
	"encoding/binary"
	"math"
	"os"

	"github.com/masatomo57/golang-playground/jingle-bell/conf"
)

type Melody []struct {
	conf.Note
	Length float64
}

// GenerateSamples generates audio samples as []float32 without writing to file
func (m Melody) GenerateSamples() []float32 {
	var samples []float32
	for _, noteWithLength := range m {
		sampleCount := int((noteWithLength.Length * conf.SamplesPerSec) / 4)
		damping := math.Pow(conf.End, 1.0/float64(sampleCount))
		for i := 0; i < sampleCount; i++ {
			sample := math.Sin((2 * math.Pi * noteWithLength.Note.Hertz() * float64(i)) / float64(conf.SamplesPerSec))
			sample = sample * math.Pow(damping, float64(i))
			samples = append(samples, float32(sample))
		}
	}
	return samples
}

// WriteTo writes the melody samples to a file
func (m Melody) WriteTo(file *os.File) {
	for _, noteWithLength := range m {
		samples := int((noteWithLength.Length * conf.SamplesPerSec) / 4)
		damping := math.Pow(conf.End, 1.0/float64(samples))
		for i := 0; i < samples; i++ {
			sample := math.Sin((2 * math.Pi * noteWithLength.Note.Hertz() * float64(i)) / float64(conf.SamplesPerSec))
			sample = sample * math.Pow(damping, float64(i))
			buf := make([]byte, 4)
			binary.LittleEndian.PutUint32(buf, math.Float32bits(float32(sample)))
			file.Write(buf)
		}
	}
}
