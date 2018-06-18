package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/joho/php/pages"
	"github.com/joho/prohttphandler"
)

func main() {
	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		portNumber = "5000"
	}
	listenOn := fmt.Sprintf(":%v", portNumber)

	handler := prohttphandler.New("public")

	home := pages.Homepage()
	handler.ExactMatchFunc(home.Path, Render(home))
	handler.ExactMatchFunc("/feed.xml", Feed(home))

	for _, post := range pages.AllPosts() {
		handler.ExactMatchFunc(post.Path, Render(post))
	}

	withRedirects := LegacyRedirectMiddleware(handler)

	log.Printf("Listening on %v", listenOn)
	log.Fatal(http.ListenAndServe(listenOn, withRedirects))
}

var securityHeaders = map[string]string{
	"Content-Security-Policy":   `upgrade-insecure-requests`,
	"Strict-Transport-Security": "max-age=2592000",
	"X-Xss-Protection":          "1; mode=block",
	"X-Frame-Options":           "DENY",
	// "X-Content-Type-Options":    "nosniff",
	"Referrer-Policy": "strict-origin-when-cross-origin",
}

func setSecurityHeaders(w http.ResponseWriter) {
	if os.Getenv("PRODUCTION") == "true" {
		for key, value := range securityHeaders {
			w.Header().Set(key, value)
		}
	}
}

func Render(page *pages.Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := page.Content()
		setSecurityHeaders(w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		t, err := template.ParseFiles("views/page.html")
		if err == nil {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			t.Execute(w, page)
		} else {
			fmt.Fprintf(w, "Error parsing %v error was:\n\n%v", r.URL.Path, err)
		}
	}
}

func Feed(page *pages.Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("views/feed.xml")
		setSecurityHeaders(w)
		if err == nil {
			w.Header().Set("Content-Type", "text/xml; charset=utf-8")
			t.Execute(w, page)
		} else {
			fmt.Fprintf(w, "Error parsing %v error was:\n\n%v", r.URL.Path, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
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
