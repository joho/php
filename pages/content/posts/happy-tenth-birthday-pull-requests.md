Happy 10th Birthday Github Pull Requests! Introduced on February 23rd 2008 in [their 3rd blog post](https://github.com/blog/3-oh-yeah-there-s-pull-requests-now).

My guess is that in discovering PRs are 10 years old you're having one of two reactions.

1. Oh my god, has it been 10 years already?
2. What do you mean Pull Requests got invented? How else would all of this work?

My reaction was definitely the first one. I'd stumbled across the introduction post while I was researching the history of code review, which in turn was prompted by the Go community's discussion regarding Gerrit vs Github. 

It's amazing to think about how over just 10 years the Pull Request has become the dominant paradigm of code review. Even if you're not using Github, you're still probably using a tool that has outright copied or been influenced by Pull Requests. Congratulations to Chris Wanstrath and Tom Preson-Werner on creating something with such enduring influence.

It wasn't an overnight success. If you read Scott Chacon's [Github Flow](http://scottchacon.com/2011/08/31/github-flow.html) post from 2011 you can see:

> GitHub has an amazing code review system called Pull Requests that *I fear not enough people know about*. Many people use it for open source work - fork a project, update the project, send a pull request to the maintainer. However, it can also easily be used as an internal code review system, which is what we do. (emphasis mine)

When I started my career code review was an esoteric practice infrequently done. Shit, even having code under source control was a rare enough practice to be worth using as a [ranking signal of team quality](https://www.joelonsoftware.com/2000/08/09/the-joel-test-12-steps-to-better-code/). The two kinds of review I had as a junior were either in person or after the fact. Branching support was relatively poor in the pre-git source control tools, so generally an old commit would be at the size you expect a branch to be nowadays. I'd ask around the office for someone to come over to my desk to do a code review, and we'd walk through the diff discussing the change before I press commit. There would also be the chance I'd get some follow up review, when my boss would read the commit in the ugly as hell email that CVS or Subversion would send out to the team every time someone committed.

That all sounds so terribly uncivilised now, thanks Github for making code review more elegant and accessible.

If you're interested in code review in general, it's kind of my general theme in writing at the moment and there'll be some more posts on the topic. I'll update this list with links as I go, but it's worth subscribing to [my feed](https://johnbarton.co/feed.xml) or [my Medium](https://medium.com/@johnbarton) to follow along.

* [Awesome code review](https://github.com/joho/awesome-code-review) - a collection of resources about code review
* Five Factor Code Review - will be an homage to [Sarah Mei](https://twitter.com/sarahmei)'s [Five Factor Testing](https://www.devmynd.com/blog/five-factor-testing/) article except about, you guessed it, code review
* [Code Review Review is the Manager's Job](https://hecate.co/blog/code-review-review-is-the-managers-job) - Why and how managers should stay engaged in the code review process
* Code Review Literature Review - I've bought a bunch of old books about code review so you don't have to
* ...and any other random ideas that spin out while I'm writing those
