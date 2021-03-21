package senbay

import (
	"time"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"gocv.io/x/gocv"
)

const keyCodeEsc = 27

type SenbayFrame struct {
	qrBoxSize         uint
	qrErrorCorrection int
}

func NewSenbayFrame(qrBoxSize uint, qrErrorCorrection int) *SenbayFrame {
	senbayFrame := &SenbayFrame{
		qrBoxSize:         qrBoxSize,
		qrErrorCorrection: qrErrorCorrection,
	}

	return senbayFrame
}

type Writer struct {
	outfile      string
	width        uint
	height       uint
	qrBoxSize    uint
	fps          uint
	cameraNumber uint
	codec        string
}

func NewSenbayWriter(outfile string, width uint, height uint, qrBoxSize uint, cameraNumber uint, codec string, fps uint) *Writer {
	writer := &Writer{
		outfile:      outfile,
		width:        width,
		height:       height,
		qrBoxSize:    qrBoxSize,
		cameraNumber: cameraNumber,
		codec:        codec,
		fps:          fps,
	}

	return writer
}

func (writer Writer) Start() {
	PN := 121

	title := "gosenbay"
	window := gocv.NewWindow(title)

	webcam, err := gocv.OpenVideoCapture(writer.cameraNumber)
	if err != nil {
		panic(err)
	}
	defer webcam.Close()

	img := gocv.NewMat()
	if ok := webcam.Read(&img); !ok {
		panic("Cannot read device")
	}

	videoCodec := writer.codec
	fps := float64(writer.fps)

	videoWriter, err := gocv.VideoWriterFile(
		writer.outfile, videoCodec, fps, img.Cols(), img.Rows(), true)
	if err != nil {
		panic(err)
	}
	defer videoWriter.Close()

	qrWriter := qrcode.NewQRCodeWriter()

	for {
		SD, err := NewSenbayData(PN)
		if err != nil {
			panic(err)
		}

		currentTime := time.Now().UnixNano() / int64(time.Millisecond)
		SD.AddInt64("TIME", currentTime)
		encodedText := SD.Encode(false)

		qrWidth := int(writer.qrBoxSize)
		qrHeight := int(writer.qrBoxSize)

		qrCode, err := qrWriter.EncodeWithoutHint(encodedText, gozxing.BarcodeFormat_QR_CODE, qrWidth, qrHeight)
		if err != nil {
			panic(err)
		}

		if img.Empty() {
			continue
		}
		webcam.Read(&img)

		ch := img.Channels()
		width := qrCode.GetWidth()
		height := qrCode.GetHeight()
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				var color uint8
				if !qrCode.Get(x, y) {
					color = 255 // White
				}
				for c := 0; c < ch; c++ {
					img.SetUCharAt(x, y*ch+c, color)
				}
			}
		}

		videoWriter.Write(img)

		window.IMShow(img)
		if window.WaitKey(1) == keyCodeEsc {
			break
		}
	}
}
