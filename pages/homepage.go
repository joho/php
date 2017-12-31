package pages

func Homepage() *Page {
	return &Page{
		Path: "/",

		Title:         "John Barton",
		ExpandedTitle: "John Barton - Engineering Lead",
		Description:   "Hi, I'm John Barton. Right now I head up the engineering team at 99designs where I herd people and ship the odd bit of Ruby and Go code. My true passion is boring friends and co-workers with slideshow presentations of Nautical History Facts.",

		PublishDate: nil,

		ImagePath:    "/john-barton-dithered.png",
		ImageCaption: "At home, with books",

		ContentPath: "home.md",
	}
}
