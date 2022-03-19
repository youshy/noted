package notes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateNeck(t *testing.T) {
	t.Run("Generate neck for standard tuned guitar, 21 frets, 6 strings", func(t *testing.T) {
		expected := Neck{
			FirstString: String{
				0:  E,
				1:  F,
				2:  Fsharp,
				3:  G,
				4:  Gsharp,
				5:  A,
				6:  Asharp,
				7:  B,
				8:  C,
				9:  Csharp,
				10: D,
				11: Dsharp,
				12: E,
				13: F,
				14: Fsharp,
				15: G,
				16: Gsharp,
				17: A,
				18: Asharp,
				19: B,
				20: C,
				21: Csharp,
			},
			SecondString: String{
				0:  B,
				1:  C,
				2:  Csharp,
				3:  D,
				4:  Dsharp,
				5:  E,
				6:  F,
				7:  Fsharp,
				8:  G,
				9:  Gsharp,
				10: A,
				11: Asharp,
				12: B,
				13: C,
				14: Csharp,
				15: D,
				16: Dsharp,
				17: E,
				18: F,
				19: Fsharp,
				20: G,
				21: Gsharp,
			},
			ThirdString: String{
				0:  G,
				1:  Gsharp,
				2:  A,
				3:  Asharp,
				4:  B,
				5:  C,
				6:  Csharp,
				7:  D,
				8:  Dsharp,
				9:  E,
				10: F,
				11: Fsharp,
				12: G,
				13: Gsharp,
				14: A,
				15: Asharp,
				16: B,
				17: C,
				18: Csharp,
				19: D,
				20: Dsharp,
				21: E,
			},
			FourthString: String{
				0:  D,
				1:  Dsharp,
				2:  E,
				3:  F,
				4:  Fsharp,
				5:  G,
				6:  Gsharp,
				7:  A,
				8:  Asharp,
				9:  B,
				10: C,
				11: Csharp,
				12: D,
				13: Dsharp,
				14: E,
				15: F,
				16: Fsharp,
				17: G,
				18: Gsharp,
				19: A,
				20: Asharp,
				21: B,
			},
			FifthString: String{
				0:  A,
				1:  Asharp,
				2:  B,
				3:  C,
				4:  Csharp,
				5:  D,
				6:  Dsharp,
				7:  E,
				8:  F,
				9:  Fsharp,
				10: G,
				11: Gsharp,
				12: A,
				13: Asharp,
				14: B,
				15: C,
				16: Csharp,
				17: D,
				18: Dsharp,
				19: E,
				20: F,
				21: Fsharp,
			},
			SixthString: String{
				0:  E,
				1:  F,
				2:  Fsharp,
				3:  G,
				4:  Gsharp,
				5:  A,
				6:  Asharp,
				7:  B,
				8:  C,
				9:  Csharp,
				10: D,
				11: Dsharp,
				12: E,
				13: F,
				14: Fsharp,
				15: G,
				16: Gsharp,
				17: A,
				18: Asharp,
				19: B,
				20: C,
				21: Csharp,
			},
		}
		firstStringTuning := E
		secondStringTuning := B
		thirdStringTuning := G
		fourthStringTuning := D
		fifthStringTuning := A
		sixthStringTuning := E

		actual, err := GenerateNeck(
			SixString,
			TwentyOneFrets,
			firstStringTuning,
			secondStringTuning,
			thirdStringTuning,
			fourthStringTuning,
			fifthStringTuning,
			sixthStringTuning,
		)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestGenerateString(t *testing.T) {
	t.Run("Generate string tuned from E, 21 fret neck", func(t *testing.T) {
		expected := String{
			0:  E,
			1:  F,
			2:  Fsharp,
			3:  G,
			4:  Gsharp,
			5:  A,
			6:  Asharp,
			7:  B,
			8:  C,
			9:  Csharp,
			10: D,
			11: Dsharp,
			12: E,
			13: F,
			14: Fsharp,
			15: G,
			16: Gsharp,
			17: A,
			18: Asharp,
			19: B,
			20: C,
			21: Csharp,
		}
		root := E
		frets := TwentyOneFrets

		actual := generateString(root, frets)

		assert.Equal(t, expected, actual)
	})
}

func TestGenerateAllNotesFromRootWithinSingleOctave(t *testing.T) {
	t.Run("Generate notes from C", func(t *testing.T) {
		expected := Notes{
			C,
			Csharp,
			D,
			Dsharp,
			E,
			F,
			Fsharp,
			G,
			Gsharp,
			A,
			Asharp,
			B,
		}
		root := C

		actual := generateAllNotesFromRootWithinSingleOctave(root)

		assert.Equal(t, expected, actual)
	})

	t.Run("Generate notes from F", func(t *testing.T) {
		expected := Notes{
			F,
			Fsharp,
			G,
			Gsharp,
			A,
			Asharp,
			B,
			C,
			Csharp,
			D,
			Dsharp,
			E,
		}
		root := F

		actual := generateAllNotesFromRootWithinSingleOctave(root)

		assert.Equal(t, expected, actual)
	})
}
