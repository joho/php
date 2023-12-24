---
title: How Should Managers Code?
description: The interesting question isn't 'should managers code?' but instead how and why should managers code.
publishdate: 2018-07-03

ExpandedTitle: How Should Managers Code?
ImagePath: /post-images/hecate-logo.jpg
ImageCaption: With intention
---
If you're writing an all time top five list of controversial engineering management topics, I can almost guarantee that the question of whether managers should be coding will find a place on the list. I’m very firmly in the camp of managers who do, and believe they should, code but it's entirely understandable why so many people come down on the other side of the issue.

According to a recent study published by Harvard Business School, employees are happier when [managed by someone who can do their job](https://hbr.org/2016/12/if-your-boss-could-do-your-job-youre-more-likely-to-be-happy-at-work). To borrow from an old [NASA engineer's rules](https://spacecraft.ssl.umd.edu/akins_laws.html) "Capabilities drive requirements, regardless of what the systems engineering textbooks say.". When teams are building new software they're operating in the margin between the possible and impossible, and that margin is entirely defined by the capabilities of the people and technology in play at the time. The further a manager gets from the underlying technology (or the further the architect gets from the people) their strategic thinking is weakened by the loss of of an intuitive sense of one half of the capability dynamic.

The backlash against managers who code is largely driven by the incredibly common phenomenon of new managers screwing up their promotion by burning out doing their old job instead of learning how to do their new job. Managers should keep coding but they must not keep doing their old job, and it's easy to mix up those two things without the right framework to decide which code to write and how to write it.

For the past year or two this is the model I've used to explain to tech leads I've managed to pick and choose which coding work they should do. It might be generalisable to individual contributor work in general, but I've only managed teams of programmers, so you get the programmers version.

Any coding task that can be done has three kinds of value. I call them the "three Is" because you can't peddle any thought-leadership without a healthy dose of alliteration. 

## The Three Is

The three kinds of value a code contribution creates are:

* **Intrinsic** the value of what the code does
* **Instructive** the value of the example the code sets
* **Investigative** the value of what you learn in writing the code

Every time someone sits down and writes some code it will deliver some mix of those three values. The code will add some new functionality that is presumably of some value, it will serve as some kind of example of what to do (or not to do), and you will have learned something in the process about the system or problem domain.

For the individual engineers on a team, the intrinsic value of each contribution is the dominating factor. This is the kind of value that's used to prioritise a backlog - the importance of the capability the contributions add. It's usually the value that drives career advancement too - who makes the most and best things wins, especially in the middle stages of a technical career. In a company that doesn't value quality too highly it's just who makes the most things wins.

The values shift a little bit as you get into more senior engineering roles. That's where what you're teaching the team through the example you set in your code comes to the fore. Senior roles, in management positions or on the technical track, are about influencing your team to be better and the best way to do that is to lead by example. 

You keep doing work high in intrinsic and instructive value and you're liable to get yourself a promotion and that's where you come unstuck.

## The New Manager Trap

It takes a lot of time and energy to deliver work of high value, especially over multiple dimensions of value. If you've just been promoted you're used to delivering a lot of that value, but you have new pressing demands on your time and energy as a manager.

The first trap is motivational. It takes time to learn how to create the new kinds of value managers create, and even longer to intuitively grasp whether you're creating that value or not. In that motivational drought of not feeling like you're creating any value you'll reflexively fall back on the value you understand. This robs you of the time you need to spend to get better at your new job, and it robs your team members of the opportunity to step up into the gap you left behind.

The second trap is in the instructive value. Thanks to your promotion, the switch for instructive value is welded permanently on and probably signal boosted too. By virtue of your change in position, your words and actions carry heavier weight. When you're unaware of the dynamic behind the instructive value of your code this can play out in one of two ways.

* In trying to deliver code of high intrinsic value within the time constraints of your new role, something has to give and you've implicitly set an example that it's OK to give it up. This can be quality, this can be timeliness, but whatever it is you have normalised deviance and standards will slip.
* You intuitively grasp that your work sets an example and you strive to do so, but once again, something has to give and it's usually your work/life balance and burnout is just around the corner.

The important thing here is that now that you've recognised there are different kinds of coding value and that you can't deliver all three in the kind of time you can dedicate to coding and still do your job as a manager. Like most tech leadership lessons, it's all about cutting scope.

## A Plan to Keep Coding

I haven't talked much about that third value, the investigative one, yet have I? This is the value that you need to focus on from here on out as a manager. When Rands back-flipped on whether [managers should code](http://randsinrepose.com/archives/technicality/) this was the value he identified. Here's a quote from that post

> Even in a monstrous company laden with policy, process, and politics, you can’t forgot how to develop software. And how to develop software is changing. Now. Right under your feet, this very second. [...] My advice is that you take action so that you stay in touch with how your team builds stuff.

So you need a plan to maximise the amount of learning you can do via coding. It needs to be time efficient, otherwise you won't have time to do your main job supporting all the people you manage. It also needs to respect the unavoidable lessons your coding will teach. You will need to pick the right mix for your context but here are a few principles to include:

* You can never, ever again pick coding tasks by their intrinsic value. This is the domain of your team. If it's on the critical path it is for full time engineers alone.
* The exception to that rule is where you are trying to be instructive about different subtypes of intrinsic value. Maybe the company is rallying around a single goal for the quarter, but engineers are prioritising tasks by other metrics - lead by example.
* Because time is always short, try and pick tasks that deliver value aligned with the management work you're doing. Trying to troubleshoot a team in your portfolio? Pick up some of their work and understand them better. Talking about improving test quality? Go fix some tests and let your actions magnify your words.
* Pick some tasks at random. Management involvement in the day to day carries signalling risk so you need to give it camouflage.
* Give all coding a healthy air-gap after your promotion. Spend the time learning the new job, just make sure to come back to coding before you've lost all touch.
* It's healthy to let your instructive quality drop from your pre-promotion peak. Leave room for your senior and staff engineers to shine. Don't let it drop too much, bad coding will undermine your voice in technical discussions more than not coding at all. 

## The Prioritised Backlog and New Hires

This lens on coding value is useful in other contexts as well. Long before I'd put names to the different value types of coding I had a habit of keeping a "new hire backlog" off the books whenever I was hiring. Like many agile or kanbany shops, we always kept our backlog in priority order of business value. This worked great when the team was in a steady state, but I'd observed over a few hires that the engineers that got up to speed fastest were the ones who got lucky with tasks that took them all over the codebase. If the team was in a focused phase of work it could take months before a new hire wandered out from the first corner of the code they started in and their productivity would suffer for it.

To counter that, when I would go into a hiring phase I'd start holding back tasks in the backlog to build up a bit of a guided tour of the codebase. Some frontend work, some backend work, some ops work. Some in the legacy bits, some in the new stuff. It in a two sided marketplace so definitely some tasks for buyers and some tasks for sellers.

With hindsight, I'd figured out that the best way to maximise the amount of intrinsic value the engineer would create over their first year I needed to front load it with investigative value. As a manager, I'd strongly recommend you keep an alternative view of your backlog for these kinds of learning value, and share those tasks out between yourself and all your new hires.

## Fin

Managers should code to stay in touch with the work they manage but they should do so for the right reasons. Coding to learn more about their teams is good. Coding to reinforce teaching is good. Coding to get the code done for itself is bad. Coding while not doing their main job is bad.

_[Photo by Raphael Koh on Unsplash](https://unsplash.com/photos/92dgYPsir9k)_

[Discuss this post on Medium](https://medium.com/@johnbarton/how-and-why-should-managers-code-323751799664)
