package pages

import (
	"io/ioutil"
	"time"

	"github.com/russross/blackfriday"
)

// Page is a page to be rendered
type Page struct {
	Title         string
	ExpandedTitle string
	Description   string

	PublishDate *time.Time

	ImagePath    string
	ImageCaption string

	Path string

	ContentPath string

	ShowPostList bool
}

var absoluteRoot = "https://johnbarton.co"

// AbsoluteURL is what it says
func (p *Page) AbsoluteURL() string {
	return absoluteRoot + p.Path
}

// AbsoluteImageURL is for og:image etc
func (p *Page) AbsoluteImageURL() string {
	return absoluteRoot + p.ImagePath
}

// Content returns HTML transformed from markdown from file specified in ContentPath
func (p *Page) Content() (string, error) {
	file, err := ioutil.ReadFile("pages/content/" + p.ContentPath)
	if err != nil {
		return "", err
	}
	content := blackfriday.Run(file)

	return string(content), nil
}

// PublishDateFormatted returns ISO8601/RFC3339 date
func (p *Page) PublishDateFormatted() string {
	if p.PublishDate != nil {
		return p.PublishDate.Format(time.RFC3339)
	}
	return ""
}

func (p *Page) RecentPosts() []*Page {
	return AllPosts()
}
