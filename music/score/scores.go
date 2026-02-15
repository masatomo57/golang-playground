package score

import (
	"encoding/binary"
	"math"
	"os"

	"github.com/masatomo57/golang-playground/music/accompaniment"
	"github.com/masatomo57/golang-playground/music/melody"
)

type Score struct {
	Melody        melody.Melody
	Accompaniment accompaniment.Accompaniment
}

var Scores = map[string]Score{
	"jingle_bell": {
		Melody:        JingleBellMelody,
		Accompaniment: JingleBellAccompaniment,
	},
}

func (s Score) WriteToFile(file *os.File) {
	melodySamples := s.Melody.GenerateSamples()
	accompanimentSamples := s.Accompaniment.GenerateSamples()
	samples := min(len(melodySamples), len(accompanimentSamples))
	for i := 0; i < samples; i++ {
		sample := 0.5*melodySamples[i] + 0.5*accompanimentSamples[i]
		buf := make([]byte, 4)
		binary.LittleEndian.PutUint32(buf, math.Float32bits(float32(sample)))
		file.Write(buf)
	}
}
