package utils

import (
	"fmt"
	"image"
	"os"
)

func MakeLogoFromPath(logoPath string) (image.Image, error) {
	logof, err := os.Open(logoPath)
	if err != nil {
		return nil, err
	}

	defer logof.Close()

	info, _, err := image.DecodeConfig(logof)
	if err != nil {
		return nil, err
	}

	if info.Width > 140 || info.Height > 140 {
		return nil, fmt.Errorf("logo dimensions must not exceed 140x140")
	}

	if _, err = logof.Seek(0, 0); err != nil {
		return nil, err
	}

	img, _, err := image.Decode(logof)
	if err != nil {
		return nil, err
	}

	return img, nil
}
