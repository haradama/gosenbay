# gosenbay

## About
Gosenbay is an implementation of Senbay in golang. It features single binary and multi-platforms.

gosenbay tested on
- macOS Catalina 10.15.2
- Ubuntu 18.04 LTS
- Windows 10

## Requirements
- go
- opencv
- github.com/kbinani/screenshot
- github.com/makiuchi-d/gozxing
- github.com/spf13/cobra
- gocv.io/x/gocv
- golang.org/x/xerrors

## Installation

### OpenCV
If you have not installed opencv yet, please follow the instructions below.

#### Mac
```
$ brew install opencv
```

#### Ubuntu
```
$ export OPENCV_VERSION="4.3.0"
$ wget https://github.com/opencv/opencv/archive/${OPENCV_VERSION}.zip
$ unzip ${OPENCV_VERSION}.zip -d .
$ mkdir ./opencv-${OPENCV_VERSION}/build
$ cd ./opencv-${OPENCV_VERSION}/build/
$ cmake -DOPENCV_GENERATE_PKGCONFIG=ON -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=OFF -D WITH_FFMPEG=ON -D WITH_TBB=ON ..
$ make -j "$(nproc)"
$ make install
```

#### Windows
```
$ go get -u -d gocv.io/x/gocv
$ chdir %GOPATH%\src\gocv.io\x\gocv
$ win_build_opencv.cmd
```

### Binary

https://github.com/haradama/gosenbay/releases

```
$ wget https://github.com/haradama/gosenbay/releases/download/v0.1/gosenbay_0.1_darwin_amd64.tar.gz
$ tar -zxvf gosenbay_0.1_darwin_amd64.tar.gz
```

### Go build
```
$ git clone https://github.com/haradama/gosenbay
$ cd ./gosenbay
$ go build
```

## Usage
```
$ gosenbay -h

Usage:
  gosenbay [flags]
  gosenbay [command]

Available Commands:
  help        Help about any command
  read        Reader to recognize the sensor data embedded in the video
  version     Print the version number of gosenbay

Flags:
  -h, --help   help for gosenbay

Use "gosenbay [command] --help" for more information about a command.
```

```
$ gosenbay read -h
Reader to recognize the sensor data embedded in the video

Usage:
  gosenbay read [flags]

Flags:
  -h, --help            help for read
  -i, --infile string   input file path
  -m, --mode int        senbay reader mode (required)
                        0: video 1: camera 2: screenshot
  -n, --nographic       disable preview
```

### Example
```
$ ./gosenbay read -i video_path -m 0
```

![](./assets/demo.gif)

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
Gosenbay is available under the Apache License, Version 2.0 license. See the LICENSE file for more info.