package prohttphandler

import (
	"net/http"
	"strings"
)

type ProHttpHandler struct {
	fsHandler             http.Handler
	exactMatchHandleFuncs map[string]func(http.ResponseWriter, *http.Request)
}

func (h *ProHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handleFunc, handleFuncPresent := h.exactMatchHandleFuncs[r.URL.Path]
	if handleFuncPresent == true {
		handleFunc(w, r)
	} else if !strings.HasSuffix(r.URL.Path, "/") {
		h.fsHandler.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func (h *ProHttpHandler) ExactMatchFunc(path string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	h.exactMatchHandleFuncs[path] = handlerFunc
}

func New() (h *ProHttpHandler) {
	h = new(ProHttpHandler)
	h.exactMatchHandleFuncs = map[string]func(http.ResponseWriter, *http.Request){}
	h.fsHandler = http.FileServer(http.Dir("public"))
	return
}
