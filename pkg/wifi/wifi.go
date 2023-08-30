package wifi

import (
	"fmt"
	"strconv"
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

	return WiFi{
		Security: security,
		SSID:     ssid,
		Password: password,
		Hidden:   hidden,
	}, nil
}
