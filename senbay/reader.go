package senbay

import (
	"encoding/json"
	"fmt"
	"image"

	"github.com/kbinani/screenshot"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"gocv.io/x/gocv"
)

const (
	modeVideoInput  = 0
	modeCameraInput = 1
	modeScreenInput = 2
)

// Reader describe an reader to interpret senbay style contents
type Reader struct {
	mode        int
	videoInput  string
	cameraInput int
	screenInput int
	captureArea image.Rectangle
	nographic   bool
}

// NewSenbayReader returns a new SenbayReader based on mode
func NewSenbayReader(mode int, videoInput string, cameraInput int, screenInput int, nographic bool) *Reader {
	senbayReader := &Reader{
		mode:        mode,
		videoInput:  videoInput,
		cameraInput: cameraInput,
		screenInput: screenInput,
		nographic:   nographic,
	}
	return senbayReader
}

// SetCaptureArea set capture area to senbay reader based on image.Rectangle
func (reader Reader) SetCaptureArea(captureArea image.Rectangle) {
	reader.captureArea = captureArea
}

// Start interpreting captured image recorded in senbay style
func (reader Reader) Start() {
	PN := 121
	SenbayData, err := NewSenbayData(PN)
	if err != nil {
		panic(err)
	}
	qrReader := qrcode.NewQRCodeReader()
	mat := gocv.NewMat()
	var cap *gocv.VideoCapture
	if reader.mode == modeVideoInput || reader.mode == modeCameraInput {
		if reader.mode == modeVideoInput {
			cap, err = gocv.VideoCaptureFile(reader.videoInput)
			if err != nil {
				panic(err)
			}
		} else {
			cap, err = gocv.VideoCaptureDevice(reader.cameraInput)
			if err != nil {
				panic(err)
			}
		}
		var title string
		var window *gocv.Window
		if !reader.nographic {
			title = "Senbay Reader"
			window = gocv.NewWindow(title)
		}
		for {
			cap.Read(&mat)
			img, err := mat.ToImage()
			if err != nil {
				panic(err)
			}
			bmp, err := gozxing.NewBinaryBitmapFromImage(img)
			if err != nil {
				panic(err)
			}
			result, err := qrReader.Decode(bmp, nil)
			if err == nil {
				senbayDict := SenbayData.Decode(result.GetText())
				bytes, err := json.Marshal(senbayDict)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(bytes))
			}
			if !reader.nographic {
				window.IMShow(mat)
				window.WaitKey(1)

				if window.WaitKey(1) == keyCodeEsc {
					break
				}
			}
		}
	} else if reader.mode == modeScreenInput {
		bounds := screenshot.GetDisplayBounds(reader.screenInput)
		_, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}

		for {
			bounds := screenshot.GetDisplayBounds(reader.screenInput)
			img, err := screenshot.CaptureRect(bounds)
			if err != nil {
				panic(err)
			}
			bmp, err := gozxing.NewBinaryBitmapFromImage(img)
			if err != nil {
				panic(err)
			}
			result, err := qrReader.Decode(bmp, nil)
			if err == nil {
				senbayDict := SenbayData.Decode(result.GetText())
				bytes, err := json.Marshal(senbayDict)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(bytes))
			}
		}
	} else {
		msg := "error: The mode value should be taken 0(=video), 1(=camera), or 2(=screen)."
		panic(msg)
	}
}
