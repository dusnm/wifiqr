package qr

import (
	"fmt"
	"io"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func Generate(data fmt.Stringer, wr io.WriteCloser) error {
	qcode, err := qrcode.NewWith(
		data.String(),
		qrcode.WithEncodingMode(qrcode.EncModeByte),
		qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionHighest),
	)

	if err != nil {
		return err
	}

	w := standard.NewWithWriter(
		wr,
		standard.WithQRWidth(40),
		standard.WithBuiltinImageEncoder(standard.PNG_FORMAT),
	)

	return qcode.Save(w)
}
