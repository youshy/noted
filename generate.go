package notes

import "fmt"

// GenerateNeck generates the full map of all notes
// available on each string of the neck.
// Given the parameters, it builds the map.
func GenerateNeck(
	strings GuitarStrings,
	frets GuitarFrets,
	tuning ...Note,
) (Neck, error) {
	var neck Neck

	switch strings {
	case SixString:
		neck = Neck{
			FirstString:  generateString(tuning[0], frets),
			SecondString: generateString(tuning[1], frets),
			ThirdString:  generateString(tuning[2], frets),
			FourthString: generateString(tuning[3], frets),
			FifthString:  generateString(tuning[4], frets),
			SixthString:  generateString(tuning[5], frets),
		}
	case SevenString:
		neck = Neck{
			FirstString:   generateString(tuning[0], frets),
			SecondString:  generateString(tuning[1], frets),
			ThirdString:   generateString(tuning[2], frets),
			FourthString:  generateString(tuning[3], frets),
			FifthString:   generateString(tuning[4], frets),
			SixthString:   generateString(tuning[5], frets),
			SeventhString: generateString(tuning[6], frets),
		}
	case EightString:
		neck = Neck{
			FirstString:   generateString(tuning[0], frets),
			SecondString:  generateString(tuning[1], frets),
			ThirdString:   generateString(tuning[2], frets),
			FourthString:  generateString(tuning[3], frets),
			FifthString:   generateString(tuning[4], frets),
			SixthString:   generateString(tuning[5], frets),
			SeventhString: generateString(tuning[6], frets),
			EightString:   generateString(tuning[7], frets),
		}
	default:
		return Neck{}, fmt.Errorf("Not available to generate neck for %v strings", strings)
	}

	return neck, nil
}

func generateString(
	rootNote Note,
	frets GuitarFrets,
) String {
	var (
		singleString = make(String, 1)
		currentNote  Note
		currentFret  Fret
	)

	allNotesFromRoot := generateAllNotesFromRootWithinSingleOctave(rootNote)

	for i := 0; i <= int(frets); i++ {
		singleString[currentFret] = allNotesFromRoot[currentNote]
		currentFret++
		currentNote++

		if currentNote == 12 {
			currentNote = 0
		}
	}

	return singleString
}

func generateAllNotesFromRootWithinSingleOctave(rootNote Note) Notes {
	var (
		currentNote Note
		notes       = make(Notes, 0)
	)

	currentNote = rootNote

	for i := Root; i < Octave; i++ {
		notes = append(notes, currentNote)
		currentNote++

		if currentNote == 12 {
			currentNote = 0
		}
	}

	return notes
}
