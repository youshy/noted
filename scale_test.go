package notes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewScale(t *testing.T) {
	t.Run("Returns a valid C minor (aeolian) scale", func(t *testing.T) {
		expected := Scale{
			Intervals: []Interval{
				Root,
				MajorSecond,
				MinorThird,
				Fourth,
				Fifth,
				MinorSixth,
				MinorSeventh,
			},
			RootNote: Note{
				Pitch:       C,
				PitchOctave: FourthOctave,
			},
			Notes: []Note{
				Note{
					Pitch:       C,
					PitchOctave: FourthOctave,
				},
				Note{
					Pitch:       D,
					PitchOctave: FourthOctave,
				},
				Note{
					Pitch:       Dsharp,
					PitchOctave: FourthOctave,
				},
				Note{
					Pitch:       F,
					PitchOctave: FourthOctave,
				},
				Note{
					Pitch:       G,
					PitchOctave: FourthOctave,
				},
				Note{
					Pitch:       Gsharp,
					PitchOctave: FourthOctave,
				},
				Note{
					Pitch:       Asharp,
					PitchOctave: FourthOctave,
				},
			},
		}

		actual, err := NewScale("aeolian",
			Root,
			MajorSecond,
			MinorThird,
			Fourth,
			Fifth,
			MinorSixth,
			MinorSeventh,
		)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestNewScaleFromRoot(t *testing.T) {
	t.Run("Returns a valid B lydian dominant scale, starting from B2", func(t *testing.T) {
		expected := Scale{
			Intervals: []Interval{
				Root,
				MajorSecond,
				MajorThird,
				Tritone,
				Fifth,
				MajorSixth,
				MinorSeventh,
			},
			RootNote: Note{
				Pitch:       B,
				PitchOctave: SecondOctave,
			},
			Notes: []Note{
				Note{
					Pitch:       B,
					PitchOctave: SecondOctave,
				},
				Note{
					Pitch:       Csharp,
					PitchOctave: ThirdOctave,
				},
				Note{
					Pitch:       Dsharp,
					PitchOctave: ThirdOctave,
				},
				Note{
					Pitch:       F,
					PitchOctave: ThirdOctave,
				},
				Note{
					Pitch:       Fsharp,
					PitchOctave: ThirdOctave,
				},
				Note{
					Pitch:       Gsharp,
					PitchOctave: ThirdOctave,
				},
				Note{
					Pitch:       A,
					PitchOctave: ThirdOctave,
				},
			},
		}

		inputRootNote := Note{
			Pitch:       B,
			PitchOctave: SecondOctave,
		}

		actual, err := NewScaleFromRoot("lydian dominant",
			inputRootNote,
			Root,
			MajorSecond,
			MajorThird,
			Tritone,
			Fifth,
			MajorSixth,
			MinorSeventh,
		)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Returns a valid D minor pentatonic scale, starting from D4", func(t *testing.T) {
		expected := Scale{
			Intervals: []Interval{
				Root,
				MinorThird,
				Fourth,
				Fifth,
				MinorSeventh,
			},
			RootNote: Note{
				Pitch:       D,
				PitchOctave: FourthOctave,
			},
			Notes: []Note{
				Note{
					Pitch:       D,
					PitchOctave: FourthOctave,
				},
				Note{
					Pitch:       F,
					PitchOctave: FourthOctave,
				},
				Note{
					Pitch:       G,
					PitchOctave: FourthOctave,
				},
				Note{
					Pitch:       A,
					PitchOctave: FourthOctave,
				},
				Note{
					Pitch:       C,
					PitchOctave: FifthOctave,
				},
			},
		}

		inputRootNote := Note{
			Pitch:       D,
			PitchOctave: FourthOctave,
		}

		actual, err := NewScaleFromRoot("minor pentatonic",
			inputRootNote,
			Root,
			MinorThird,
			Fourth,
			Fifth,
			MinorSeventh,
		)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestBuildNewScale(t *testing.T) {
	t.Run("Returns a valid C major (ionian) scale, starting from C4", func(t *testing.T) {
		expected := Scale{
			Intervals: []Interval{
				Root,
				MajorSecond,
				MajorThird,
				Fourth,
				Fifth,
				MajorSixth,
				MajorSeventh,
			},
			RootNote: Note{
				Pitch:       C,
				PitchOctave: FourthOctave,
			},
			Notes: []Note{
				Note{
					Pitch:       C,
					PitchOctave: FourthOctave,
				},
				Note{
					Pitch:       D,
					PitchOctave: FourthOctave,
				},
				Note{
					Pitch:       E,
					PitchOctave: FourthOctave,
				},
				Note{
					Pitch:       F,
					PitchOctave: FourthOctave,
				},
				Note{
					Pitch:       G,
					PitchOctave: FourthOctave,
				},
				Note{
					Pitch:       A,
					PitchOctave: FourthOctave,
				},
				Note{
					Pitch:       B,
					PitchOctave: FourthOctave,
				},
			},
		}

		inputRootNote := Note{
			Pitch:       C,
			PitchOctave: FourthOctave,
		}
		inputIntervals := []Interval{
			Root,
			MajorSecond,
			MajorThird,
			Fourth,
			Fifth,
			MajorSixth,
			MajorSeventh,
		}

		actual, err := buildNewScale("ionian", inputRootNote, inputIntervals)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Returns a valid whole-tone scale, starting from G3", func(t *testing.T) {
		expected := Scale{
			Intervals: []Interval{
				Root,
				MajorSecond,
				MajorThird,
				Tritone,
				MinorSixth,
				MinorSeventh,
			},
			RootNote: Note{
				Pitch:       G,
				PitchOctave: ThirdOctave,
			},
			Notes: []Note{
				Note{
					Pitch:       G,
					PitchOctave: ThirdOctave,
				},
				Note{
					Pitch:       A,
					PitchOctave: ThirdOctave,
				},
				Note{
					Pitch:       B,
					PitchOctave: ThirdOctave,
				},
				Note{
					Pitch:       Csharp,
					PitchOctave: FourthOctave,
				},
				Note{
					Pitch:       Dsharp,
					PitchOctave: FourthOctave,
				},
				Note{
					Pitch:       F,
					PitchOctave: FourthOctave,
				},
			},
		}

		inputRootNote := Note{
			Pitch:       G,
			PitchOctave: ThirdOctave,
		}
		inputIntervals := []Interval{
			Root,
			MajorSecond,
			MajorThird,
			Tritone,
			MinorSixth,
			MinorSeventh,
		}

		actual, err := buildNewScale("whole-tone", inputRootNote, inputIntervals)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)

	})

	t.Run("Returns an error if the input scale intervals' are unsorted", func(t *testing.T) {
		var expected Scale

		inputRootNote := Note{}
		inputIntervals := []Interval{
			MajorThird,
			Root,
			MajorSeventh,
			Tritone,
		}

		actual, err := buildNewScale("", inputRootNote, inputIntervals)

		assert.Error(t, err)
		assert.EqualError(t, err, ErrUnsortedIntervals)
		assert.Equal(t, expected, actual)
	})

	t.Run("Reassigns the intervals without modifying their order", func(t *testing.T) {
		expected := Scale{
			Intervals: []Interval{
				Root,
				MinorThird,
				Tritone,
				MajorSixth,
			},
		}

		inputRootNote := Note{}
		inputIntervals := []Interval{
			Root,
			MinorThird,
			Tritone,
			MajorSixth,
		}

		actual, err := buildNewScale("", inputRootNote, inputIntervals)

		assert.NoError(t, err)
		assert.Equal(t, expected.Intervals, actual.Intervals)
	})
}

func TestValidateIntervals(t *testing.T) {
	t.Run("Returns true for sorted intervals of a major (ionian) scale", func(t *testing.T) {
		expected := true
		input := []Interval{
			Root,
			MajorSecond,
			MajorThird,
			Fourth,
			Fifth,
			MajorSixth,
			MajorSeventh,
		}

		actual := validateIntervals(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Returns false for unsorted intervals of a minor (aeolian) scale", func(t *testing.T) {
		expected := false
		input := []Interval{
			MinorThird,
			Fourth,
			MajorSecond,
			Fifth,
			Root,
			MinorSixth,
			MinorSeventh,
		}

		actual := validateIntervals(input)

		assert.Equal(t, expected, actual)
	})
}

func TestGenerateNotesWithinSingleOctavel(t *testing.T) {
	t.Run("Generates all notes for Lydian scale, starting from Fsharp3", func(t *testing.T) {
		expected := Notes{
			Note{
				Pitch:       Fsharp,
				PitchOctave: ThirdOctave,
			},
			Note{
				Pitch:       Gsharp,
				PitchOctave: ThirdOctave,
			},
			Note{
				Pitch:       Asharp,
				PitchOctave: ThirdOctave,
			},
			Note{
				Pitch:       C,
				PitchOctave: FourthOctave,
			},
			Note{
				Pitch:       Csharp,
				PitchOctave: FourthOctave,
			},
			Note{
				Pitch:       Dsharp,
				PitchOctave: FourthOctave,
			},
			Note{
				Pitch:       F,
				PitchOctave: FourthOctave,
			},
		}

		inputRootNote := Note{
			Pitch:       Fsharp,
			PitchOctave: ThirdOctave,
		}
		inputIntervals := []Interval{
			Root,
			MajorSecond,
			MajorThird,
			Tritone,
			Fifth,
			MajorSixth,
			MajorSeventh,
		}

		actual := generateNotesWithinSingleOctave(inputRootNote, inputIntervals)

		assert.Equal(t, expected, actual)
	})
}

func TestRaiseNoteByInterval(t *testing.T) {
	t.Run("Raise C4 by Major Third", func(t *testing.T) {
		expected := Note{
			Pitch:       E,
			PitchOctave: FourthOctave,
		}

		actual := Note{
			Pitch:       C,
			PitchOctave: FourthOctave,
		}

		actual.raiseNoteByInterval(MajorThird)

		assert.Equal(t, expected, actual)
	})

	t.Run("Raise F1 by Minor Seventh", func(t *testing.T) {
		expected := Note{
			Pitch:       Dsharp,
			PitchOctave: SecondOctave,
		}

		actual := Note{
			Pitch:       F,
			PitchOctave: FirstOctave,
		}

		actual.raiseNoteByInterval(MinorSeventh)

		assert.Equal(t, expected, actual)
	})

	t.Run("raise B5 by Minor Second", func(t *testing.T) {
		expected := Note{
			Pitch:       C,
			PitchOctave: SixthOctave,
		}

		actual := Note{
			Pitch:       B,
			PitchOctave: FifthOctave,
		}

		actual.raiseNoteByInterval(MinorSecond)

		assert.Equal(t, expected, actual)
	})
}
