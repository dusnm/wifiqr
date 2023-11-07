package utils

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInvalidStatesOfHexColorParsing(t *testing.T) {
	t.Parallel()

	// arrange
	assert := require.New(t)
	fixtures := []struct {
		Name                 string
		ExpectedErrorMessage string
		Input                string
	}{
		{
			Name:                 "invalid length of hex string",
			ExpectedErrorMessage: "invalid length for hex color: 3 or 6 characters required",
			Input:                "a",
		},
		{
			Name:                 "not parsable as hex color",
			ExpectedErrorMessage: "encoding/hex: invalid byte: U+0078 'x'",
			Input:                "xxxxxx",
		},
	}

	for _, fixture := range fixtures {
		t.Log(fixture.Name)

		// act
		_, err := ParseHexColor(fixture.Input)

		// assert
		assert.Errorf(err, fixture.ExpectedErrorMessage)
	}
}

func TestValidStatesOfHexColorParsing(t *testing.T) {
	t.Parallel()

	// arrange
	assert := require.New(t)
	fixtures := []struct {
		Name     string
		Expected color.RGBA
		Input    string
	}{
		{
			Name: "Hex color is 6 characters long",
			Expected: color.RGBA{
				0xea,
				0x00,
				0xeb,
				0xff,
			},
			Input: "ea00eb",
		},
		{
			Name: "Hex color is 3 characters long",
			Expected: color.RGBA{
				0xee,
				0xcc,
				0xbb,
				0xff,
			},
			Input: "ecb",
		},
	}

	for _, fixture := range fixtures {
		t.Log(fixture.Name)

		// act
		c, err := ParseHexColor(fixture.Input)

		// assert
		assert.NoError(err)
		assert.Equal(c, fixture.Expected)
	}
}
