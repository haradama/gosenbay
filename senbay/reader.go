package senbay

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/kbinani/screenshot"
)

type SenbayReader struct {
	readerMode int
	videoInput string
	cameraInput int
	screenInput int
	captureArea image.Rectangle
}

func NewSenbayReader(readerMode int, videoInput string, cameraInput int, screenInput int) *SenbayReader {
	senbayReader := &SenbayReader{
		readerMode: readerMode,
		videoInput: videoInput,
		cameraInput: cameraInput,
		screenInput: screenInput,
	}
	return senbayReader
}

func (senbayReader SenbayReader) SetCaptureAre(captureArea image.Rectangle) {
	senbayReader.captureArea = captureArea
}

func (senbayReader SenbayReader) Start() {
	var cap *gocv.VideoCapture
	var err error
	switch senbayReader.readerMode {
		case 0:
			cap, err = gocv.VideoCaptureFile(senbayReader.videoInput)
			if err != nil {
				panic(err)
			}
		case 1:
			cap, err = gocv.VideoCaptureDevice(senbayReader.cameraInput)
			if err != nil {
				panic(err)
			}
		case 2:
			bounds := screenshot.GetDisplayBounds(senbayReader.screenInput)
			_, err := screenshot.CaptureRect(bounds)
			if err != nil {
				panic(err)
			}
		default:
			msg := "error: The mode value should be taken 0(=video), 1(=camera), or 2(=screen)."
			panic(msg)
	}

	PN := 121
	senbayFrame := NewSenbayFrame(PN)
	qrReader := qrcode.NewQRCodeReader()
	mat := gocv.NewMat()
	if senbayReader.readerMode <= 1 {
		window := gocv.NewWindow("Senbay")
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
				senbayDict := senbayFrame.Decode(result.GetText())
				fmt.Println(senbayDict)
			}

			window.WaitKey(1)
		}
	} else if senbayReader.readerMode == 2 {
		for {
			fmt.Println("hello")
		}
	}
}