package wifi

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

const (
	SEC_NONE = "nopass"
	SEC_WEP  = "WEP"
	SEC_WPA  = "WPA"
)

type (
	WiFi struct {
		Security string
		SSID     string
		Password string
		Hidden   bool
	}
)

func (w WiFi) String() string {
	template := "WIFI:T:%s;S:%s;P:%s;H:%s;;"

	return fmt.Sprintf(
		template,
		w.Security,
		w.SSID,
		w.Password,
		strconv.FormatBool(w.Hidden),
	)
}

func New(
	security string,
	ssid string,
	password string,
	hidden bool,
) (WiFi, error) {
	if security != SEC_NONE &&
		security != SEC_WEP &&
		security != SEC_WPA {
		return WiFi{}, fmt.Errorf("invalid security: %s", security)
	}

	if ssid == "" {
		return WiFi{}, fmt.Errorf("ssid cannot be empty")
	}

	ssid = escapeHexString(escapeSpecialCharacters(ssid))
	password = escapeHexString(escapeSpecialCharacters(password))

	return WiFi{
		Security: security,
		SSID:     ssid,
		Password: password,
		Hidden:   hidden,
	}, nil
}

func escapeHexString(input string) string {
	if len(input)%2 != 0 {
		return input
	}

	if _, err := hex.DecodeString(input); err != nil {
		return input
	}

	// The string can be interpreted as a hex number
	// therefore we wrap it in double quotes
	builder := strings.Builder{}
	builder.WriteRune('"')
	builder.WriteString(input)
	builder.WriteRune('"')

	return builder.String()
}

func escapeSpecialCharacters(input string) string {
	builder := strings.Builder{}
	for _, c := range input {
		switch c {
		case '\\', ';', ',', '"', ':':
			builder.WriteRune('\\')
			fallthrough
		default:
			builder.WriteRune(c)
		}
	}

	return builder.String()
}
