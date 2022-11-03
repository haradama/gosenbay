// Package senbay provides the functions to encode and
// decode to the senbay format.
package senbay

import (
	"reflect"
	"testing"
)

func TestNewBaseX(t *testing.T) {
	type args struct {
		positionalNotation int
	}
	tests := []struct {
		name    string
		args    args
		want    *BaseX
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBaseX(tt.args.positionalNotation)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBaseX() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBaseX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseX_encodeLongValue(t *testing.T) {
	type fields struct {
		PN           int
		Table        []int
		ReverseTable []int
	}
	type args struct {
		lVal int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []rune
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseX := BaseX{
				PN:           tt.fields.PN,
				Table:        tt.fields.Table,
				ReverseTable: tt.fields.ReverseTable,
			}
			if got := baseX.encodeLongValue(tt.args.lVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseX.encodeLongValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseX_encodeDoubleValue(t *testing.T) {
	type fields struct {
		PN           int
		Table        []int
		ReverseTable []int
	}
	type args struct {
		dVal float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []rune
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseX := BaseX{
				PN:           tt.fields.PN,
				Table:        tt.fields.Table,
				ReverseTable: tt.fields.ReverseTable,
			}
			if got := baseX.encodeDoubleValue(tt.args.dVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseX.encodeDoubleValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseX_decodeLongValue(t *testing.T) {
	type fields struct {
		PN           int
		Table        []int
		ReverseTable []int
	}
	type args struct {
		sVal []rune
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseX := BaseX{
				PN:           tt.fields.PN,
				Table:        tt.fields.Table,
				ReverseTable: tt.fields.ReverseTable,
			}
			if got := baseX.decodeLongValue(tt.args.sVal); got != tt.want {
				t.Errorf("BaseX.decodeLongValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseX_decodeDoubleValue(t *testing.T) {
	type fields struct {
		PN           int
		Table        []int
		ReverseTable []int
	}
	type args struct {
		sVal []rune
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseX := BaseX{
				PN:           tt.fields.PN,
				Table:        tt.fields.Table,
				ReverseTable: tt.fields.ReverseTable,
			}
			if got := baseX.decodeDoubleValue(tt.args.sVal); got != tt.want {
				t.Errorf("BaseX.decodeDoubleValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSenbayFormat(t *testing.T) {
	type args struct {
		PN int
	}
	tests := []struct {
		name    string
		args    args
		want    *Format
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSenbayFormat(tt.args.PN)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSenbayFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSenbayFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormat_getReservedShortKey(t *testing.T) {
	type fields struct {
		ReversedKeys map[string]string
		PN           int
		baseX        *BaseX
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			senbayFormat := Format{
				ReversedKeys: tt.fields.ReversedKeys,
				PN:           tt.fields.PN,
				baseX:        tt.fields.baseX,
			}
			if got := senbayFormat.getReservedShortKey(tt.args.key); got != tt.want {
				t.Errorf("Format.getReservedShortKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormat_getReservedOriginalKey(t *testing.T) {
	type fields struct {
		ReversedKeys map[string]string
		PN           int
		baseX        *BaseX
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			senbayFormat := Format{
				ReversedKeys: tt.fields.ReversedKeys,
				PN:           tt.fields.PN,
				baseX:        tt.fields.baseX,
			}
			if got := senbayFormat.getReservedOriginalKey(tt.args.key); got != tt.want {
				t.Errorf("Format.getReservedOriginalKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormat_encode(t *testing.T) {
	type fields struct {
		ReversedKeys map[string]string
		PN           int
		baseX        *BaseX
	}
	type args struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			senbayFormat := Format{
				ReversedKeys: tt.fields.ReversedKeys,
				PN:           tt.fields.PN,
				baseX:        tt.fields.baseX,
			}
			if got := senbayFormat.encode(tt.args.text); got != tt.want {
				t.Errorf("Format.encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormat_decode(t *testing.T) {
	type fields struct {
		ReversedKeys map[string]string
		PN           int
		baseX        *BaseX
	}
	type args struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			senbayFormat := Format{
				ReversedKeys: tt.fields.ReversedKeys,
				PN:           tt.fields.PN,
				baseX:        tt.fields.baseX,
			}
			if got := senbayFormat.decode(tt.args.text); got != tt.want {
				t.Errorf("Format.decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSenbayData(t *testing.T) {
	type args struct {
		PN int
	}
	tests := []struct {
		name    string
		args    args
		want    *Data
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSenbayData(tt.args.PN)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSenbayData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSenbayData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_AddInt(t *testing.T) {
	type fields struct {
		senbayData map[string]string
		PN         int
		SF         *Format
	}
	type args struct {
		key   string
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SD := Data{
				senbayData: tt.fields.senbayData,
				PN:         tt.fields.PN,
				SF:         tt.fields.SF,
			}
			SD.AddInt(tt.args.key, tt.args.value)
		})
	}
}

func TestData_AddInt64(t *testing.T) {
	type fields struct {
		senbayData map[string]string
		PN         int
		SF         *Format
	}
	type args struct {
		key   string
		value int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SD := Data{
				senbayData: tt.fields.senbayData,
				PN:         tt.fields.PN,
				SF:         tt.fields.SF,
			}
			SD.AddInt64(tt.args.key, tt.args.value)
		})
	}
}

func TestData_AddFloat(t *testing.T) {
	type fields struct {
		senbayData map[string]string
		PN         int
		SF         *Format
	}
	type args struct {
		key   string
		value float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SD := Data{
				senbayData: tt.fields.senbayData,
				PN:         tt.fields.PN,
				SF:         tt.fields.SF,
			}
			SD.AddFloat(tt.args.key, tt.args.value)
		})
	}
}

func TestData_AddFloat64(t *testing.T) {
	type fields struct {
		senbayData map[string]string
		PN         int
		SF         *Format
	}
	type args struct {
		key   string
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SD := Data{
				senbayData: tt.fields.senbayData,
				PN:         tt.fields.PN,
				SF:         tt.fields.SF,
			}
			SD.AddFloat64(tt.args.key, tt.args.value)
		})
	}
}

func TestData_AddText(t *testing.T) {
	type fields struct {
		senbayData map[string]string
		PN         int
		SF         *Format
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SD := Data{
				senbayData: tt.fields.senbayData,
				PN:         tt.fields.PN,
				SF:         tt.fields.SF,
			}
			SD.AddText(tt.args.key, tt.args.value)
		})
	}
}

func TestData_Clear(t *testing.T) {
	type fields struct {
		senbayData map[string]string
		PN         int
		SF         *Format
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SD := Data{
				senbayData: tt.fields.senbayData,
				PN:         tt.fields.PN,
				SF:         tt.fields.SF,
			}
			SD.Clear()
		})
	}
}

func TestData_Encode(t *testing.T) {
	type fields struct {
		senbayData map[string]string
		PN         int
		SF         *Format
	}
	type args struct {
		compress bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SD := Data{
				senbayData: tt.fields.senbayData,
				PN:         tt.fields.PN,
				SF:         tt.fields.SF,
			}
			if got := SD.Encode(tt.args.compress); got != tt.want {
				t.Errorf("Data.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Decode(t *testing.T) {
	type fields struct {
		senbayData map[string]string
		PN         int
		SF         *Format
	}
	type args struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SD := Data{
				senbayData: tt.fields.senbayData,
				PN:         tt.fields.PN,
				SF:         tt.fields.SF,
			}
			if got := SD.Decode(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
