package notes

// Pitch symbolises all available pitches in twelve-tone system
type Pitch int

// PitchOctave symbolises all available octaves for a single note
type PitchOctave int

// Note symbolises all available notes within all available octaves.
// IsInScale checks if the note is available in a scale
// generated for the neck
type Note struct {
	Pitch       Pitch
	PitchOctave PitchOctave
	IsInScale   bool
}

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

// Interval depict all available intervals within one octave
type Interval int

// Notes is a collection of all notes generated from root
type Notes []Note

// Scale symbolises any custom scale that will allow for
// propagating all available notes on the neck
type Scale struct {
	// Name is the name of the scale
	Name string
	// Intervals describes how the scale is built within one octave
	Intervals []Interval
	// RootNote sets the scale starting point
	RootNote Note
	// Notes describes all available notes
	// within one octave from the Root note
	Notes Notes
}

const (
	C Pitch = iota
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

const (
	MinusOneOctave PitchOctave = iota - 1
	ZeroOctave
	FirstOctave
	SecondOctave
	ThirdOctave
	FourthOctave
	FifthOctave
	SixthOctave
	SeventhOctave
	EightOctave
	NinthOctave
)
