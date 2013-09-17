package homepage

import (
	"math/rand"
)

type link struct {
	Title string
	Url   string
}

type links []link

type quote struct {
	Text   string
	Author string
	Url    string
}

type homePage struct {
	Writings links
	Projects links
	Social   links
	quotes   []quote
}

func (p homePage) Quote() quote {
	quoteIndex := rand.Intn(len(p.quotes))
	return p.quotes[quoteIndex]
}

func New() *homePage {
	p := new(homePage)
	// writings
	p.Writings = links{
		{"Adopt The Pace of Nature", "http://blog.whoisjohnbarton.com/post/60153459731/adopt-the-pace-of-nature"},
		{"The Rules For Not Hiring Someone", "http://blog.whoisjohnbarton.com/post/60250771346/the-rules-for-not-hiring-someone"},
		{"Web Scale for the Rest of Us (RubyConfAU Talk)", "http://vimeo.com/61342269"},
		{"Interview for The Fetch", "http://blog.thefetch.com/2012/10/29/interview-melbourne-local-john-barton/"},
		{"The Minimum Viable Rails Stack", "http://goodfil.ms/blog/posts/2012/09/15/minimum-viable-rails-stack/"},
		{"Will App.net be the connective tissue founders & developers can rely on?", "http://thenextweb.com/apps/2012/08/18/why-10000-people-care-app-net/"},
		{"Scaling Rails @ Melbourne Ruby Meetup", "http://jrb.tumblr.com/post/30570014929/scaling-rails-at-melbourne-roro"},
	}

	// projects
	p.Projects = links{
		{"This site (Sinatra + SCSS)", "https://github.com/joho/whoisjohnbarton"},
		{"Spoilerless Tour De France", "https://github.com/joho/letour"},
		{"Golang port of dotenv (godotenv)", "https://github.com/joho/godotenv"},
		{"Netflix API wrapper for Ruby", "https://github.com/goodfilms/netflix-ruby"},
		{"HTTP Status codes 7xx (Developer Fouls)", "https://github.com/joho/7XX-rfc"},
	}

	// social urls
	p.Social = links{
		{"twitter.com/johnbarton", "http://twitter.com/johnbarton"},
		{"goodfil.ms/john", "http://goodfil.ms/john"},
		{"github.com/joho", "http://github.com/joho"},
		{"jrb.tumblr.com", "http://jrb.tumblr.com"},
		{"goodreads.com/johnbarton", "http://goodreads.com/johnbarton"},
		{"au.linkedin.com/in/johnbarton", "http://au.linkedin.com/in/johnbarton"},
	}

	// quotes
	p.quotes = []quote{
		{"There is nothing to writing. All you do is sit down at a typewriter and bleed.", "Ernest Hemingway", ""},
		{"Don't you know who I am?", "Ben Schwarz", "http://germanforhipster.com"},
		{"I think we underestimate the complexity of deep sea cabling", "@deepjohnbarton", "https://twitter.com/deepjohnbarton/status/250782278349897728"},
		{"Ninety percent of everything is crap.", "Sturgeon's Law", "http://en.wikipedia.org/wiki/Sturgeon's_Law"},
	}

	return p
}
