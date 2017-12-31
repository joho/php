package pages

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

			PublishDate: nil,

			ImagePath:    "/john-barton-dithered.png",
			ImageCaption: "testing",

			ContentPath: "posts/hello-world.md",
		},
	}
}
