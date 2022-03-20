package senbay

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
