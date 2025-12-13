package main

import (
	"flag"
	"log"
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
	a := score.Accompaniment

	file := "out.bin"
	f, _ := os.Create(file)
	defer f.Close()

	a.WriteToFile(f)
}
