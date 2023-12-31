# wifiqr

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
![semver](https://img.shields.io/badge/semver-1.2.2-blue)
![Tests](https://github.com/dusnm/jmbg/actions/workflows/test.yml/badge.svg?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/dusnm/wifiqr)](https://goreportcard.com/report/github.com/dusnm/wifiqr)

A simple program to generate a QR code to connect to a WiFi network

<img src="https://github.com/dusnm/wifiqr/blob/main/res/example.png?raw=true" height="320" width="300" alt="Example QR code">

## Usage
To get a list of all the available flags run `wifiqr --help` or `wifiqr -h`
```
Copyright (C) 2023 Dušan Mitrović <dusan@dusanmitrovic.xyz>
Licensed under the terms of the GNU GPL v3 only

Usage:
  wifiqr [flags]

Flags:
  -c, --color string      The color of the text in the frame. Accepts 3 or 6 hex characters. (default "8178e4")
  -h, --help              help for wifiqr
  -i, --invisible         The visibility of your WiFi network. Use this switch if your WiFi is hidden.
  -l, --logo string       The path to a raster image logo that will be put in the center of the QR code. Max dimensions: 140x140.
  -n, --name string       The name of your WiFi network.
  -x, --no-logo           Use this switch if you don't want a logo in the center. Takes precedence over the -l and --logo options.
  -o, --output string     The output filename. Default is in the current directory. (default "/home/dusan/Projects/personal/wifiqr/qr.png")
  -p, --password string   The password of your wifi network.
  -s, --security string   The security of your WiFi network. Can be one of nopass, WEP, WPA. (default "nopass")
  -v, --version           version for wifiqr
```

### Example
If you wanted to encode a QR code with the following information
```
SSID: Test_Network
Encryption: WPA3
Password: test
```

```shell
wifiqr -n Test_Network -s WPA -p test -o ./your_file_name.png
```

Or with long options
```shell
wifiqr --name Test_Network --security WPA -p test --output ./your_file_name.png
```

## Installation
You can use the precompiled binaries for your operating system and cpu architecture in the [release](https://github.com/dusnm/wifiqr/releases/latest) section.
Make sure you put the binary somewhere in your `$PATH`

For those using Arch Linux, a PKGBUILD is available in the AUR. Use either `makepkg` directly or your favorite AUR helper.

### Example
* makepkg
```shell
git clone https://aur.archlinux.org/wifiqr.git && cd wifiqr
makepkg -si
```
* yay
```shell
yay -S wifiqr
```

## Building from source
Ensure that you have a recent version of go installed (>=1.21)

Clone the repository
```shell
git clone https://github.com/dusnm/wifiqr.git && cd wifiqr
```

Build
```shell
CGO_ENABLED=0 go build -a -ldflags '-X "github.com/dusnm/wifiqr/cmd.version=1.2.2" -extldflags "-static"' -o ./bin/wifiqr .
```

(Optional) You can also use the `Taskfile.yml` if you happen to have `go-task` installed.
#### Linux
```shell
task build-linux
```

#### MacOS
```shell
task build-macos
```

#### Windows
```shell
task build-windows
```

#### All platforms
```shell
task build
```

## Licensing
Licensed under the terms of the GNU General Public License, version 3
