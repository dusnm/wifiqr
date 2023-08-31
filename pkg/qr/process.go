package qr

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"math"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

const (
	fontSize = 50
	fontDPI  = 72
)

func AddHeader(
	title string,
	imageBytes []byte,
	wr io.WriteCloser,
) error {
	fg := image.NewUniform(
		color.RGBA{
			0x81,
			0x78,
			0xe4,
			0xff,
		},
	)

	fontBytes, err := resources.ReadFile("res/font.ttf")
	if err != nil {
		return err
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return err
	}

	r := bytes.NewReader(imageBytes)
	imgc, err := png.DecodeConfig(r)
	if err != nil {
		return err
	}

	if _, err = r.Seek(0, 0); err != nil {
		return err
	}

	img, err := png.Decode(r)
	if err != nil {
		return err
	}

	drawableImage, ok := img.(draw.Image)
	if !ok {
		return fmt.Errorf("image not drawable")
	}

	d := &font.Drawer{
		Dst: drawableImage,
		Src: fg,
		Face: truetype.NewFace(
			f,
			&truetype.Options{
				Size:    fontSize,
				DPI:     fontDPI,
				Hinting: font.HintingFull,
			},
		),
	}

	d.Dot = fixed.Point26_6{
		X: (fixed.I(imgc.Width) - d.MeasureString(title)) / 2,
		Y: fixed.I(int(10 + math.Ceil(fontSize*fontDPI/72))),
	}

	d.DrawString(title)

	return png.Encode(wr, drawableImage)
}
