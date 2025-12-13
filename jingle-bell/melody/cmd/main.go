package main

import (
	"os"

	"github.com/masatomo57/golang-playground/jingle-bell/conf"
	"github.com/masatomo57/golang-playground/jingle-bell/melody"
)

func main() {
	file := "out.bin"
	f, _ := os.Create(file)
	defer f.Close()

	m := melody.Melody{
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

	m.WriteTo(f)
}
