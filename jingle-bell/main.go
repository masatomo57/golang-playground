package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	melodyFile, _ := os.Open("./melody/cmd/out.bin")
	defer melodyFile.Close()
	melodyF := []float32{}
	for {
		var f float32
		err := binary.Read(melodyFile, binary.LittleEndian, &f)
		if err == io.EOF {
			break
		}
		melodyF = append(melodyF, f)
	}
	fmt.Println(len(melodyF))

	chordFile, _ := os.Open("./accompaniment/cmd/out.bin")
	defer chordFile.Close()
	chordF := []float32{}
	for {
		var f float32
		err := binary.Read(chordFile, binary.LittleEndian, &f)
		if err == io.EOF {
			break
		}
		chordF = append(chordF, f)
	}
	fmt.Println(len(chordF))

	file := "out.bin"
	f, _ := os.Create(file)
	samples := min(len(melodyF), len(chordF))
	for i := 0; i < samples; i++ {
		sample := 0.5*melodyF[i] + 0.5*chordF[i]
		buf := make([]byte, 4)

		// バイト順序=LittleEndian
		binary.LittleEndian.PutUint32(buf, math.Float32bits(float32(sample)))
		f.Write(buf)
	}
}
