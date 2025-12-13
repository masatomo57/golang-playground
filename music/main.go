package main

import (
	"encoding/binary"
	"flag"
	"log"
	"math"
	"os"

	"github.com/masatomo57/golang-playground/music/score"
)

func main() {
	title := flag.String("title", "jingle_bell", "")
	flag.Parse()

	score, ok := score.Scores[*title]
	if !ok {
		log.Fatalf("score not found: %s", *title)
	}
	m := score.Melody
	a := score.Accompaniment

	melodySamples := m.GenerateSamples()
	accompanimentSamples := a.GenerateSamples()

	file := "out.bin"
	f, _ := os.Create(file)
	defer f.Close()
	samples := min(len(melodySamples), len(accompanimentSamples))
	for i := 0; i < samples; i++ {
		sample := 0.5*melodySamples[i] + 0.5*accompanimentSamples[i]
		buf := make([]byte, 4)

		binary.LittleEndian.PutUint32(buf, math.Float32bits(float32(sample)))
		f.Write(buf)
	}
}
