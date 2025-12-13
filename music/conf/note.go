package conf

type Note float64

const (
	B2 Note = 123.471

	C3  Note = 130.813
	D3  Note = 146.832
	E3  Note = 164.814
	F3  Note = 174.614
	Fs3 Note = 184.997
	G3  Note = 195.998
	A3  Note = 220.000
	As3 Note = 233.082
	B3  Note = 246.942

	C4  Note = 261.626
	Cs4 Note = 277.183
	D4  Note = 293.665
	E4  Note = 329.628
	F4  Note = 349.228
	Fs4 Note = 369.994
	G4  Note = 391.995
	A4  Note = 440.000
	B4  Note = 493.883

	C5  Note = 523.251
	Cs5 Note = 554.365
	D5  Note = 587.330
	Ds5 Note = 622.254
	E5  Note = 659.255
	F5  Note = 698.456
	Fs5 Note = 739.989
	G5  Note = 783.991
	A5  Note = 880.000
	B5  Note = 987.767

	None Note = 0
)

func (n Note) Hertz() float64 {
	return float64(n)
}