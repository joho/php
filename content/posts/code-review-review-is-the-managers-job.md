---
title: Code Review Review is the Manager's Job
description: Code review is such a high value, high leverage team activity that every manager should ensure the process is done well.
publishdate: 2018-08-15

ExpandedTitle: Code Review Review is the Manager's Job
ImagePath: /post-images/hecate-logo.jpg
ImageCaption: Who watches the watchers?
---
In the classic management book [High Output Management](https://www.goodreads.com/book/show/324750.High_Output_Management) Andy Grove argues that "training is the manager’s job" because it’s the highest leverage activity a manager can do to increase the output of their organisation.

It’s still great advice for managers applicable across most teams, but for the manager of the modern software development team the fulcrum point has shifted. The code review, usually performed using GitHub pull requests or the like, is now the highest leverage point for improving an engineering team’s output.

## More than just quality control

Historically the code review process has been regarded as a quality control tool. Early empirical validation of the technique was overwhelmingly focused on its value in preventing defects. In the ten or so years since [pull requests entered the mainstream](https://johnbarton.co/posts/happy-tenth-birthday-pull-requests) the practice has evolved to be so much more.

Code review, whether via a pull request or other means, still is one of the best bang for buck methods for finding defects in code. Beyond the ordinary "first order" definition of quality which is whether it works or not, it’s also a great place to gatekeep “second order” quality, which is defects in the maintainability of a codebase.

Pull requests have also become the place where the team trains each other peer-to-peer, partially subsuming the role of manager as trainer. It’s one of the primary places where the team’s culture develops, especially if the team is distributed. It’s also the de facto information radiator for a development team, the best way to know how a product and codebase is changing over time is to be inside the code review loop.

A team with a good code review process is improving their codebase and improving each other. A team with a bad code review process will at best be stuck at a local maximum, and more likely be suffering from declining code quality and actively hurting each other. 

Given code review is such a high stakes process, how can any good leader afford to neglect it? It can be tempting to view the process as the sole domain of the engineers themselves, but like in [how (and why) should managers code?](https://hecate.co/blog/how-should-managers-code) I’d argue that’s a failure of outlook and approach rather than a domain in which management should not participate.

## Quality

Ensuring software meets the quality objectives of the organisation has been historically the largest driver of code review processes. In his [2006 post on code review](https://blog.codinghorror.com/code-reviews-just-do-it/) Jeff Atwood argued that every engineer should be getting their code reviewed and 100% of his argument rested on quality.

Discourse in the developer community is now swinging away from that framing. In his 2015 Railsconf talk [Implementing a Strong Code-Review Culture](https://www.youtube.com/watch?v=PJjmw9TRB7s) Derek Prior argues that "Code reviews are not about catching bugs. Modern code reviews are about socialization, learning, and teaching." I’m not sure it’s an either or situation, but whether you agree or not, ensuring all stakeholders in the review process have a consistent view of the purpose is important.

One really important thing to keep an eye on is whether a downstream department such as the QA or security teams has a process dependency assuming that code review has at least some coverage on defect detection. The example that springs to mind is during audit or security vendor questionnaire responses teams often use their peer review process to demonstrate "effective change control".

If your team drifts away from including defect detection in the code review mandate in such a context, as the engineering manager you should either gently nudge them back on track or instead go manage the stakeholder expectations.

## Training

I’ll borrow the summary of Andy Grove’s argument about training from [Ian Tien’s handy summary of the book](https://medium.com/@iantien/top-takeaways-from-andy-grove-s-high-output-management-2e0ecfb1ea63) 

> Training is the highest leverage activity a manager can do to increase the output of an organization. If a manager spends 12 hours preparing training for 10 team members that increases their output by 1% on average, the result is 200 hours of increased output from the 10 employees (each works about 2000 hours a year). Don’t leave training to outsiders, do it yourself.

When you can embed a good peer-to-peer training culture inside your code review process you take an already outstanding leverage point for team improvement and raise it to the next power. It can be particularly powerful if you can reinforce top down training initiatives with bottom up practices. Try hosting a lunch and learn around a specific coding practice and get a couple of your engineers to support the message and double down by focusing their review comments for the week around the theme.

If your team doesn’t currently have much of a training mindset within code review practice your top priority as leader should be to foster it. It is possible to over-do it and slide into an overly academic culture where introspection can take too much focus away from delivery, but I think it’s far more common to err on the side of under-investing in the educative benefits. Use your judgement to pick the right balance between sharpening your axe and swinging it at trees.

How specifically to develop that traning mindset in code review is a big topic and I recommend starting with the talks linked to on [Awesome Code Review](https://github.com/joho/awesome-code-review#talks-and-podcasts).

## Behaviour & Culture

For many teams, both distributed and not, GitHub is a "place" where work is done and culture is fostered. You can learn a lot about a team’s dynamics by diving into their pull requests. The professional and personal often blurs in pull request discussion, where comments can shift quickly from detailed line by line code analysis to sharing the meme of the hour to boost or criticise another developer.

[You can’t delegate company culture](https://relate.zendesk.com/articles/cant-delegate-culture-cant-buy-culture-where-to-get-company-culture/). As a manager, you’re responsible and accountable for the culture of your team. **Any manager responsible for the culture of an engineering group who has drawn a box around GitHub and said "this is an engineering space" and never visits is ****derelict in the performance of their duty**. 

If it sounds like I feel strongly on this topic it’s because I do. If I had to name a single casus belli for founding Hecate it would be when Coraline Ada Ehmke described [her year at GitHub](https://where.coraline.codes/blog/my-year-at-github/). The fact that an engineer was being singled out and treated so differently within the pull request process was appalling and I believe it is 100% a management failure.

As the manager, it’s your role to ensure that all the interpersonal dynamics at play within your sphere of influence are healthy and productive. You should stay across the review process enough to know what the social norms within your team and dig in to investigate when there is a departure from the norms. If the norms are unhealthy, you should be nudging the group towards healthier norms, and if that doesn’t work look to more drastic measures than mere nudges.

In the word’s of Australia’s former head of the Army David Morrison ["the standard you walk past is the standard you accept"](https://www.youtube.com/watch?v=QaqpoeVgr8U) - and if you disengage from the review process you are de facto walking past whatever the current standard is.

## But how?

So you agree that managing the overall health of your pull request process and culture is important. The natural next question is how to engage with that process?

There’s a fine line to be walked between managing the pull request process and micro-managing it. What happens within any specific pull request is the domain of the team so you should resist the urge to intervene frequently within the process itself. The behavioural trends in the process are your domain and if you’re running a meta-review process well your work will be predominantly behind the scenes.

Step one is to stay engaged with the pull requests themselves. When the team is small this is easy enough - you can follow along with the regular GitHub emails or Slack integrations. As the team grows and your meeting load grows heavier that can get trickier - coming back to your desk after four hours of meetings to a hundred GitHub emails is a bit of a killer. This is the use case for what I build the first Hecate product [Dispatch - batched PR notifications for managers](https://hecate.co/products/dispatch).

Step two is to develop your feedback patterns. Make sure you call out good reviews when you see them. If you do get the chance to contribute some code here and there, ensure that you set a shining example of code review etiquette as a reviewee.

Be careful of the "praise in public, criticise in private" trap. While generally good advice, because pull request comments are effectively a permanent public record (within the context of the org), leaving bad behaviour publicly unchecked can give the impression that such behaviour is OK. 

I’d advocate to generally start with a private nudge with the offender. You may have totally misread the interaction and can leave it at that. In the best case you can get them to apologise and correct the behaviour themselves. As an aside, I recommend leaving all but the worst comments unedited and adding an explanatory or apologetic comment afterwards rather than rewriting history. You can’t take back a hurtful comment, but by expunging them from the record it can gaslight the reviewee and undermine the group learning opportunity. If the offender declines to apologise, leave a public but gentle rebuke.

It’s important that the publicly viewable rewards and punishments for behaviour in code review align with the group’s values.

Step three is to track the trends in your team. In my first management role I kept a weekly diary of team activity which included particularly good or bad pull requests I’d seen that week. Schedule time monthly or quarterly to look back at your records and get a qualitative sense of how your team’s process is evolving. We’re working on a product right now called [Heartbeat - qualitative tracking of code review practices](https://hecate.co/products/heartbeat) which you can register interest for now if you’d like to try it out.

## Fin

Code review practice and tooling is the fulcrum point to improve so many dimensions of your team’s performance - don’t neglect it. While managers should be out of the day to day of code review they should maintain an active hand in the process.

If you like what we’ve got to say on the topic of management and code reviews, check us out at [hecate.co](https://hecate.co/). We’ve got a great newsletter, the form to subscribe is in the footer! This post is also [shared on Medium](https://hackernoon.com/code-review-review-is-the-managers-job-d412827a66c9) if you wanted to give it a clap.

_Photo by [Maia Habegger on Unsplash](https://unsplash.com/photos/Th6p15WAPP0)_