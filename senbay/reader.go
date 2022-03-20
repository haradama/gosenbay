package senbay

import (
	"encoding/json"
	"fmt"
	"image"
	"log"
	"strconv"

	"github.com/kbinani/screenshot"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"gocv.io/x/gocv"
)

type InputMode int

const (
	modeVideoInput InputMode = iota
	modeCameraInput
	modeScreenInput
)

// Reader describe an reader to interpret senbay style contents
type Reader struct {
	mode        InputMode
	videoInput  string
	cameraInput int
	screenInput int
	captureArea image.Rectangle
	nographic   bool
}

func NewSenbayReader(mode InputMode, videoInput string, cameraInput int, screenInput int, nographic bool) *Reader {
	senbayReader := &Reader{
		mode:        mode,
		videoInput:  videoInput,
		cameraInput: cameraInput,
		screenInput: screenInput,
		nographic:   nographic,
	}
	return senbayReader
}

// NewSenbayVideoReader returns a new SenbayReader for a video
func NewSenbayVideoReader(videoInput string, nographic bool) *Reader {
	mode := modeVideoInput
	senbayReader := &Reader{
		mode:       mode,
		videoInput: videoInput,
		nographic:  nographic,
	}
	return senbayReader
}

// NewSenbayVideoReader returns a new SenbayReader for a camera
func NewSenbayCameraReader(cameraInput int, nographic bool) *Reader {
	mode := modeCameraInput
	senbayReader := &Reader{
		mode:        mode,
		cameraInput: cameraInput,
		nographic:   nographic,
	}
	return senbayReader
}

// NewSenbayVideoReader returns a new SenbayReader for a screen
func NewSenbayScreenReader(screenInput int, nographic bool) *Reader {
	mode := modeCameraInput
	senbayReader := &Reader{
		mode:        mode,
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
func (reader Reader) Start(fn HandlerFunc) {
	handler := NewHandler(fn)
	PN := 121
	SenbayData, err := NewSenbayData(PN)
	if err != nil {
		log.Fatal(err)
	}
	qrReader := qrcode.NewQRCodeReader()
	mat := gocv.NewMat()
	var cap *gocv.VideoCapture
	if reader.mode == modeVideoInput || reader.mode == modeCameraInput {
		if reader.mode == modeVideoInput {
			cap, err = gocv.VideoCaptureFile(reader.videoInput)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			cap, err = gocv.VideoCaptureDevice(reader.cameraInput)
			if err != nil {
				log.Fatal(err)
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
				log.Fatal(err)
			}
			bmp, err := gozxing.NewBinaryBitmapFromImage(img)
			if err != nil {
				log.Fatal(err)
			}
			result, err := qrReader.Decode(bmp, nil)
			if err == nil {
				senbayDict := map[string]interface{}{}
				for key, value := range SenbayData.Decode(result.GetText()) {
					parsedvalue, err := strconv.ParseFloat(value, 64)
					if err == nil {
						senbayDict[key] = parsedvalue
						continue
					} else {
						senbayDict[key] = value
					}
				}
				handler.Handle(senbayDict)
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
			log.Fatal(err)
		}

		for {
			bounds := screenshot.GetDisplayBounds(reader.screenInput)
			img, err := screenshot.CaptureRect(bounds)
			if err != nil {
				log.Fatal(err)
			}
			bmp, err := gozxing.NewBinaryBitmapFromImage(img)
			if err != nil {
				log.Fatal(err)
			}
			result, err := qrReader.Decode(bmp, nil)
			if err == nil {
				senbayDict := map[string]interface{}{}
				for key, value := range SenbayData.Decode(result.GetText()) {
					parsedvalue, err := strconv.ParseFloat(value, 64)
					if err == nil {
						senbayDict[key] = parsedvalue
						continue
					} else {
						senbayDict[key] = value
					}
				}
				handler.Handle(senbayDict)
			}
		}
	} else {
		msg := "error: The mode value should be taken 0(=video), 1(=camera), or 2(=screen)."
		log.Fatal(msg)
	}
}

func ShowResult(senbayDict map[string]interface{}) {
	bytes, err := json.Marshal(senbayDict)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}
