package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

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
	handler.ExactMatchFunc("/", Index)

	withRedirects := LegacyRedirectMiddleware(handler)

	log.Printf("Listening on %v", listenOn)
	log.Fatal(http.ListenAndServe(listenOn, withRedirects))
}

func Index(w http.ResponseWriter, r *http.Request) {
	p := homepage.New()
	t, err := template.ParseFiles("views/index.html")
	if err == nil {
		t.Execute(w, p)
	} else {
		fmt.Fprintf(w, "Error parsing %v error was:\n\n%v", r.URL.Path, err)
	}
}

func LegacyRedirectMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		devRequest := r.Host == "localhost:5000"
		legacyDomain := r.Host != "johnbarton.co"

		if !devRequest && legacyDomain {
			log.Printf("Redirecting %v dev:%v legacydomain:%v", r.URL, devRequest, legacyDomain)
			URL := *r.URL
			URL.Scheme = "https"
			URL.Host = "johnbarton.co"
			URL.Path = strings.ToLower(URL.Path)

			http.Redirect(w, r, URL.String(), http.StatusPermanentRedirect)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
