package senbay

import (
	"encoding/json"
	"fmt"
	"log"
)

// Handler is an interface that defines the method to handle the Senbay data.
type Handler interface {
	Handle(senbayDict map[string]interface{})
}

// HandlerFunc is a type that defines the function to handle the Senbay data.
type HandlerFunc func(senbayDict map[string]interface{})

// handler is a struct that implements the Handler interface.
type handler struct {
	fn HandlerFunc
}

// NewHandler returns a new Handler that wraps the given HandlerFunc.
func NewHandler(fn HandlerFunc) Handler {
	return &handler{fn}
}

// Handle calls the HandlerFunc to handle the Senbay data.
func (h *handler) Handle(senbayDict map[string]interface{}) {
	h.fn(senbayDict)
}

// ShowResult is a function that shows the result of the Senbay data in JSON format.
func ShowResult(senbayDict map[string]interface{}) {
	bytes, err := json.Marshal(senbayDict)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}
