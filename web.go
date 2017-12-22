package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/php/homepage"
	"github.com/joho/prohttphandler"
)

func main() {
	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		portNumber = "5000"
	}
	listenOn := fmt.Sprintf(":%v", portNumber)

	handler := prohttphandler.New("public")
	handler.ExactMatchFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := homepage.New()
		t, err := template.ParseFiles("views/index.html")
		if err == nil {
			t.Execute(w, p)
		} else {
			fmt.Fprintf(w, "Error parsing %v error was:\n\n%v", r.URL.Path, err)
		}
	})

	log.Printf("Listening on %v", listenOn)
	log.Fatal(http.ListenAndServe(listenOn, handler))
}
