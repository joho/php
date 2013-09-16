package homepage

import (
	"math/rand"
)

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

func New() *homePage {
	p := new(homePage)
	// writings
	p.Writings = []link{
		link{"Adopt The Pace of Nature", "http://blog.whoisjohnbarton.com/post/60153459731/adopt-the-pace-of-nature"},
		link{"The Rules For Not Hiring Someone", "http://blog.whoisjohnbarton.com/post/60250771346/the-rules-for-not-hiring-someone"},
		link{"Web Scale for the Rest of Us (RubyConfAU Talk)", "http://vimeo.com/61342269"},
		link{"Interview for The Fetch", "http://blog.thefetch.com/2012/10/29/interview-melbourne-local-john-barton/"},
		link{"The Minimum Viable Rails Stack", "http://goodfil.ms/blog/posts/2012/09/15/minimum-viable-rails-stack/"},
		link{"Will App.net be the connective tissue founders & developers can rely on?", "http://thenextweb.com/apps/2012/08/18/why-10000-people-care-app-net/"},
		link{"Scaling Rails @ Melbourne Ruby Meetup", "http://jrb.tumblr.com/post/30570014929/scaling-rails-at-melbourne-roro"},
	}

	// projects
	p.Projects = []link{
		link{"This site (Sinatra + SCSS)", "https://github.com/joho/whoisjohnbarton"},
		link{"Spoilerless Tour De France", "https://github.com/joho/letour"},
		link{"Golang port of dotenv (godotenv)", "https://github.com/joho/godotenv"},
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
