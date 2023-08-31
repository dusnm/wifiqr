# wifiqr

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
![semver](https://img.shields.io/badge/semver-1.1.0-blue)
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
  -h, --help              help for wifiqr
  -i, --invisible         The visibility of your WiFi network. Set this to true if your WiFi is hidden.
  -n, --name string       The name of your WiFi network.
  -o, --output string     The output filename. Default is in the current directory. (default "/home/dusan/qr.png")
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

## Building from source
Ensure that you have a recent version of go installed

Clone the repository
```shell
git clone https://github.com/dusnm/wifiqr.git && cd wifiqr
```

Build
```shell
go build -o wifiqr .
```

## Licensing
Licensed under the terms of the GNU General Public License, version 3
