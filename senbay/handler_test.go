package senbay

import (
	"reflect"
	"testing"
)

func TestNewHandler(t *testing.T) {
	type args struct {
		fn HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handler_Handle(t *testing.T) {
	type fields struct {
		fn HandlerFunc
	}
	type args struct {
		senbayDict map[string]interface{}
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
			h := &handler{
				fn: tt.fields.fn,
			}
			h.Handle(tt.args.senbayDict)
		})
	}
}

func TestShowResult(t *testing.T) {
	type args struct {
		senbayDict map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShowResult(tt.args.senbayDict)
		})
	}
}
