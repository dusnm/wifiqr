package qr

import (
	"bytes"
	"fmt"
	"image"
	"io"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

type (
	nopCloser struct {
		io.Writer
	}
)

func (nopCloser) Close() error {
	return nil
}

func Generate(data fmt.Stringer) ([]byte, error) {
	qcode, err := qrcode.NewWith(
		data.String(),
		qrcode.WithEncodingMode(qrcode.EncModeByte),
		qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionHighest),
	)

	if err != nil {
		return nil, err
	}

	logof, err := resources.ReadFile("res/logo.png")
	if err != nil {
		return nil, err
	}

	logo, _, err := image.Decode(bytes.NewReader(logof))
	if err != nil {
		return nil, err
	}

	b := bytes.NewBuffer(nil)
	w := standard.NewWithWriter(
		nopCloser{Writer: b},
		// top, right, bottom, left
		standard.WithBorderWidth(90, 40, 40, 40),
		standard.WithCircleShape(),
		standard.WithBuiltinImageEncoder(standard.PNG_FORMAT),
		standard.WithLogoImage(logo),
	)

	if err = qcode.Save(w); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
