package senbay

import (
	"reflect"
	"testing"
)

func TestNewSenbayFrame(t *testing.T) {
	type args struct {
		qrBoxSize         uint
		qrErrorCorrection int
	}
	tests := []struct {
		name string
		args args
		want *SenbayFrame
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSenbayFrame(tt.args.qrBoxSize, tt.args.qrErrorCorrection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSenbayFrame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSenbayWriter(t *testing.T) {
	type args struct {
		outfile      string
		width        uint
		height       uint
		qrBoxSize    uint
		cameraNumber uint
		codec        string
		fps          uint
	}
	tests := []struct {
		name string
		args args
		want *Writer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSenbayWriter(tt.args.outfile, tt.args.width, tt.args.height, tt.args.qrBoxSize, tt.args.cameraNumber, tt.args.codec, tt.args.fps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSenbayWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriter_Start(t *testing.T) {
	type fields struct {
		outfile      string
		width        uint
		height       uint
		qrBoxSize    uint
		fps          uint
		cameraNumber uint
		codec        string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := Writer{
				outfile:      tt.fields.outfile,
				width:        tt.fields.width,
				height:       tt.fields.height,
				qrBoxSize:    tt.fields.qrBoxSize,
				fps:          tt.fields.fps,
				cameraNumber: tt.fields.cameraNumber,
				codec:        tt.fields.codec,
			}
			writer.Start()
		})
	}
}
