---
title: 2017 In Review2
description: A quick braindump of the year primarily to test out the new blogging feature I built for my site.
publish_date: 2017-12-31

ExpandedTitle: 2017 In Review
ImagePath: /post-images/top-nine-2017.png
ImageCaption: My Instagram top nine of 2017
---
Over the past few days Charity Majors has been at the centre of a wide ranging twitter debate around whether developers should go on call. Her position, which I agree with 100%, is that developers should.

Because it's a conversation that's gone on many tangents, it's hard to pick a single twitter thread that's representative. This one in the middle is pretty good. 

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">All this heated talk about on call is certainly revealing the particular pathologies of where those engineers work.  Listen:<br><br>1) engineering is about building *and maintaining* services<br>2) on call should not be life-impacting<br>3) services are *better* when feedback loops are short</p>&mdash; Charity Majors (@mipsytipsy) <a href="https://twitter.com/mipsytipsy/status/962151928741285888?ref_src=twsrc%5Etfw">February 10, 2018</a></blockquote>
<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

## The haters

I can completely empathise with why discussing on call can be challenging. Charity herself mentions that a lot of people have PTSD from bad on call experiences. On call has been appalling for many people in many places, but none of that is inherent in on call itself.

There are two frustrating themes in the opposing side of the debate right now.

1. An almost wilful denial that a humane model of on call can (and does) exist
2. A betrayal between castes of engineers, that developers can't be on call because they have valid personal reasons, but ops should because it's their job

My feeling is that the first is a reflexive defensive posture of developers to avoid thinking through their unacknowledged position on the second. The really sad part of this dynamic is that while there remains a small moral failing in the general area of "developer privilege", that in examining rather than rejecting humane on call programs you'd discover the underlying systemic problem of companies trying to get a free ride on 24/7 availability hurting everyone, including the companies themselves.

## Developer Privilege

Let's deal with the elephant in the room first. Ops engineers typically get paid less and treated worse than developers. They're expected to carry the pager and keep software running 24/7. The two biggest factors in the quality of life for this role is the rate usage of the software grows and the quality of the software as designed and built, neither of which they have much direct influence on.

I cannot rightly comprehend a universe where the accident of your technical speciality makes you accountable for the quality of someone else's work.

<blockquote class="twitter-tweet" data-conversation="none" data-lang="en"><p lang="en" dir="ltr">I *do* mean to be hurtful and dismissive to those who use their status to push their share of the software engineering maintenance load onto lower status engineers, instead of putting shoulder to wheel and making it *better*.<br><br>Solidarity, man.</p>&mdash; Charity Majors (@mipsytipsy) <a href="https://twitter.com/mipsytipsy/status/962204888871591936?ref_src=twsrc%5Etfw">February 10, 2018</a></blockquote>
<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

I don't think I can say it any better than that.

I'm also not a fan of the arguments that "I can't go on call because I have a family" or "I like camping so I can't be on call". Do none of you have any friends in other industries? I guess nurses can't have families. I guess no doctors go camping. Electricians can't join organised sports leagues. No train drivers can go to the movies.

Systems that run around the clock need around the clock coverage, whether that's keeping people alive or selling doo-dads via a webform. Part of why software is so remunerative is that it facilitates round the clock sales (or activity or whatever) 24 hours a day, every day of the year.

When developers are happy to share in that windfall, and then lean on arguments that fly in the face of the lived experience of other professions to avoid sharing in the toil, I run out of all empathy. There will always be exceptional circumstances, but the answer to that is to participate by default and handle the exceptions, not opt an entire industry out to avoid the edge cases.

## Humane On Call

That said, if your company has nothing but a hellish on call experience to offer, you'd be insane to opt in. There are (I assume) many ways to design a humane on call system, but I just want to share a bunch of quick points about policy I've instituted across two teams that makes on call workable.

* Pay for on call. Pay a retainer for the rotation period (a week where I've worked) to cover the opportunity cost of the cinema you didn't go to or nights drinking/camping you can't do. Pay a call out fee for responding to an alert. I'll talk more details about why this is important in the next section.
* Rotations should be between 5-10 people. Too few and you have an out-size and unreasonable impact on out of work activities. Too many and the rotation becomes ineffective. Skills need to be exercised to be retained, and one in ten weeks is just barely frequent enough.
* Heavily curate which alerts go to a pager. Not all systems or errors are created equal, and when you can put a dollar price on paging out you would be amazed at how easy that process becomes.
* Let people swap on call time, from a few hours (to see a movie) to a whole rotation (to go on holiday). Only allow swaps, not giving them away. Everyone takes a fair share, but timing is flexible.
* Have an escalation policy. It all becomes more manageable with support. Don't punish people pulling in help from the next level.
* Let the team fix things that break. Product engineers who have just responded to an alert have the perfect mix of motivation and context to get a fix in place.
* Train your people. On board them well to on call. Do workshops. Have a game day. Schedule frequent refresher courses.

I can't say the above would work in every context, but it's worked for me across a couple of roles so far.

## The Free Rider Problem

I mentioned above that in my teams I run paid on call rotations. It's now a core part of my management philosophy, but I arrived there by pure copy and paste accident. The short version is that in a past startup we (two out of three engineers) were told that we needed to share the on call support to lighten the burden on the founding engineer of the company. After a knee jerk reaction of "fuck that", I remembered that my mum - a nurse - was on call relatively frequently. I called her up, asked how on call worked for her - which was she was paid an extra retainer of 10% for every hour she was on call, and was paid double time if she was called in to work, and a minimum payout for response. If mum could do it so could I. Clinton (the other non-founding engineer) and I came back to the table and said we'd participate on those terms.

That startup was called Envato and it runs [one of the highest traffic Ruby on Rails sites in the world](https://www.alexa.com/siteinfo/themeforest.net). They still run the program roughly that way, and I made the same policy at 99designs when I joined the management team there.

It wasn't until much, much later that I realised it (roughly) solves [the free rider](https://en.wikipedia.org/wiki/Free-rider_problem) and sets up a virtuous quality cycle by aligning business and engineering incentives. Unpaid on call allows a business to profit from round the clock revenue generation from an always-on platform without paying its share in the costs of maintaining round the clock revenue generation.

* Once there's a set cost to maintaining a rotation, deciding which products and services need round the clock availability becomes an easy business decision
* Once each page comes with a minimum cost (one hour at double pay), separating "informative" alerts from actionable ones becomes an easy business decision
* Once software quality degrades to the point of services frequently paging and the costs stack up, ensuring the remediation work is completed becomes an easy business decision
* Engineers still tend to value personal time (sleep, etc) more than the call out fee, and so fixing their software becomes an easy personal decision

Paying for on call isn't a panacea. Incentives aren't a silver bullet for behavioural change, but when a company can page an engineer for free, it's no surprise there are so many bad actors.

## Let me sum up

Most companies offer a shitty on call experience, but that sets the floor, not the ceiling for what on call can be. Teams where the people creating the software help support the software achieve better quality through aligned incentives and increased awareness. This has not been the historical practice in industry, but companies can make the change by acknowledging and addressing their role in making quality an economic externality on ops teams. Developers who have benefited from the status quo should engage more deeply with the underlying issues rather than pushing quality onto another class of engineer.

