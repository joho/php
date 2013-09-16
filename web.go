package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
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

type link struct {
	Title string
	Url   string
}

type quote struct {
	Text   string
	Author string
	Url    string
}

type homePage struct {
	Writings []link
	Projects []link
	Social   []link
	quotes   []quote
}

func (p homePage) Quote() quote {
	quoteIndex := rand.Intn(len(p.quotes))
	return p.quotes[quoteIndex]
}

func newHomePage() *homePage {
	p := new(homePage)
	// writings
	p.Writings = []link{
		link{"Web Scale for the Rest of Us (RubyConfAU Talk)", "http://vimeo.com/61342269"},
		link{"Interview for The Fetch", "http://blog.thefetch.com/2012/10/29/interview-melbourne-local-john-barton/"},
		link{"The Minimum Viable Rails Stack", "http://goodfil.ms/blog/posts/2012/09/15/minimum-viable-rails-stack/"},
		link{"Will App.net be the connective tissue founders & developers can rely on?", "http://thenextweb.com/apps/2012/08/18/why-10000-people-care-app-net/"},
		link{"Scaling Rails @ Melbourne Ruby Meetup", "http://jrb.tumblr.com/post/30570014929/scaling-rails-at-melbourne-roro"},
	}

	// projects
	p.Projects = []link{
		link{"This site (Sinatra + SCSS)", "https://github.com/joho/whoisjohnbarton.com"},
		link{"Spoilerless Tour De France", "https://github.com/joho/letour"},
		link{"Loose (and fast) Postgres Row Counts", "https://github.com/goodfilms/postgres_loose_table_counts"},
		link{"Netflix API wrapper for Ruby", "https://github.com/goodfilms/netflix-ruby"},
		link{"HTTP Status codes 7xx (Developer Fouls)", "https://github.com/joho/7XX-rfc"},
	}

	// social urls
	p.Social = []link{
		link{"twitter.com/johnbarton", "http://twitter.com/johnbarton"},
		link{"goodfil.ms/john", "http://goodfil.ms/john"},
		link{"github.com/joho", "http://github.com/joho"},
		link{"jrb.tumblr.com", "http://jrb.tumblr.com"},
		link{"goodreads.com/johnbarton", "http://goodreads.com/johnbarton"},
		link{"au.linkedin.com/in/johnbarton", "http://au.linkedin.com/in/johnbarton"},
	}

	// quotes
	p.quotes = []quote{
		quote{"There is nothing to writing. All you do is sit down at a typewriter and bleed.", "Ernest Hemingway", ""},
		quote{"Don't you know who I am?", "Ben Schwarz", "http://germanforhipster.com"},
		quote{"I think we underestimate the complexity of deep sea cabling", "@deepjohnbarton", "https://twitter.com/deepjohnbarton/status/250782278349897728"},
		quote{"Ninety percent of everything is crap.", "Sturgeon's Law", "http://en.wikipedia.org/wiki/Sturgeon's_Law"},
	}

	return p
}

func main() {
	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		portNumber = "5000"
	}
	listenOn := fmt.Sprintf(":%v", portNumber)

	handler := newProHttpHandler()
	handler.ExactMatchFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := newHomePage()
		t, err := template.ParseFiles("views/index.html")
		if err == nil {
			t.Execute(w, p)
		} else {
			fmt.Fprintf(w, "Error parsing %v error was:\n\n%v", r.URL.Path, err)
		}
	})

	log.Fatal(http.ListenAndServe(listenOn, handler))
}
