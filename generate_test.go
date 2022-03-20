package notes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateNeck(t *testing.T) {
	t.Run("Generate neck for standard tuned guitar, 21 frets, 6 strings", func(t *testing.T) {
		expected := Neck{
			FirstString: String{
				0: Note{
					Pitch:       E,
					PitchOctave: FourthOctave,
				},
				1: Note{
					Pitch:       F,
					PitchOctave: FourthOctave,
				},
				2: Note{
					Pitch:       Fsharp,
					PitchOctave: FourthOctave,
				},
				3: Note{
					Pitch:       G,
					PitchOctave: FourthOctave,
				},
				4: Note{
					Pitch:       Gsharp,
					PitchOctave: FourthOctave,
				},
				5: Note{
					Pitch:       A,
					PitchOctave: FourthOctave,
				},
				6: Note{
					Pitch:       Asharp,
					PitchOctave: FourthOctave,
				},
				7: Note{
					Pitch:       B,
					PitchOctave: FourthOctave,
				},
				8: Note{
					Pitch:       C,
					PitchOctave: FifthOctave,
				},
				9: Note{
					Pitch:       Csharp,
					PitchOctave: FifthOctave,
				},
				10: Note{
					Pitch:       D,
					PitchOctave: FifthOctave,
				},
				11: Note{
					Pitch:       Dsharp,
					PitchOctave: FifthOctave,
				},
				12: Note{
					Pitch:       E,
					PitchOctave: FifthOctave,
				},
				13: Note{
					Pitch:       F,
					PitchOctave: FifthOctave,
				},
				14: Note{
					Pitch:       Fsharp,
					PitchOctave: FifthOctave,
				},
				15: Note{
					Pitch:       G,
					PitchOctave: FifthOctave,
				},
				16: Note{
					Pitch:       Gsharp,
					PitchOctave: FifthOctave,
				},
				17: Note{
					Pitch:       A,
					PitchOctave: FifthOctave,
				},
				18: Note{
					Pitch:       Asharp,
					PitchOctave: FifthOctave,
				},
				19: Note{
					Pitch:       B,
					PitchOctave: FifthOctave,
				},
				20: Note{
					Pitch:       C,
					PitchOctave: SixthOctave,
				},
				21: Note{
					Pitch:       Csharp,
					PitchOctave: SixthOctave,
				},
			},
			SecondString: String{
				0: Note{
					Pitch:       B,
					PitchOctave: ThirdOctave,
				},
				1: Note{
					Pitch:       C,
					PitchOctave: FourthOctave,
				},
				2: Note{
					Pitch:       Csharp,
					PitchOctave: FourthOctave,
				},
				3: Note{
					Pitch:       D,
					PitchOctave: FourthOctave,
				},
				4: Note{
					Pitch:       Dsharp,
					PitchOctave: FourthOctave,
				},
				5: Note{
					Pitch:       E,
					PitchOctave: FourthOctave,
				},
				6: Note{
					Pitch:       F,
					PitchOctave: FourthOctave,
				},
				7: Note{
					Pitch:       Fsharp,
					PitchOctave: FourthOctave,
				},
				8: Note{
					Pitch:       G,
					PitchOctave: FourthOctave,
				},
				9: Note{
					Pitch:       Gsharp,
					PitchOctave: FourthOctave,
				},
				10: Note{
					Pitch:       A,
					PitchOctave: FourthOctave,
				},
				11: Note{
					Pitch:       Asharp,
					PitchOctave: FourthOctave,
				},
				12: Note{
					Pitch:       B,
					PitchOctave: FourthOctave,
				},
				13: Note{
					Pitch:       C,
					PitchOctave: FifthOctave,
				},
				14: Note{
					Pitch:       Csharp,
					PitchOctave: FifthOctave,
				},
				15: Note{
					Pitch:       D,
					PitchOctave: FifthOctave,
				},
				16: Note{
					Pitch:       Dsharp,
					PitchOctave: FifthOctave,
				},
				17: Note{
					Pitch:       E,
					PitchOctave: FifthOctave,
				},
				18: Note{
					Pitch:       F,
					PitchOctave: FifthOctave,
				},
				19: Note{
					Pitch:       Fsharp,
					PitchOctave: FifthOctave,
				},
				20: Note{
					Pitch:       G,
					PitchOctave: FifthOctave,
				},
				21: Note{
					Pitch:       Gsharp,
					PitchOctave: FifthOctave,
				},
			},
			ThirdString: String{
				0: Note{
					Pitch:       G,
					PitchOctave: ThirdOctave,
				},
				1: Note{
					Pitch:       Gsharp,
					PitchOctave: ThirdOctave,
				},
				2: Note{
					Pitch:       A,
					PitchOctave: ThirdOctave,
				},
				3: Note{
					Pitch:       Asharp,
					PitchOctave: ThirdOctave,
				},
				4: Note{
					Pitch:       B,
					PitchOctave: ThirdOctave,
				},
				5: Note{
					Pitch:       C,
					PitchOctave: FourthOctave,
				},
				6: Note{
					Pitch:       Csharp,
					PitchOctave: FourthOctave,
				},
				7: Note{
					Pitch:       D,
					PitchOctave: FourthOctave,
				},
				8: Note{
					Pitch:       Dsharp,
					PitchOctave: FourthOctave,
				},
				9: Note{
					Pitch:       E,
					PitchOctave: FourthOctave,
				},
				10: Note{
					Pitch:       F,
					PitchOctave: FourthOctave,
				},
				11: Note{
					Pitch:       Fsharp,
					PitchOctave: FourthOctave,
				},
				12: Note{
					Pitch:       G,
					PitchOctave: FourthOctave,
				},
				13: Note{
					Pitch:       Gsharp,
					PitchOctave: FourthOctave,
				},
				14: Note{
					Pitch:       A,
					PitchOctave: FourthOctave,
				},
				15: Note{
					Pitch:       Asharp,
					PitchOctave: FourthOctave,
				},
				16: Note{
					Pitch:       B,
					PitchOctave: FourthOctave,
				},
				17: Note{
					Pitch:       C,
					PitchOctave: FifthOctave,
				},
				18: Note{
					Pitch:       Csharp,
					PitchOctave: FifthOctave,
				},
				19: Note{
					Pitch:       D,
					PitchOctave: FifthOctave,
				},
				20: Note{
					Pitch:       Dsharp,
					PitchOctave: FifthOctave,
				},
				21: Note{
					Pitch:       E,
					PitchOctave: FifthOctave,
				},
			},
			FourthString: String{
				0: Note{
					Pitch:       D,
					PitchOctave: ThirdOctave,
				},
				1: Note{
					Pitch:       Dsharp,
					PitchOctave: ThirdOctave,
				},
				2: Note{
					Pitch:       E,
					PitchOctave: ThirdOctave,
				},
				3: Note{
					Pitch:       F,
					PitchOctave: ThirdOctave,
				},
				4: Note{
					Pitch:       Fsharp,
					PitchOctave: ThirdOctave,
				},
				5: Note{
					Pitch:       G,
					PitchOctave: ThirdOctave,
				},
				6: Note{
					Pitch:       Gsharp,
					PitchOctave: ThirdOctave,
				},
				7: Note{
					Pitch:       A,
					PitchOctave: ThirdOctave,
				},
				8: Note{
					Pitch:       Asharp,
					PitchOctave: ThirdOctave,
				},
				9: Note{
					Pitch:       B,
					PitchOctave: ThirdOctave,
				},
				10: Note{
					Pitch:       C,
					PitchOctave: FourthOctave,
				},
				11: Note{
					Pitch:       Csharp,
					PitchOctave: FourthOctave,
				},
				12: Note{
					Pitch:       D,
					PitchOctave: FourthOctave,
				},
				13: Note{
					Pitch:       Dsharp,
					PitchOctave: FourthOctave,
				},
				14: Note{
					Pitch:       E,
					PitchOctave: FourthOctave,
				},
				15: Note{
					Pitch:       F,
					PitchOctave: FourthOctave,
				},
				16: Note{
					Pitch:       Fsharp,
					PitchOctave: FourthOctave,
				},
				17: Note{
					Pitch:       G,
					PitchOctave: FourthOctave,
				},
				18: Note{
					Pitch:       Gsharp,
					PitchOctave: FourthOctave,
				},
				19: Note{
					Pitch:       A,
					PitchOctave: FourthOctave,
				},
				20: Note{
					Pitch:       Asharp,
					PitchOctave: FourthOctave,
				},
				21: Note{
					Pitch:       B,
					PitchOctave: FourthOctave,
				},
			},
			FifthString: String{
				0: Note{
					Pitch:       A,
					PitchOctave: SecondOctave,
				},
				1: Note{
					Pitch:       Asharp,
					PitchOctave: SecondOctave,
				},
				2: Note{
					Pitch:       B,
					PitchOctave: SecondOctave,
				},
				3: Note{
					Pitch:       C,
					PitchOctave: ThirdOctave,
				},
				4: Note{
					Pitch:       Csharp,
					PitchOctave: ThirdOctave,
				},
				5: Note{
					Pitch:       D,
					PitchOctave: ThirdOctave,
				},
				6: Note{
					Pitch:       Dsharp,
					PitchOctave: ThirdOctave,
				},
				7: Note{
					Pitch:       E,
					PitchOctave: ThirdOctave,
				},
				8: Note{
					Pitch:       F,
					PitchOctave: ThirdOctave,
				},
				9: Note{
					Pitch:       Fsharp,
					PitchOctave: ThirdOctave,
				},
				10: Note{
					Pitch:       G,
					PitchOctave: ThirdOctave,
				},
				11: Note{
					Pitch:       Gsharp,
					PitchOctave: ThirdOctave,
				},
				12: Note{
					Pitch:       A,
					PitchOctave: ThirdOctave,
				},
				13: Note{
					Pitch:       Asharp,
					PitchOctave: ThirdOctave,
				},
				14: Note{
					Pitch:       B,
					PitchOctave: ThirdOctave,
				},
				15: Note{
					Pitch:       C,
					PitchOctave: FourthOctave,
				},
				16: Note{
					Pitch:       Csharp,
					PitchOctave: FourthOctave,
				},
				17: Note{
					Pitch:       D,
					PitchOctave: FourthOctave,
				},
				18: Note{
					Pitch:       Dsharp,
					PitchOctave: FourthOctave,
				},
				19: Note{
					Pitch:       E,
					PitchOctave: FourthOctave,
				},
				20: Note{
					Pitch:       F,
					PitchOctave: FourthOctave,
				},
				21: Note{
					Pitch:       Fsharp,
					PitchOctave: FourthOctave,
				},
			},
			SixthString: String{
				0: Note{
					Pitch:       E,
					PitchOctave: SecondOctave,
				},
				1: Note{
					Pitch:       F,
					PitchOctave: SecondOctave,
				},
				2: Note{
					Pitch:       Fsharp,
					PitchOctave: SecondOctave,
				},
				3: Note{
					Pitch:       G,
					PitchOctave: SecondOctave,
				},
				4: Note{
					Pitch:       Gsharp,
					PitchOctave: SecondOctave,
				},
				5: Note{
					Pitch:       A,
					PitchOctave: SecondOctave,
				},
				6: Note{
					Pitch:       Asharp,
					PitchOctave: SecondOctave,
				},
				7: Note{
					Pitch:       B,
					PitchOctave: SecondOctave,
				},
				8: Note{
					Pitch:       C,
					PitchOctave: ThirdOctave,
				},
				9: Note{
					Pitch:       Csharp,
					PitchOctave: ThirdOctave,
				},
				10: Note{
					Pitch:       D,
					PitchOctave: ThirdOctave,
				},
				11: Note{
					Pitch:       Dsharp,
					PitchOctave: ThirdOctave,
				},
				12: Note{
					Pitch:       E,
					PitchOctave: ThirdOctave,
				},
				13: Note{
					Pitch:       F,
					PitchOctave: ThirdOctave,
				},
				14: Note{
					Pitch:       Fsharp,
					PitchOctave: ThirdOctave,
				},
				15: Note{
					Pitch:       G,
					PitchOctave: ThirdOctave,
				},
				16: Note{
					Pitch:       Gsharp,
					PitchOctave: ThirdOctave,
				},
				17: Note{
					Pitch:       A,
					PitchOctave: ThirdOctave,
				},
				18: Note{
					Pitch:       Asharp,
					PitchOctave: ThirdOctave,
				},
				19: Note{
					Pitch:       B,
					PitchOctave: ThirdOctave,
				},
				20: Note{
					Pitch:       C,
					PitchOctave: FourthOctave,
				},
				21: Note{
					Pitch:       Csharp,
					PitchOctave: FourthOctave,
				},
			},
		}

		firstStringTuning := Note{
			Pitch:       E,
			PitchOctave: FourthOctave,
		}
		secondStringTuning := Note{
			Pitch:       B,
			PitchOctave: ThirdOctave,
		}
		thirdStringTuning := Note{
			Pitch:       G,
			PitchOctave: ThirdOctave,
		}
		fourthStringTuning := Note{
			Pitch:       D,
			PitchOctave: ThirdOctave,
		}
		fifthStringTuning := Note{
			Pitch:       A,
			PitchOctave: SecondOctave,
		}
		sixthStringTuning := Note{
			Pitch:       E,
			PitchOctave: SecondOctave,
		}

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

	t.Run("Generate neck for drop G, 24 frets, 7 strings", func(t *testing.T) {
		expected := Neck{
			FirstString: String{
				0: Note{
					Pitch:       D,
					PitchOctave: FourthOctave,
				},
				1: Note{
					Pitch:       Dsharp,
					PitchOctave: FourthOctave,
				},
				2: Note{
					Pitch:       E,
					PitchOctave: FourthOctave,
				},
				3: Note{
					Pitch:       F,
					PitchOctave: FourthOctave,
				},
				4: Note{
					Pitch:       Fsharp,
					PitchOctave: FourthOctave,
				},
				5: Note{
					Pitch:       G,
					PitchOctave: FourthOctave,
				},
				6: Note{
					Pitch:       Gsharp,
					PitchOctave: FourthOctave,
				},
				7: Note{
					Pitch:       A,
					PitchOctave: FourthOctave,
				},
				8: Note{
					Pitch:       Asharp,
					PitchOctave: FourthOctave,
				},
				9: Note{
					Pitch:       B,
					PitchOctave: FourthOctave,
				},
				10: Note{
					Pitch:       C,
					PitchOctave: FifthOctave,
				},
				11: Note{
					Pitch:       Csharp,
					PitchOctave: FifthOctave,
				},
				12: Note{
					Pitch:       D,
					PitchOctave: FifthOctave,
				},
				13: Note{
					Pitch:       Dsharp,
					PitchOctave: FifthOctave,
				},
				14: Note{
					Pitch:       E,
					PitchOctave: FifthOctave,
				},
				15: Note{
					Pitch:       F,
					PitchOctave: FifthOctave,
				},
				16: Note{
					Pitch:       Fsharp,
					PitchOctave: FifthOctave,
				},
				17: Note{
					Pitch:       G,
					PitchOctave: FifthOctave,
				},
				18: Note{
					Pitch:       Gsharp,
					PitchOctave: FifthOctave,
				},
				19: Note{
					Pitch:       A,
					PitchOctave: FifthOctave,
				},
				20: Note{
					Pitch:       Asharp,
					PitchOctave: FifthOctave,
				},
				21: Note{
					Pitch:       B,
					PitchOctave: FifthOctave,
				},
				22: Note{
					Pitch:       C,
					PitchOctave: SixthOctave,
				},
				23: Note{
					Pitch:       Csharp,
					PitchOctave: SixthOctave,
				},
				24: Note{
					Pitch:       D,
					PitchOctave: SixthOctave,
				},
			},
			SecondString: String{
				0: Note{
					Pitch:       A,
					PitchOctave: ThirdOctave,
				},
				1: Note{
					Pitch:       Asharp,
					PitchOctave: ThirdOctave,
				},
				2: Note{
					Pitch:       B,
					PitchOctave: ThirdOctave,
				},
				3: Note{
					Pitch:       C,
					PitchOctave: FourthOctave,
				},
				4: Note{
					Pitch:       Csharp,
					PitchOctave: FourthOctave,
				},
				5: Note{
					Pitch:       D,
					PitchOctave: FourthOctave,
				},
				6: Note{
					Pitch:       Dsharp,
					PitchOctave: FourthOctave,
				},
				7: Note{
					Pitch:       E,
					PitchOctave: FourthOctave,
				},
				8: Note{
					Pitch:       F,
					PitchOctave: FourthOctave,
				},
				9: Note{
					Pitch:       Fsharp,
					PitchOctave: FourthOctave,
				},
				10: Note{
					Pitch:       G,
					PitchOctave: FourthOctave,
				},
				11: Note{
					Pitch:       Gsharp,
					PitchOctave: FourthOctave,
				},
				12: Note{
					Pitch:       A,
					PitchOctave: FourthOctave,
				},
				13: Note{
					Pitch:       Asharp,
					PitchOctave: FourthOctave,
				},
				14: Note{
					Pitch:       B,
					PitchOctave: FourthOctave,
				},
				15: Note{
					Pitch:       C,
					PitchOctave: FifthOctave,
				},
				16: Note{
					Pitch:       Csharp,
					PitchOctave: FifthOctave,
				},
				17: Note{
					Pitch:       D,
					PitchOctave: FifthOctave,
				},
				18: Note{
					Pitch:       Dsharp,
					PitchOctave: FifthOctave,
				},
				19: Note{
					Pitch:       E,
					PitchOctave: FifthOctave,
				},
				20: Note{
					Pitch:       F,
					PitchOctave: FifthOctave,
				},
				21: Note{
					Pitch:       Fsharp,
					PitchOctave: FifthOctave,
				},
				22: Note{
					Pitch:       G,
					PitchOctave: FifthOctave,
				},
				23: Note{
					Pitch:       Gsharp,
					PitchOctave: FifthOctave,
				},
				24: Note{
					Pitch:       A,
					PitchOctave: FifthOctave,
				},
			},
			ThirdString: String{
				0: Note{
					Pitch:       F,
					PitchOctave: ThirdOctave,
				},
				1: Note{
					Pitch:       Fsharp,
					PitchOctave: ThirdOctave,
				},
				2: Note{
					Pitch:       G,
					PitchOctave: ThirdOctave,
				},
				3: Note{
					Pitch:       Gsharp,
					PitchOctave: ThirdOctave,
				},
				4: Note{
					Pitch:       A,
					PitchOctave: ThirdOctave,
				},
				5: Note{
					Pitch:       Asharp,
					PitchOctave: ThirdOctave,
				},
				6: Note{
					Pitch:       B,
					PitchOctave: ThirdOctave,
				},
				7: Note{
					Pitch:       C,
					PitchOctave: FourthOctave,
				},
				8: Note{
					Pitch:       Csharp,
					PitchOctave: FourthOctave,
				},
				9: Note{
					Pitch:       D,
					PitchOctave: FourthOctave,
				},
				10: Note{
					Pitch:       Dsharp,
					PitchOctave: FourthOctave,
				},
				11: Note{
					Pitch:       E,
					PitchOctave: FourthOctave,
				},
				12: Note{
					Pitch:       F,
					PitchOctave: FourthOctave,
				},
				13: Note{
					Pitch:       Fsharp,
					PitchOctave: FourthOctave,
				},
				14: Note{
					Pitch:       G,
					PitchOctave: FourthOctave,
				},
				15: Note{
					Pitch:       Gsharp,
					PitchOctave: FourthOctave,
				},
				16: Note{
					Pitch:       A,
					PitchOctave: FourthOctave,
				},
				17: Note{
					Pitch:       Asharp,
					PitchOctave: FourthOctave,
				},
				18: Note{
					Pitch:       B,
					PitchOctave: FourthOctave,
				},
				19: Note{
					Pitch:       C,
					PitchOctave: FifthOctave,
				},
				20: Note{
					Pitch:       Csharp,
					PitchOctave: FifthOctave,
				},
				21: Note{
					Pitch:       D,
					PitchOctave: FifthOctave,
				},
				22: Note{
					Pitch:       Dsharp,
					PitchOctave: FifthOctave,
				},
				23: Note{
					Pitch:       E,
					PitchOctave: FifthOctave,
				},
				24: Note{
					Pitch:       F,
					PitchOctave: FifthOctave,
				},
			},
			FourthString: String{
				0: Note{
					Pitch:       C,
					PitchOctave: ThirdOctave,
				},
				1: Note{
					Pitch:       Csharp,
					PitchOctave: ThirdOctave,
				},
				2: Note{
					Pitch:       D,
					PitchOctave: ThirdOctave,
				},
				3: Note{
					Pitch:       Dsharp,
					PitchOctave: ThirdOctave,
				},
				4: Note{
					Pitch:       E,
					PitchOctave: ThirdOctave,
				},
				5: Note{
					Pitch:       F,
					PitchOctave: ThirdOctave,
				},
				6: Note{
					Pitch:       Fsharp,
					PitchOctave: ThirdOctave,
				},
				7: Note{
					Pitch:       G,
					PitchOctave: ThirdOctave,
				},
				8: Note{
					Pitch:       Gsharp,
					PitchOctave: ThirdOctave,
				},
				9: Note{
					Pitch:       A,
					PitchOctave: ThirdOctave,
				},
				10: Note{
					Pitch:       Asharp,
					PitchOctave: ThirdOctave,
				},
				11: Note{
					Pitch:       B,
					PitchOctave: ThirdOctave,
				},
				12: Note{
					Pitch:       C,
					PitchOctave: FourthOctave,
				},
				13: Note{
					Pitch:       Csharp,
					PitchOctave: FourthOctave,
				},
				14: Note{
					Pitch:       D,
					PitchOctave: FourthOctave,
				},
				15: Note{
					Pitch:       Dsharp,
					PitchOctave: FourthOctave,
				},
				16: Note{
					Pitch:       E,
					PitchOctave: FourthOctave,
				},
				17: Note{
					Pitch:       F,
					PitchOctave: FourthOctave,
				},
				18: Note{
					Pitch:       Fsharp,
					PitchOctave: FourthOctave,
				},
				19: Note{
					Pitch:       G,
					PitchOctave: FourthOctave,
				},
				20: Note{
					Pitch:       Gsharp,
					PitchOctave: FourthOctave,
				},
				21: Note{
					Pitch:       A,
					PitchOctave: FourthOctave,
				},
				22: Note{
					Pitch:       Asharp,
					PitchOctave: FourthOctave,
				},
				23: Note{
					Pitch:       B,
					PitchOctave: FourthOctave,
				},
				24: Note{
					Pitch:       C,
					PitchOctave: FifthOctave,
				},
			},
			FifthString: String{
				0: Note{
					Pitch:       G,
					PitchOctave: SecondOctave,
				},
				1: Note{
					Pitch:       Gsharp,
					PitchOctave: SecondOctave,
				},
				2: Note{
					Pitch:       A,
					PitchOctave: SecondOctave,
				},
				3: Note{
					Pitch:       Asharp,
					PitchOctave: SecondOctave,
				},
				4: Note{
					Pitch:       B,
					PitchOctave: SecondOctave,
				},
				5: Note{
					Pitch:       C,
					PitchOctave: ThirdOctave,
				},
				6: Note{
					Pitch:       Csharp,
					PitchOctave: ThirdOctave,
				},
				7: Note{
					Pitch:       D,
					PitchOctave: ThirdOctave,
				},
				8: Note{
					Pitch:       Dsharp,
					PitchOctave: ThirdOctave,
				},
				9: Note{
					Pitch:       E,
					PitchOctave: ThirdOctave,
				},
				10: Note{
					Pitch:       F,
					PitchOctave: ThirdOctave,
				},
				11: Note{
					Pitch:       Fsharp,
					PitchOctave: ThirdOctave,
				},
				12: Note{
					Pitch:       G,
					PitchOctave: ThirdOctave,
				},
				13: Note{
					Pitch:       Gsharp,
					PitchOctave: ThirdOctave,
				},
				14: Note{
					Pitch:       A,
					PitchOctave: ThirdOctave,
				},
				15: Note{
					Pitch:       Asharp,
					PitchOctave: ThirdOctave,
				},
				16: Note{
					Pitch:       B,
					PitchOctave: ThirdOctave,
				},
				17: Note{
					Pitch:       C,
					PitchOctave: FourthOctave,
				},
				18: Note{
					Pitch:       Csharp,
					PitchOctave: FourthOctave,
				},
				19: Note{
					Pitch:       D,
					PitchOctave: FourthOctave,
				},
				20: Note{
					Pitch:       Dsharp,
					PitchOctave: FourthOctave,
				},
				21: Note{
					Pitch:       E,
					PitchOctave: FourthOctave,
				},
				22: Note{
					Pitch:       F,
					PitchOctave: FourthOctave,
				},
				23: Note{
					Pitch:       Fsharp,
					PitchOctave: FourthOctave,
				},
				24: Note{
					Pitch:       G,
					PitchOctave: FourthOctave,
				},
			},
			SixthString: String{
				0: Note{
					Pitch:       D,
					PitchOctave: SecondOctave,
				},
				1: Note{
					Pitch:       Dsharp,
					PitchOctave: SecondOctave,
				},
				2: Note{
					Pitch:       E,
					PitchOctave: SecondOctave,
				},
				3: Note{
					Pitch:       F,
					PitchOctave: SecondOctave,
				},
				4: Note{
					Pitch:       Fsharp,
					PitchOctave: SecondOctave,
				},
				5: Note{
					Pitch:       G,
					PitchOctave: SecondOctave,
				},
				6: Note{
					Pitch:       Gsharp,
					PitchOctave: SecondOctave,
				},
				7: Note{
					Pitch:       A,
					PitchOctave: SecondOctave,
				},
				8: Note{
					Pitch:       Asharp,
					PitchOctave: SecondOctave,
				},
				9: Note{
					Pitch:       B,
					PitchOctave: SecondOctave,
				},
				10: Note{
					Pitch:       C,
					PitchOctave: ThirdOctave,
				},
				11: Note{
					Pitch:       Csharp,
					PitchOctave: ThirdOctave,
				},
				12: Note{
					Pitch:       D,
					PitchOctave: ThirdOctave,
				},
				13: Note{
					Pitch:       Dsharp,
					PitchOctave: ThirdOctave,
				},
				14: Note{
					Pitch:       E,
					PitchOctave: ThirdOctave,
				},
				15: Note{
					Pitch:       F,
					PitchOctave: ThirdOctave,
				},
				16: Note{
					Pitch:       Fsharp,
					PitchOctave: ThirdOctave,
				},
				17: Note{
					Pitch:       G,
					PitchOctave: ThirdOctave,
				},
				18: Note{
					Pitch:       Gsharp,
					PitchOctave: ThirdOctave,
				},
				19: Note{
					Pitch:       A,
					PitchOctave: ThirdOctave,
				},
				20: Note{
					Pitch:       Asharp,
					PitchOctave: ThirdOctave,
				},
				21: Note{
					Pitch:       B,
					PitchOctave: ThirdOctave,
				},
				22: Note{
					Pitch:       C,
					PitchOctave: FourthOctave,
				},
				23: Note{
					Pitch:       Csharp,
					PitchOctave: FourthOctave,
				},
				24: Note{
					Pitch:       D,
					PitchOctave: FourthOctave,
				},
			},
			SeventhString: String{
				0: Note{
					Pitch:       G,
					PitchOctave: FirstOctave,
				},
				1: Note{
					Pitch:       Gsharp,
					PitchOctave: FirstOctave,
				},
				2: Note{
					Pitch:       A,
					PitchOctave: FirstOctave,
				},
				3: Note{
					Pitch:       Asharp,
					PitchOctave: FirstOctave,
				},
				4: Note{
					Pitch:       B,
					PitchOctave: FirstOctave,
				},
				5: Note{
					Pitch:       C,
					PitchOctave: SecondOctave,
				},
				6: Note{
					Pitch:       Csharp,
					PitchOctave: SecondOctave,
				},
				7: Note{
					Pitch:       D,
					PitchOctave: SecondOctave,
				},
				8: Note{
					Pitch:       Dsharp,
					PitchOctave: SecondOctave,
				},
				9: Note{
					Pitch:       E,
					PitchOctave: SecondOctave,
				},
				10: Note{
					Pitch:       F,
					PitchOctave: SecondOctave,
				},
				11: Note{
					Pitch:       Fsharp,
					PitchOctave: SecondOctave,
				},
				12: Note{
					Pitch:       G,
					PitchOctave: SecondOctave,
				},
				13: Note{
					Pitch:       Gsharp,
					PitchOctave: SecondOctave,
				},
				14: Note{
					Pitch:       A,
					PitchOctave: SecondOctave,
				},
				15: Note{
					Pitch:       Asharp,
					PitchOctave: SecondOctave,
				},
				16: Note{
					Pitch:       B,
					PitchOctave: SecondOctave,
				},
				17: Note{
					Pitch:       C,
					PitchOctave: ThirdOctave,
				},
				18: Note{
					Pitch:       Csharp,
					PitchOctave: ThirdOctave,
				},
				19: Note{
					Pitch:       D,
					PitchOctave: ThirdOctave,
				},
				20: Note{
					Pitch:       Dsharp,
					PitchOctave: ThirdOctave,
				},
				21: Note{
					Pitch:       E,
					PitchOctave: ThirdOctave,
				},
				22: Note{
					Pitch:       F,
					PitchOctave: ThirdOctave,
				},
				23: Note{
					Pitch:       Fsharp,
					PitchOctave: ThirdOctave,
				},
				24: Note{
					Pitch:       G,
					PitchOctave: ThirdOctave,
				},
			},
		}

		firstStringTuning := Note{
			Pitch:       D,
			PitchOctave: FourthOctave,
		}
		secondStringTuning := Note{
			Pitch:       A,
			PitchOctave: ThirdOctave,
		}
		thirdStringTuning := Note{
			Pitch:       F,
			PitchOctave: ThirdOctave,
		}
		fourthStringTuning := Note{
			Pitch:       C,
			PitchOctave: ThirdOctave,
		}
		fifthStringTuning := Note{
			Pitch:       G,
			PitchOctave: SecondOctave,
		}
		sixthStringTuning := Note{
			Pitch:       D,
			PitchOctave: SecondOctave,
		}
		seventhStringTuning := Note{
			Pitch:       G,
			PitchOctave: FirstOctave,
		}

		actual, err := GenerateNeck(
			SevenString,
			TwentyFourFrets,
			firstStringTuning,
			secondStringTuning,
			thirdStringTuning,
			fourthStringTuning,
			fifthStringTuning,
			sixthStringTuning,
			seventhStringTuning,
		)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Generate neck for half-step down, 22 frets, 8 strings", func(t *testing.T) {
		expected := Neck{
			FirstString: String{
				0: Note{
					Pitch:       Dsharp,
					PitchOctave: FourthOctave,
				},
				1: Note{
					Pitch:       E,
					PitchOctave: FourthOctave,
				},
				2: Note{
					Pitch:       F,
					PitchOctave: FourthOctave,
				},
				3: Note{
					Pitch:       Fsharp,
					PitchOctave: FourthOctave,
				},
				4: Note{
					Pitch:       G,
					PitchOctave: FourthOctave,
				},
				5: Note{
					Pitch:       Gsharp,
					PitchOctave: FourthOctave,
				},
				6: Note{
					Pitch:       A,
					PitchOctave: FourthOctave,
				},
				7: Note{
					Pitch:       Asharp,
					PitchOctave: FourthOctave,
				},
				8: Note{
					Pitch:       B,
					PitchOctave: FourthOctave,
				},
				9: Note{
					Pitch:       C,
					PitchOctave: FifthOctave,
				},
				10: Note{
					Pitch:       Csharp,
					PitchOctave: FifthOctave,
				},
				11: Note{
					Pitch:       D,
					PitchOctave: FifthOctave,
				},
				12: Note{
					Pitch:       Dsharp,
					PitchOctave: FifthOctave,
				},
				13: Note{
					Pitch:       E,
					PitchOctave: FifthOctave,
				},
				14: Note{
					Pitch:       F,
					PitchOctave: FifthOctave,
				},
				15: Note{
					Pitch:       Fsharp,
					PitchOctave: FifthOctave,
				},
				16: Note{
					Pitch:       G,
					PitchOctave: FifthOctave,
				},
				17: Note{
					Pitch:       Gsharp,
					PitchOctave: FifthOctave,
				},
				18: Note{
					Pitch:       A,
					PitchOctave: FifthOctave,
				},
				19: Note{
					Pitch:       Asharp,
					PitchOctave: FifthOctave,
				},
				20: Note{
					Pitch:       B,
					PitchOctave: FifthOctave,
				},
				21: Note{
					Pitch:       C,
					PitchOctave: SixthOctave,
				},
				22: Note{
					Pitch:       Csharp,
					PitchOctave: SixthOctave,
				},
			},
			SecondString: String{
				0: Note{
					Pitch:       Asharp,
					PitchOctave: ThirdOctave,
				},
				1: Note{
					Pitch:       B,
					PitchOctave: ThirdOctave,
				},
				2: Note{
					Pitch:       C,
					PitchOctave: FourthOctave,
				},
				3: Note{
					Pitch:       Csharp,
					PitchOctave: FourthOctave,
				},
				4: Note{
					Pitch:       D,
					PitchOctave: FourthOctave,
				},
				5: Note{
					Pitch:       Dsharp,
					PitchOctave: FourthOctave,
				},
				6: Note{
					Pitch:       E,
					PitchOctave: FourthOctave,
				},
				7: Note{
					Pitch:       F,
					PitchOctave: FourthOctave,
				},
				8: Note{
					Pitch:       Fsharp,
					PitchOctave: FourthOctave,
				},
				9: Note{
					Pitch:       G,
					PitchOctave: FourthOctave,
				},
				10: Note{
					Pitch:       Gsharp,
					PitchOctave: FourthOctave,
				},
				11: Note{
					Pitch:       A,
					PitchOctave: FourthOctave,
				},
				12: Note{
					Pitch:       Asharp,
					PitchOctave: FourthOctave,
				},
				13: Note{
					Pitch:       B,
					PitchOctave: FourthOctave,
				},
				14: Note{
					Pitch:       C,
					PitchOctave: FifthOctave,
				},
				15: Note{
					Pitch:       Csharp,
					PitchOctave: FifthOctave,
				},
				16: Note{
					Pitch:       D,
					PitchOctave: FifthOctave,
				},
				17: Note{
					Pitch:       Dsharp,
					PitchOctave: FifthOctave,
				},
				18: Note{
					Pitch:       E,
					PitchOctave: FifthOctave,
				},
				19: Note{
					Pitch:       F,
					PitchOctave: FifthOctave,
				},
				20: Note{
					Pitch:       Fsharp,
					PitchOctave: FifthOctave,
				},
				21: Note{
					Pitch:       G,
					PitchOctave: FifthOctave,
				},
				22: Note{
					Pitch:       Gsharp,
					PitchOctave: FifthOctave,
				},
			},
			ThirdString: String{
				0: Note{
					Pitch:       Fsharp,
					PitchOctave: ThirdOctave,
				},
				1: Note{
					Pitch:       G,
					PitchOctave: ThirdOctave,
				},
				2: Note{
					Pitch:       Gsharp,
					PitchOctave: ThirdOctave,
				},
				3: Note{
					Pitch:       A,
					PitchOctave: ThirdOctave,
				},
				4: Note{
					Pitch:       Asharp,
					PitchOctave: ThirdOctave,
				},
				5: Note{
					Pitch:       B,
					PitchOctave: ThirdOctave,
				},
				6: Note{
					Pitch:       C,
					PitchOctave: FourthOctave,
				},
				7: Note{
					Pitch:       Csharp,
					PitchOctave: FourthOctave,
				},
				8: Note{
					Pitch:       D,
					PitchOctave: FourthOctave,
				},
				9: Note{
					Pitch:       Dsharp,
					PitchOctave: FourthOctave,
				},
				10: Note{
					Pitch:       E,
					PitchOctave: FourthOctave,
				},
				11: Note{
					Pitch:       F,
					PitchOctave: FourthOctave,
				},
				12: Note{
					Pitch:       Fsharp,
					PitchOctave: FourthOctave,
				},
				13: Note{
					Pitch:       G,
					PitchOctave: FourthOctave,
				},
				14: Note{
					Pitch:       Gsharp,
					PitchOctave: FourthOctave,
				},
				15: Note{
					Pitch:       A,
					PitchOctave: FourthOctave,
				},
				16: Note{
					Pitch:       Asharp,
					PitchOctave: FourthOctave,
				},
				17: Note{
					Pitch:       B,
					PitchOctave: FourthOctave,
				},
				18: Note{
					Pitch:       C,
					PitchOctave: FifthOctave,
				},
				19: Note{
					Pitch:       Csharp,
					PitchOctave: FifthOctave,
				},
				20: Note{
					Pitch:       D,
					PitchOctave: FifthOctave,
				},
				21: Note{
					Pitch:       Dsharp,
					PitchOctave: FifthOctave,
				},
				22: Note{
					Pitch:       E,
					PitchOctave: FifthOctave,
				},
			},
			FourthString: String{
				0: Note{
					Pitch:       Csharp,
					PitchOctave: ThirdOctave,
				},
				1: Note{
					Pitch:       D,
					PitchOctave: ThirdOctave,
				},
				2: Note{
					Pitch:       Dsharp,
					PitchOctave: ThirdOctave,
				},
				3: Note{
					Pitch:       E,
					PitchOctave: ThirdOctave,
				},
				4: Note{
					Pitch:       F,
					PitchOctave: ThirdOctave,
				},
				5: Note{
					Pitch:       Fsharp,
					PitchOctave: ThirdOctave,
				},
				6: Note{
					Pitch:       G,
					PitchOctave: ThirdOctave,
				},
				7: Note{
					Pitch:       Gsharp,
					PitchOctave: ThirdOctave,
				},
				8: Note{
					Pitch:       A,
					PitchOctave: ThirdOctave,
				},
				9: Note{
					Pitch:       Asharp,
					PitchOctave: ThirdOctave,
				},
				10: Note{
					Pitch:       B,
					PitchOctave: ThirdOctave,
				},
				11: Note{
					Pitch:       C,
					PitchOctave: FourthOctave,
				},
				12: Note{
					Pitch:       Csharp,
					PitchOctave: FourthOctave,
				},
				13: Note{
					Pitch:       D,
					PitchOctave: FourthOctave,
				},
				14: Note{
					Pitch:       Dsharp,
					PitchOctave: FourthOctave,
				},
				15: Note{
					Pitch:       E,
					PitchOctave: FourthOctave,
				},
				16: Note{
					Pitch:       F,
					PitchOctave: FourthOctave,
				},
				17: Note{
					Pitch:       Fsharp,
					PitchOctave: FourthOctave,
				},
				18: Note{
					Pitch:       G,
					PitchOctave: FourthOctave,
				},
				19: Note{
					Pitch:       Gsharp,
					PitchOctave: FourthOctave,
				},
				20: Note{
					Pitch:       A,
					PitchOctave: FourthOctave,
				},
				21: Note{
					Pitch:       Asharp,
					PitchOctave: FourthOctave,
				},
				22: Note{
					Pitch:       B,
					PitchOctave: FourthOctave,
				},
			},
			FifthString: String{
				0: Note{
					Pitch:       Gsharp,
					PitchOctave: SecondOctave,
				},
				1: Note{
					Pitch:       A,
					PitchOctave: SecondOctave,
				},
				2: Note{
					Pitch:       Asharp,
					PitchOctave: SecondOctave,
				},
				3: Note{
					Pitch:       B,
					PitchOctave: SecondOctave,
				},
				4: Note{
					Pitch:       C,
					PitchOctave: ThirdOctave,
				},
				5: Note{
					Pitch:       Csharp,
					PitchOctave: ThirdOctave,
				},
				6: Note{
					Pitch:       D,
					PitchOctave: ThirdOctave,
				},
				7: Note{
					Pitch:       Dsharp,
					PitchOctave: ThirdOctave,
				},
				8: Note{
					Pitch:       E,
					PitchOctave: ThirdOctave,
				},
				9: Note{
					Pitch:       F,
					PitchOctave: ThirdOctave,
				},
				10: Note{
					Pitch:       Fsharp,
					PitchOctave: ThirdOctave,
				},
				11: Note{
					Pitch:       G,
					PitchOctave: ThirdOctave,
				},
				12: Note{
					Pitch:       Gsharp,
					PitchOctave: ThirdOctave,
				},
				13: Note{
					Pitch:       A,
					PitchOctave: ThirdOctave,
				},
				14: Note{
					Pitch:       Asharp,
					PitchOctave: ThirdOctave,
				},
				15: Note{
					Pitch:       B,
					PitchOctave: ThirdOctave,
				},
				16: Note{
					Pitch:       C,
					PitchOctave: FourthOctave,
				},
				17: Note{
					Pitch:       Csharp,
					PitchOctave: FourthOctave,
				},
				18: Note{
					Pitch:       D,
					PitchOctave: FourthOctave,
				},
				19: Note{
					Pitch:       Dsharp,
					PitchOctave: FourthOctave,
				},
				20: Note{
					Pitch:       E,
					PitchOctave: FourthOctave,
				},
				21: Note{
					Pitch:       F,
					PitchOctave: FourthOctave,
				},
				22: Note{
					Pitch:       Fsharp,
					PitchOctave: FourthOctave,
				},
			},
			SixthString: String{
				0: Note{
					Pitch:       Dsharp,
					PitchOctave: SecondOctave,
				},
				1: Note{
					Pitch:       E,
					PitchOctave: SecondOctave,
				},
				2: Note{
					Pitch:       F,
					PitchOctave: SecondOctave,
				},
				3: Note{
					Pitch:       Fsharp,
					PitchOctave: SecondOctave,
				},
				4: Note{
					Pitch:       G,
					PitchOctave: SecondOctave,
				},
				5: Note{
					Pitch:       Gsharp,
					PitchOctave: SecondOctave,
				},
				6: Note{
					Pitch:       A,
					PitchOctave: SecondOctave,
				},
				7: Note{
					Pitch:       Asharp,
					PitchOctave: SecondOctave,
				},
				8: Note{
					Pitch:       B,
					PitchOctave: SecondOctave,
				},
				9: Note{
					Pitch:       C,
					PitchOctave: ThirdOctave,
				},
				10: Note{
					Pitch:       Csharp,
					PitchOctave: ThirdOctave,
				},
				11: Note{
					Pitch:       D,
					PitchOctave: ThirdOctave,
				},
				12: Note{
					Pitch:       Dsharp,
					PitchOctave: ThirdOctave,
				},
				13: Note{
					Pitch:       E,
					PitchOctave: ThirdOctave,
				},
				14: Note{
					Pitch:       F,
					PitchOctave: ThirdOctave,
				},
				15: Note{
					Pitch:       Fsharp,
					PitchOctave: ThirdOctave,
				},
				16: Note{
					Pitch:       G,
					PitchOctave: ThirdOctave,
				},
				17: Note{
					Pitch:       Gsharp,
					PitchOctave: ThirdOctave,
				},
				18: Note{
					Pitch:       A,
					PitchOctave: ThirdOctave,
				},
				19: Note{
					Pitch:       Asharp,
					PitchOctave: ThirdOctave,
				},
				20: Note{
					Pitch:       B,
					PitchOctave: ThirdOctave,
				},
				21: Note{
					Pitch:       C,
					PitchOctave: FourthOctave,
				},
				22: Note{
					Pitch:       Csharp,
					PitchOctave: FourthOctave,
				},
			},
			SeventhString: String{
				0: Note{
					Pitch:       Asharp,
					PitchOctave: FirstOctave,
				},
				1: Note{
					Pitch:       B,
					PitchOctave: FirstOctave,
				},
				2: Note{
					Pitch:       C,
					PitchOctave: SecondOctave,
				},
				3: Note{
					Pitch:       Csharp,
					PitchOctave: SecondOctave,
				},
				4: Note{
					Pitch:       D,
					PitchOctave: SecondOctave,
				},
				5: Note{
					Pitch:       Dsharp,
					PitchOctave: SecondOctave,
				},
				6: Note{
					Pitch:       E,
					PitchOctave: SecondOctave,
				},
				7: Note{
					Pitch:       F,
					PitchOctave: SecondOctave,
				},
				8: Note{
					Pitch:       Fsharp,
					PitchOctave: SecondOctave,
				},
				9: Note{
					Pitch:       G,
					PitchOctave: SecondOctave,
				},
				10: Note{
					Pitch:       Gsharp,
					PitchOctave: SecondOctave,
				},
				11: Note{
					Pitch:       A,
					PitchOctave: SecondOctave,
				},
				12: Note{
					Pitch:       Asharp,
					PitchOctave: SecondOctave,
				},
				13: Note{
					Pitch:       B,
					PitchOctave: SecondOctave,
				},
				14: Note{
					Pitch:       C,
					PitchOctave: ThirdOctave,
				},
				15: Note{
					Pitch:       Csharp,
					PitchOctave: ThirdOctave,
				},
				16: Note{
					Pitch:       D,
					PitchOctave: ThirdOctave,
				},
				17: Note{
					Pitch:       Dsharp,
					PitchOctave: ThirdOctave,
				},
				18: Note{
					Pitch:       E,
					PitchOctave: ThirdOctave,
				},
				19: Note{
					Pitch:       F,
					PitchOctave: ThirdOctave,
				},
				20: Note{
					Pitch:       Fsharp,
					PitchOctave: ThirdOctave,
				},
				21: Note{
					Pitch:       G,
					PitchOctave: ThirdOctave,
				},
				22: Note{
					Pitch:       Gsharp,
					PitchOctave: ThirdOctave,
				},
			},
			EightString: String{
				0: Note{
					Pitch:       F,
					PitchOctave: FirstOctave,
				},
				1: Note{
					Pitch:       Fsharp,
					PitchOctave: FirstOctave,
				},
				2: Note{
					Pitch:       G,
					PitchOctave: FirstOctave,
				},
				3: Note{
					Pitch:       Gsharp,
					PitchOctave: FirstOctave,
				},
				4: Note{
					Pitch:       A,
					PitchOctave: FirstOctave,
				},
				5: Note{
					Pitch:       Asharp,
					PitchOctave: FirstOctave,
				},
				6: Note{
					Pitch:       B,
					PitchOctave: FirstOctave,
				},
				7: Note{
					Pitch:       C,
					PitchOctave: SecondOctave,
				},
				8: Note{
					Pitch:       Csharp,
					PitchOctave: SecondOctave,
				},
				9: Note{
					Pitch:       D,
					PitchOctave: SecondOctave,
				},
				10: Note{
					Pitch:       Dsharp,
					PitchOctave: SecondOctave,
				},
				11: Note{
					Pitch:       E,
					PitchOctave: SecondOctave,
				},
				12: Note{
					Pitch:       F,
					PitchOctave: SecondOctave,
				},
				13: Note{
					Pitch:       Fsharp,
					PitchOctave: SecondOctave,
				},
				14: Note{
					Pitch:       G,
					PitchOctave: SecondOctave,
				},
				15: Note{
					Pitch:       Gsharp,
					PitchOctave: SecondOctave,
				},
				16: Note{
					Pitch:       A,
					PitchOctave: SecondOctave,
				},
				17: Note{
					Pitch:       Asharp,
					PitchOctave: SecondOctave,
				},
				18: Note{
					Pitch:       B,
					PitchOctave: SecondOctave,
				},
				19: Note{
					Pitch:       C,
					PitchOctave: ThirdOctave,
				},
				20: Note{
					Pitch:       Csharp,
					PitchOctave: ThirdOctave,
				},
				21: Note{
					Pitch:       D,
					PitchOctave: ThirdOctave,
				},
				22: Note{
					Pitch:       Dsharp,
					PitchOctave: ThirdOctave,
				},
			},
		}

		firstStringTuning := Note{
			Pitch:       Dsharp,
			PitchOctave: FourthOctave,
		}
		secondStringTuning := Note{
			Pitch:       Asharp,
			PitchOctave: ThirdOctave,
		}
		thirdStringTuning := Note{
			Pitch:       Fsharp,
			PitchOctave: ThirdOctave,
		}
		fourthStringTuning := Note{
			Pitch:       Csharp,
			PitchOctave: ThirdOctave,
		}
		fifthStringTuning := Note{
			Pitch:       Gsharp,
			PitchOctave: SecondOctave,
		}
		sixthStringTuning := Note{
			Pitch:       Dsharp,
			PitchOctave: SecondOctave,
		}
		seventhStringTuning := Note{
			Pitch:       Asharp,
			PitchOctave: FirstOctave,
		}
		eightStringTuning := Note{
			Pitch:       F,
			PitchOctave: FirstOctave,
		}

		actual, err := GenerateNeck(
			EightString,
			TwentyTwoFrets,
			firstStringTuning,
			secondStringTuning,
			thirdStringTuning,
			fourthStringTuning,
			fifthStringTuning,
			sixthStringTuning,
			seventhStringTuning,
			eightStringTuning,
		)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestGenerateString(t *testing.T) {
	t.Run("Generate string tuned from E, 21 fret neck, first string", func(t *testing.T) {
		expected := String{
			0: Note{
				Pitch:       E,
				PitchOctave: FourthOctave,
			},
			1: Note{
				Pitch:       F,
				PitchOctave: FourthOctave,
			},
			2: Note{
				Pitch:       Fsharp,
				PitchOctave: FourthOctave,
			},
			3: Note{
				Pitch:       G,
				PitchOctave: FourthOctave,
			},
			4: Note{
				Pitch:       Gsharp,
				PitchOctave: FourthOctave,
			},
			5: Note{
				Pitch:       A,
				PitchOctave: FourthOctave,
			},
			6: Note{
				Pitch:       Asharp,
				PitchOctave: FourthOctave,
			},
			7: Note{
				Pitch:       B,
				PitchOctave: FourthOctave,
			},
			8: Note{
				Pitch:       C,
				PitchOctave: FifthOctave,
			},
			9: Note{
				Pitch:       Csharp,
				PitchOctave: FifthOctave,
			},
			10: Note{
				Pitch:       D,
				PitchOctave: FifthOctave,
			},
			11: Note{
				Pitch:       Dsharp,
				PitchOctave: FifthOctave,
			},
			12: Note{
				Pitch:       E,
				PitchOctave: FifthOctave,
			},
			13: Note{
				Pitch:       F,
				PitchOctave: FifthOctave,
			},
			14: Note{
				Pitch:       Fsharp,
				PitchOctave: FifthOctave,
			},
			15: Note{
				Pitch:       G,
				PitchOctave: FifthOctave,
			},
			16: Note{
				Pitch:       Gsharp,
				PitchOctave: FifthOctave,
			},
			17: Note{
				Pitch:       A,
				PitchOctave: FifthOctave,
			},
			18: Note{
				Pitch:       Asharp,
				PitchOctave: FifthOctave,
			},
			19: Note{
				Pitch:       B,
				PitchOctave: FifthOctave,
			},
			20: Note{
				Pitch:       C,
				PitchOctave: SixthOctave,
			},
			21: Note{
				Pitch:       Csharp,
				PitchOctave: SixthOctave,
			},
		}
		root := Note{
			Pitch:       E,
			PitchOctave: FourthOctave,
		}
		frets := TwentyOneFrets

		actual := generateString(root, frets)

		assert.Equal(t, expected, actual)
	})

	t.Run("Generate string tuned from C, 24 fret neck, sixth string", func(t *testing.T) {
		expected := String{
			0: Note{
				Pitch:       C,
				PitchOctave: SecondOctave,
			},
			1: Note{
				Pitch:       Csharp,
				PitchOctave: SecondOctave,
			},
			2: Note{
				Pitch:       D,
				PitchOctave: SecondOctave,
			},
			3: Note{
				Pitch:       Dsharp,
				PitchOctave: SecondOctave,
			},
			4: Note{
				Pitch:       E,
				PitchOctave: SecondOctave,
			},
			5: Note{
				Pitch:       F,
				PitchOctave: SecondOctave,
			},
			6: Note{
				Pitch:       Fsharp,
				PitchOctave: SecondOctave,
			},
			7: Note{
				Pitch:       G,
				PitchOctave: SecondOctave,
			},
			8: Note{
				Pitch:       Gsharp,
				PitchOctave: SecondOctave,
			},
			9: Note{
				Pitch:       A,
				PitchOctave: SecondOctave,
			},
			10: Note{
				Pitch:       Asharp,
				PitchOctave: SecondOctave,
			},
			11: Note{
				Pitch:       B,
				PitchOctave: SecondOctave,
			},
			12: Note{
				Pitch:       C,
				PitchOctave: ThirdOctave,
			},
			13: Note{
				Pitch:       Csharp,
				PitchOctave: ThirdOctave,
			},
			14: Note{
				Pitch:       D,
				PitchOctave: ThirdOctave,
			},
			15: Note{
				Pitch:       Dsharp,
				PitchOctave: ThirdOctave,
			},
			16: Note{
				Pitch:       E,
				PitchOctave: ThirdOctave,
			},
			17: Note{
				Pitch:       F,
				PitchOctave: ThirdOctave,
			},
			18: Note{
				Pitch:       Fsharp,
				PitchOctave: ThirdOctave,
			},
			19: Note{
				Pitch:       G,
				PitchOctave: ThirdOctave,
			},
			20: Note{
				Pitch:       Gsharp,
				PitchOctave: ThirdOctave,
			},
			21: Note{
				Pitch:       A,
				PitchOctave: ThirdOctave,
			},
			22: Note{
				Pitch:       Asharp,
				PitchOctave: ThirdOctave,
			},
			23: Note{
				Pitch:       B,
				PitchOctave: ThirdOctave,
			},
			24: Note{
				Pitch:       C,
				PitchOctave: FourthOctave,
			},
		}
		root := Note{
			Pitch:       C,
			PitchOctave: SecondOctave,
		}
		frets := TwentyFourFrets

		actual := generateString(root, frets)

		assert.Equal(t, expected, actual)
	})
}

func TestGenerateAllNotesFromRootWithinSingleOctave(t *testing.T) {
	t.Run("Generate notes from C", func(t *testing.T) {
		expected := Notes{
			Note{
				Pitch:       C,
				PitchOctave: SecondOctave,
			},
			Note{
				Pitch:       Csharp,
				PitchOctave: SecondOctave,
			},
			Note{
				Pitch:       D,
				PitchOctave: SecondOctave,
			},
			Note{
				Pitch:       Dsharp,
				PitchOctave: SecondOctave,
			},
			Note{
				Pitch:       E,
				PitchOctave: SecondOctave,
			},
			Note{
				Pitch:       F,
				PitchOctave: SecondOctave,
			},
			Note{
				Pitch:       Fsharp,
				PitchOctave: SecondOctave,
			},
			Note{
				Pitch:       G,
				PitchOctave: SecondOctave,
			},
			Note{
				Pitch:       Gsharp,
				PitchOctave: SecondOctave,
			},
			Note{
				Pitch:       A,
				PitchOctave: SecondOctave,
			},
			Note{
				Pitch:       Asharp,
				PitchOctave: SecondOctave,
			},
			Note{
				Pitch:       B,
				PitchOctave: SecondOctave,
			},
		}
		root := Note{
			Pitch:       C,
			PitchOctave: SecondOctave,
		}

		actual := generateAllNotesFromRootWithinSingleOctave(root)

		assert.Equal(t, expected, actual)
	})

	t.Run("Generate notes from F, fourth octave", func(t *testing.T) {
		expected := Notes{
			Note{
				Pitch:       F,
				PitchOctave: FourthOctave,
			},
			Note{
				Pitch:       Fsharp,
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
				Pitch:       A,
				PitchOctave: FourthOctave,
			},
			Note{
				Pitch:       Asharp,
				PitchOctave: FourthOctave,
			},
			Note{
				Pitch:       B,
				PitchOctave: FourthOctave,
			},
			Note{
				Pitch:       C,
				PitchOctave: FifthOctave,
			},
			Note{
				Pitch:       Csharp,
				PitchOctave: FifthOctave,
			},
			Note{
				Pitch:       D,
				PitchOctave: FifthOctave,
			},
			Note{
				Pitch:       Dsharp,
				PitchOctave: FifthOctave,
			},
			Note{
				Pitch:       E,
				PitchOctave: FifthOctave,
			},
		}

		root := Note{
			Pitch:       F,
			PitchOctave: FourthOctave,
		}

		actual := generateAllNotesFromRootWithinSingleOctave(root)

		assert.Equal(t, expected, actual)
	})
}
