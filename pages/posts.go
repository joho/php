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
			Path: "/posts/three-flavours-of-iteration",

			Title:         "Three flavours of iteration",
			ExpandedTitle: "Three flavours of iteration",
			Description:   "In the project kickoff meeting your team has all agreed to iterate towards a solution. Sounds great. Everyone's agreed. (Narrator's voice) They hadn't",

			PublishDate: publishDate("2018-01-11"),

			ImagePath:    "/post-images/catlike-agility.jpg",
			ImageCaption: "This kind of agility?",

			ContentPath: "posts/three-flavours-of-iteration.md",
		},
		&Page{
			Path: "/posts/2017-in-review",

			Title:         "2017 In Review",
			ExpandedTitle: "2017 In Review",
			Description:   "A quick braindump of the year primarily to test out the new blogging feature I built for my site.",

			PublishDate: publishDate("2017-12-31"),

			ImagePath:    "/post-images/top-nine-2017.png",
			ImageCaption: "My Instagram top nine of 2017",

			ContentPath: "posts/2017-in-review.md",
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
