package pages

import (
	"time"
)

var postsReverseChronological []*Page

func AllPosts() []*Page {
	return postsReverseChronological
}

func init() {
	postsReverseChronological = []*Page{
		&Page{
			Path: "/posts/hello-world",

			Title:         "Hello World",
			ExpandedTitle: "Hello World",
			Description:   "A test post",

			PublishDate: publishDate("2017-12-31"),

			ImagePath:    "/john-barton-dithered.png",
			ImageCaption: "testing",

			ContentPath: "posts/hello-world.md",
		},
	}
}

func publishDate(date string) *time.Time {
	time, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil
	}
	return &time
}
