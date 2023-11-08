package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/dusnm/wifiqr/pkg/qr"
	"github.com/dusnm/wifiqr/pkg/utils"
	"github.com/dusnm/wifiqr/pkg/wifi"
	"github.com/spf13/cobra"
)

var (
	security   string
	ssid       string
	password   string
	hidden     bool
	output     string
	titleColor string
	logoPath   string
	noLogo     bool
	version    string
)

var rootCmd = &cobra.Command{
	Use:     "wifiqr",
	Version: version,
	Short:   "This program helps you generate QR codes to connect to WiFi networks",
	Long: `Copyright (C) 2023 Dušan Mitrović <dusan@dusanmitrovic.xyz>
Licensed under the terms of the GNU GPL v3 only
    `,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := utils.ParseHexColor(titleColor)
		if err != nil {
			return err
		}

		wf, err := wifi.New(
			security,
			ssid,
			password,
			hidden,
		)
		if err != nil {
			return err
		}

		if !filepath.IsAbs(output) {
			output, err = filepath.Abs(output)
			if err != nil {
				return err
			}
		}

		if !noLogo && logoPath != "" && !filepath.IsAbs(logoPath) {
			logoPath, err = filepath.Abs(logoPath)
			if err != nil {
				return err
			}
		}

		f, err := os.OpenFile(output, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0o644)
		if err != nil {
			return err
		}

		defer f.Close()

		b, err := qr.Generate(wf, noLogo, logoPath)
		if err != nil {
			_ = os.Remove(output)

			return err
		}

		// removes the escape characters from the generated image
		replacer := strings.NewReplacer(
			"\\", "",
			"\"", "",
		)

		if err = qr.AddHeader(replacer.Replace(wf.SSID), b, f, c); err != nil {
			_ = os.Remove(output)

			return err
		}

		return nil
	},
}

func Execute() {
	rootCmd.Flags().StringVarP(
		&security,
		"security",
		"s",
		wifi.SEC_NONE,
		"The security of your WiFi network. Can be one of nopass, WEP, WPA.",
	)

	rootCmd.Flags().StringVarP(
		&ssid,
		"name",
		"n",
		"",
		"The name of your WiFi network.",
	)

	rootCmd.Flags().StringVarP(
		&password,
		"password",
		"p",
		"",
		"The password of your wifi network.",
	)

	rootCmd.Flags().BoolVarP(
		&hidden,
		"invisible",
		"i",
		false,
		"The visibility of your WiFi network. Use this switch if your WiFi is hidden.",
	)

	rootCmd.Flags().StringVarP(
		&output,
		"output",
		"o",
		func() string {
			d, _ := os.Getwd()
			return filepath.Join(d, "qr.png")
		}(),
		"The output filename. Default is in the current directory.",
	)

	rootCmd.Flags().StringVarP(
		&titleColor,
		"color",
		"c",
		"8178e4",
		"The color of the text in the frame. Accepts 3 or 6 hex characters.",
	)

	rootCmd.Flags().StringVarP(
		&logoPath,
		"logo",
		"l",
		"",
		"The path to a raster image logo that will be put in the center of the QR code. Max dimensions: 140x140.",
	)

	rootCmd.Flags().BoolVarP(
		&noLogo,
		"no-logo",
		"x",
		false,
		"Use this switch if you don't want a logo in the center. Takes precedence over the -l and --logo options.",
	)

	rootCmd.MarkFlagRequired("name")

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
