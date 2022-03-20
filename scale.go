package notes

import (
	"errors"
)

const (
	ErrUnsortedIntervals = "Please sort the intervals for the scale on the way in"
)

// NewScale allows for an easy build of the new scale.
// By default, it creates a scale from a middle C (C4)
func NewScale(
	name string,
	intervals ...Interval,
) (Scale, error) {
	defaultRoot := Note{
		Pitch:       C,
		PitchOctave: FourthOctave,
	}

	return buildNewScale(name, defaultRoot, intervals)
}

// NewScaleWithRoot creates the scale within a given key.
func NewScaleFromRoot(
	name string,
	rootNote Note,
	intervals ...Interval,
) (Scale, error) {

	return buildNewScale(name, rootNote, intervals)
}

func buildNewScale(
	name string,
	rootNote Note,
	intervals []Interval,
) (Scale, error) {
	var s Scale
	s.RootNote = rootNote

	ok := validateIntervals(intervals)
	if !ok {
		return Scale{}, errors.New(ErrUnsortedIntervals)
	}

	s.Intervals = intervals

	s.Notes = generateNotesWithinSingleOctave(rootNote, intervals)

	return s, nil
}

func validateIntervals(intervals []Interval) bool {
	for i := 0; i < len(intervals); i++ {
		if i+1 == len(intervals) {
			return true
		}

		if intervals[i] > intervals[i+1] {
			return false
		}
	}

	return true
}

func generateNotesWithinSingleOctave(rootNote Note, intervals []Interval) Notes {
	var (
		notes    = make(Notes, 0)
		tempNote Note
	)

	for i := 0; i < len(intervals); i++ {
		tempNote = rootNote
		tempNote.raiseNoteByInterval(intervals[i])

		notes = append(notes, tempNote)
	}

	return notes
}

func (n *Note) raiseNoteByInterval(interval Interval) {
	temp := n.Pitch + Pitch(interval)

	if temp > 11 {
		n.Pitch = (n.Pitch + Pitch(interval)) - 12
		n.PitchOctave++
	} else {
		n.Pitch = n.Pitch + Pitch(interval)
	}
}
