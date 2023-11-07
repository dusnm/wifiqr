package wifi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type (
	input struct {
		Security string
		SSID     string
		Password string
		Hidden   bool
	}
)

func TestWifiStringEncodingInvalidCases(t *testing.T) {
	t.Parallel()

	// arrange
	assert := require.New(t)
	fixtures := []struct {
		Name                 string
		ExpectedErrorMessage string
		Input                input
	}{
		{
			Name:                 "Security left out entirely",
			ExpectedErrorMessage: "invalid security: ",
			Input: input{
				Security: "",
				SSID:     "Journey Through Pressure",
				Password: "",
				Hidden:   false,
			},
		},
		{
			Name:                 "Security is not one of the defined values",
			ExpectedErrorMessage: "invalid security: Teargas",
			Input: input{
				Security: "Teargas",
				SSID:     "I Am Nothing",
				Password: "",
				Hidden:   false,
			},
		},
		{
			Name:                 "SSID left out",
			ExpectedErrorMessage: "ssid cannot be empty",
			Input: input{
				Security: SEC_WPA,
				SSID:     "",
				Password: "",
				Hidden:   false,
			},
		},
	}

	for _, fixture := range fixtures {
		t.Log(fixture.Name)

		// act
		_, err := New(
			fixture.Input.Security,
			fixture.Input.SSID,
			fixture.Input.Password,
			fixture.Input.Hidden,
		)

		// assert
		assert.Errorf(err, fixture.ExpectedErrorMessage)
	}
}

func TestWifiStringEncodingValidCases(t *testing.T) {
	t.Parallel()

	// arrange
	assert := require.New(t)
	fixtures := []struct {
		Name           string
		ExpectedOutput string
		Input          input
	}{
		{
			Name:           "All optional fields left empty",
			ExpectedOutput: "WIFI:T:WPA;S:Impermanence;P:\"\";H:false;;",
			Input: input{
				Security: SEC_WPA,
				SSID:     "Impermanence",
				Password: "",
				Hidden:   false,
			},
		},
		{
			Name:           "All fields filled",
			ExpectedOutput: "WIFI:T:WEP;S:Impermanence;P:Foo;H:true;;",
			Input: input{
				Security: SEC_WEP,
				SSID:     "Impermanence",
				Password: "Foo",
				Hidden:   true,
			},
		},
		{
			Name:           "SSID and Password contain characters that can be interpreted as hex digits",
			ExpectedOutput: "WIFI:T:WPA;S:\"AEFF3200\";P:\"aeff3200\";H:false;;",
			Input: input{
				Security: SEC_WPA,
				SSID:     "AEFF3200",
				Password: "aeff3200",
				Hidden:   false,
			},
		},
		{
			Name:           "SSID and Password contain special characters",
			ExpectedOutput: "WIFI:T:nopass;S:\\\\The\\;Fall\\,Of\\\"Hearts\\:;P:\\:The\\\"Great\\,Cold\\;Distance\\\\;H:false;;",
			Input: input{
				Security: SEC_NONE,
				SSID:     "\\The;Fall,Of\"Hearts:",
				Password: ":The\"Great,Cold;Distance\\",
				Hidden:   false,
			},
		},
	}

	for _, fixture := range fixtures {
		t.Log(fixture.Name)

		// act
		w, err := New(
			fixture.Input.Security,
			fixture.Input.SSID,
			fixture.Input.Password,
			fixture.Input.Hidden,
		)

		// assert
		assert.NoError(err)
		assert.Equal(fixture.ExpectedOutput, w.String())
	}
}
