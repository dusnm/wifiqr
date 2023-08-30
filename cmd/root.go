package cmd

import (
	"os"
	"path/filepath"

	"github.com/dusnm/wifiqr/pkg/qr"
	"github.com/dusnm/wifiqr/pkg/wifi"
	"github.com/spf13/cobra"
)

var security string
var ssid string
var password string
var hidden bool
var output string

var rootCmd = &cobra.Command{
	Use:     "wifiqr",
	Version: "1.0.0",
	Short:   "This program helps you generate QR codes to connect to WiFi networks",
	Long: `Copyright (C) 2023 Dušan Mitrović <dusan@dusanmitrovic.xyz>
Licensed under the terms of the GNU GPL v3 only
    `,
	RunE: func(cmd *cobra.Command, args []string) error {
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

		f, err := os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		if err != nil {
			return err
		}

		err = qr.Generate(wf, f)

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
		"The visibility of your WiFi network. Set this to true if your WiFi is hidden.",
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

	rootCmd.MarkFlagRequired("name")

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
