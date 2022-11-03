package senbay

import (
	"encoding/json"
	"fmt"
	"log"
)

type Handler interface {
	Handle(senbayDict map[string]interface{})
}

type HandlerFunc func(senbayDict map[string]interface{})

type handler struct {
	fn HandlerFunc
}

func NewHandler(fn HandlerFunc) Handler {
	return &handler{fn}
}

func (h *handler) Handle(senbayDict map[string]interface{}) {
	h.fn(senbayDict)
}

func ShowResult(senbayDict map[string]interface{}) {
	bytes, err := json.Marshal(senbayDict)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}
