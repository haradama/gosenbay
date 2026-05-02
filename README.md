# gosenbay

gosenbay is a Go implementation of Senbay that is intended to work on multiple platforms.

Senbay is a format and platform concept for capturing, embedding, integrating, and restreaming synchronized sensor-data streams. This repository provides a Go package for encoding and decoding Senbay data, and also includes a browser-based camera application powered by WebAssembly.

[![GoDoc](https://godoc.org/github.com/haradama/gosenbay?status.svg)](https://godoc.org/github.com/haradama/gosenbay)
[![Go Report Card](https://goreportcard.com/badge/github.com/haradama/gosenbay)](https://goreportcard.com/report/github.com/haradama/gosenbay)

## Web Camera App

A browser-based Senbay camera app is available here:

https://haradama.github.io/gosenbay/

The web app allows you to:

- Open the camera from a browser
- Collect sensor data such as time, location, acceleration, and device orientation
- Encode sensor data into Senbay format using Go WebAssembly
- Embed the encoded Senbay data as a QR code on top of the camera video
- Move the QR code position by dragging with a mouse or finger
- Record the camera video with the embedded Senbay QR code
- Use the app on smartphones and PCs through the browser

The web app is designed so that the final recorded video contains only the camera image and the Senbay QR code. UI elements such as recording indicators and drag hints are displayed in the browser UI and are not burned into the recorded video.

### Browser Permissions

The web app may request the following permissions depending on the device and browser:

- Camera access
- Location access
- Motion sensor access
- Orientation sensor access

For best results on smartphones, open the app over HTTPS. GitHub Pages already serves the app over HTTPS.

## Senbay Data Format

gosenbay supports encoding and decoding key-value sensor data.

Common reserved Senbay keys include:

| Key | Meaning |
| --- | --- |
| `TIME` | Timestamp |
| `LONG` | Longitude |
| `LATI` | Latitude |
| `ALTI` | Altitude |
| `ACCX` | Acceleration X |
| `ACCY` | Acceleration Y |
| `ACCZ` | Acceleration Z |
| `YAW` | Yaw / alpha |
| `ROLL` | Roll / gamma |
| `PITC` | Pitch / beta |
| `HEAD` | Heading |
| `SPEE` | Speed |
| `BRIG` | Brightness |
| `AIRP` | Air pressure |
| `HTBT` | Heartbeat |

The default positional notation used by this implementation is `PN = 121`.

## Usage as a Go Library

Install or import the package in your Go project:

```go
import "github.com/haradama/gosenbay/senbay"
```

### Encode Data

In this example, integer and string values are converted to Senbay format.

```go
package main

import (
	"fmt"

	"github.com/haradama/gosenbay/senbay"
)

func main() {
	PN := 121
	SD, _ := senbay.NewSenbayData(PN)

	value1 := 123
	SD.AddInt("KEY1", value1)
	SD.AddText("KEY2", "value2")

	result := SD.Encode(true)
	fmt.Println(result)
}
```

### Encode Sensor-like Data

```go
package main

import (
	"fmt"
	"time"

	"github.com/haradama/gosenbay/senbay"
)

func main() {
	SD, _ := senbay.NewSenbayData(121)

	SD.AddInt64("TIME", time.Now().UnixMilli())
	SD.AddFloat64("LATI", 35.681236)
	SD.AddFloat64("LONG", 139.767125)
	SD.AddFloat64("ACCX", 0.12)
	SD.AddFloat64("ACCY", -0.03)
	SD.AddFloat64("ACCZ", 9.81)

	encoded := SD.Encode(true)

	fmt.Println(encoded)
}
```

### Decode Data

```go
package main

import (
	"fmt"

	"github.com/haradama/gosenbay/senbay"
)

func main() {
	SD, _ := senbay.NewSenbayData(121)

	encoded := "V:4,KEY1:\u0002\u0003,KEY2:'value2'"
	decoded := SD.Decode(encoded)

	fmt.Println(decoded)
}
```

## Native GoCV Reader / Writer

The repository also includes native reader and writer components that use GoCV.

These native components are intended for desktop or server environments where OpenCV is available. They are excluded from the WebAssembly build because browser camera access, drawing, and recording are handled by Web APIs instead.

On Ubuntu, install OpenCV dependencies before building the native GoCV parts:

```bash
sudo apt update
sudo apt install -y libopencv-dev pkg-config
```

Then build and test:

```bash
go build -v ./...
go test -v ./...
```

## WebAssembly Architecture

The browser app uses a hybrid architecture:

```text
Browser App
├── Camera / Canvas / Recording: TypeScript + Web APIs
├── QR generation: TypeScript
├── Sensor collection: TypeScript + Browser APIs
└── Senbay encode/decode: Go compiled to WebAssembly
```

The Go Senbay core is compiled to WebAssembly and exposed to JavaScript as:

- `senbayEncode(payloadJson, compress)`
- `senbayDecode(senbayText)`

This keeps the Senbay format implementation shared with the Go package while letting the browser handle camera, sensor, touch, and recording features.

## Development

### Build the WebAssembly Module

```bash
./scripts/build-wasm.sh
```

This generates:

```text
web/public/senbay.wasm
web/public/wasm_exec.js
```

These files are generated artifacts and do not need to be committed.

### Run the Web App Locally

```bash
cd web
npm install
npm run dev
```

Then open the local Vite URL in your browser.

When testing camera and sensor features on a smartphone, use HTTPS or a browser/device setup that allows camera access from your development server.

### Build the Web App

```bash
./scripts/build-wasm.sh

cd web
npm ci
npm run build
```

The built static files are generated in:

```text
web/dist
```

## GitHub Pages Deployment

The web app is deployed to GitHub Pages:

https://haradama.github.io/gosenbay/

The deployment workflow builds the Go WebAssembly module in CI, installs the web dependencies, builds the Vite app, and publishes `web/dist` to GitHub Pages.

The generated WebAssembly files are created during CI:

```text
web/public/senbay.wasm
web/public/wasm_exec.js
```

They are not required to be stored in the repository.

### Related Links
- [tetujin/SenbayKit-CLI](https://github.com/tetujin/SenbayKit-CLI) (The original)
- [Senbay Platform Website](http://www.senbay.info)
- [Senbay YouTube Channel](https://www.youtube.com/channel/UCbnQUEc3KpE1M9auxwMh2dA/videos)

### Reference

```
@inproceedings{Nishiyama:2018:SPI:3236112.3236154,
    author = {Nishiyama, Yuuki and Dey, Anind K. and Ferreira, Denzil and Yonezawa, Takuro and Nakazawa, Jin},
    title = {Senbay: A Platform for Instantly Capturing, Integrating, and Restreaming of Synchronized Multiple Sensor-data Stream},
    booktitle = {Proceedings of the 20th International Conference on Human-Computer Interaction with Mobile Devices and Services Adjunct},
    series = {MobileHCI '18},
    year = {2018},
    location = {Barcelona, Spain},
    publisher = {ACM},
} 
```

### License
gosenbay is available under the Apache License, Version 2.0 license. See the LICENSE file for more info.
