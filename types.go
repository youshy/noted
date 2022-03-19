package notes

// Note symbolises all available notes, calculated as semitones
type Note int

// Fret symbolises the fret on a guitar
type Fret int

// String symbolises a collection of notes on each fret of a string
type String map[Fret]Note

// GuitarStrings symbolises all available guitar strings permutations
// For now, only 6, 7 and 8 strings are available
type GuitarStrings int

// GuitarFrets symbolises all available guitar frets permutations
type GuitarFrets int

// Neck symbolises a collection of notes on each fret on each string
// 0 fret symbolises to which note each string is tuned
type Neck struct {
	FirstString   String
	SecondString  String
	ThirdString   String
	FourthString  String
	FifthString   String
	SixthString   String
	SeventhString String
	EightString   String
}

const (
	C Note = iota
	Csharp
	D
	Dsharp
	E
	F
	Fsharp
	G
	Gsharp
	A
	Asharp
	B
)

const (
	SixString GuitarStrings = iota + 6
	SevenString
	EightString
)

const (
	TwentyOneFrets GuitarFrets = iota + 21
	TwentyTwoFrets
	_
	TwentyFourFrets
	_
	_
	TwentySevenFrets // Because Majesty II has 27, I want to be able to use them all...
)

// Interval depict all available intervals within one octave
type Interval int

const (
	Root Interval = iota
	MinorSecond
	MajorSecond
	MinorThird
	MajorThird
	Fourth
	Tritone
	Fifth
	MinorSixth
	MajorSixth
	MinorSeventh
	MajorSeventh
	Octave
)

// Notes is a collection of all notes generated from root
type Notes []Note
