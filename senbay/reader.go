package senbay

import (
	"fmt"
	"image"

	"github.com/kbinani/screenshot"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"gocv.io/x/gocv"
)

// Reader describe an reader to interpret senbay style contents
type Reader struct {
	mode        int
	videoInput  string
	cameraInput int
	screenInput int
	captureArea image.Rectangle
}

// NewSenbayReader returns a new SenbayReader based on mode
func NewSenbayReader(mode int, videoInput string, cameraInput int, screenInput int) *Reader {
	senbayReader := &Reader{
		mode:        mode,
		videoInput:  videoInput,
		cameraInput: cameraInput,
		screenInput: screenInput,
	}
	return senbayReader
}

// SetCaptureArea set capture area to senbay reader based on image.Rectangle
func (reader Reader) SetCaptureArea(captureArea image.Rectangle) {
	reader.captureArea = captureArea
}

// Start interpreting captured image recorded in senbay style
func (reader Reader) Start() {
	var cap *gocv.VideoCapture
	var err error
	switch reader.mode {
	case 0:
		cap, err = gocv.VideoCaptureFile(reader.videoInput)
		if err != nil {
			panic(err)
		}
	case 1:
		cap, err = gocv.VideoCaptureDevice(reader.cameraInput)
		if err != nil {
			panic(err)
		}
	case 2:
		bounds := screenshot.GetDisplayBounds(reader.screenInput)
		_, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
	default:
		msg := "error: The mode value should be taken 0(=video), 1(=camera), or 2(=screen)."
		panic(msg)
	}

	PN := 121
	SenbayData, err := NewSenbayData(PN)
	if err != nil {
		panic(err)
	}
	qrReader := qrcode.NewQRCodeReader()
	mat := gocv.NewMat()
	if reader.mode <= 1 {
		title := "Senbay Reader"
		window := gocv.NewWindow(title)
		for {
			cap.Read(&mat)
			window.IMShow(mat)
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
				fmt.Println(senbayDict)
			}

			window.WaitKey(1)
		}
	} else if reader.mode == 2 {
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
				fmt.Println(senbayDict)
			}
		}
	}
}
