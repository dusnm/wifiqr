package utils

import (
	"encoding/hex"
	"fmt"
	"image/color"
	"strings"
)

func ParseHexColor(hexColor string) (color.RGBA, error) {
	hexColor = strings.ToLower(hexColor)
	length := len(hexColor)
	if length != 3 && length != 6 {
		return color.RGBA{}, fmt.Errorf("invalid length for hex color: 3 or 6 characters required")
	}

	if length == 3 {
		builder := strings.Builder{}
		for _, c := range hexColor {
			builder.WriteRune(c)
			builder.WriteRune(c)
		}

		hexColor = builder.String()
	}

	decoded, err := hex.DecodeString(hexColor)
	if err != nil {
		return color.RGBA{}, err
	}

	return color.RGBA{
		decoded[0],
		decoded[1],
		decoded[2],
		0xff,
	}, nil
}
