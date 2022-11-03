package senbay

import (
	"log"
	"strconv"

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
	mode       InputMode
	videoInput string
	nographic  bool
}

func NewSenbayReader(videoInput string, nographic bool) *Reader {
	senbayReader := &Reader{
		videoInput: videoInput,
		nographic:  nographic,
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

	cap, err = gocv.VideoCaptureFile(reader.videoInput)
	if err != nil {
		log.Fatal(err)
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
}
