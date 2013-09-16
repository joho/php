package main

import (
	"fmt"
	"github.com/joho/whoisjohnbarton.com/homepage"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type proHttpHandler struct {
	fsHandler             http.Handler
	exactMatchHandleFuncs map[string]func(http.ResponseWriter, *http.Request)
}

func (h *proHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handleFunc, handleFuncPresent := h.exactMatchHandleFuncs[r.URL.Path]
	if handleFuncPresent == true {
		handleFunc(w, r)
	} else if !strings.HasSuffix(r.URL.Path, "/") {
		h.fsHandler.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func (h *proHttpHandler) ExactMatchFunc(path string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	h.exactMatchHandleFuncs[path] = handlerFunc
}

func newProHttpHandler() (h *proHttpHandler) {
	h = new(proHttpHandler)
	h.exactMatchHandleFuncs = map[string]func(http.ResponseWriter, *http.Request){}
	h.fsHandler = http.FileServer(http.Dir("public"))
	return
}

func main() {
	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		portNumber = "5000"
	}
	listenOn := fmt.Sprintf(":%v", portNumber)

	handler := newProHttpHandler()
	handler.ExactMatchFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := homepage.New()
		t, err := template.ParseFiles("views/index.html")
		if err == nil {
			t.Execute(w, p)
		} else {
			fmt.Fprintf(w, "Error parsing %v error was:\n\n%v", r.URL.Path, err)
		}
	})

	log.Fatal(http.ListenAndServe(listenOn, handler))
}
