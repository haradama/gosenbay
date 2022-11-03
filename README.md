# gosenbay

gosenbay is a golang implementation of senbay that is intended to work on multiple platforms.

[![GoDoc](https://godoc.org/github.com/haradama/gosenbay?status.svg)](https://godoc.org/github.com/haradama/gosenbay)
[![Go Report Card](https://goreportcard.com/badge/github.com/haradama/gosenbay)](https://goreportcard.com/report/github.com/haradama/gosenbay)

## Usage

In this example, int and string value are converted to senbay format.

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
