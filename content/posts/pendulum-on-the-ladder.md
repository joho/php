---
title: "Engineering Management: The Pendulum on the Ladder"
description: Is it necessary to swing the pendulum between management and engineering roles to remain a competent technical leader?
publishdate: 2019-03-21

ExpandedTitle: "Engineering Management: The Pendulum on the Ladder"
ImagePath: /post-images/hecate-logo.jpg
ImageCaption: Bing. Bong.
---
In January this year Charity Majors wrote [Engineering Management: The Pendulum or the Ladder](https://charity.wtf/2019/01/04/engineering-management-the-pendulum-or-the-ladder/) as a follow up to [The Engineer/Manager Pendulum](https://charity.wtf/2017/05/11/the-engineer-manager-pendulum/). It's a great article with a realist lens on the career paths of software engineers going into management. It's obviously a position built out of real world experience and provides a viable career template if you want to keep your technical and leadership skills in balance. It's damned fine advice and it works.

Despite all the positives for some reason I never shared the article within my community. It's the career path I've followed: developer for a few roles, ran the engineering team at [Envato](https://envato.com/) for a few years, went back to the tools as co-founder of [Goodfilms](https://goodfil.ms/blog/), lead the engineering team at [99designs](https://99designs.com/), and now back on the tools as the founder of this fine site.

It's taken me a while to gather my thoughts on why despite personally following the same career path I've not passed that advice on to others. **It's because while Charity's advice is great for succeeding within the status quo, the status quo isn't good enough** and we should aim for something better.

Before getting into it you should read the post yourself but to briefly summarise her position:

* people management is a new skill that takes time (two years) and effort to develop to a serviceable level
* time spent developing management skills is time spent losing technical skills
* the implication of that dynamic is that there are career traps, both of employability and fulfilment when you naively climb the ladder of the default career path
* the best way out of the trap is to either deeply commit to the people management ladder with eyes wide open or to swing on the pendulum between leadership and individual contributor roles

As I said before, great advice, but I feel echoes of another past false dichotomy.

# DevOps: the last irreconcilable tech identity crisis

The last time I had a philosophical disagreement with a long standing role split in our field was during the emergence of the "DevOps movement".

Ten years ago conventional wisdom was that you had two main career paths working on the web 2.0: to build the app as a developer or to run the app as a sysadmin and never the twain shall meet. The highest performing teams in that era - Flickr is the one that comes to mind - said that if you worked really hard you could have your dev team and your ops team get along really nicely.

So I try and transport Charity's advice backwards in time and see if it fits that discourse too. You're a developer. Becoming a sysadmin is not a promotion. It's a new job, with new skills. You need to stick at it for a couple of years. Your developer chops will wane after a couple of years, so to really stay current you need to jump back to your old job. To be the best engineer you can be you do need a balanced understanding of developing and running your software so you should pendulum between those two roles.

It's a clunky simplification but largely holds up. What's really interesting to me is how it would have been great advice back then but it completely misses the rise of the developers supporting their own software. I'm hoping that the same is happening again now, that while the pendulum is great advice, that **hopefully we're coming into an era where us engineers can manage ourselves**. Specialised roles in operations and development didn't go away, but for huge swathes of the industry more rounded out engineers was a big benefit.

# Is having technical managers a worthwhile goal?

That question is worth an entire post on its own but let's assume it's a yes for now.

Employees are happier [working for someone who can do their job](https://hbr.org/2016/12/if-your-boss-could-do-your-job-youre-more-likely-to-be-happy-at-work). By being [managed by bureaucrats](https://twitter.com/caseyrosenthal/status/1039726304818601985) rather than managing ourselves we can be pointed towards purposes we may not agree with, and managed using methods we do not like. Hackernews routinely fills up with posts and comments decrying sprint to sprint marches building things of little to negative net value. **As an industry do we expect this to get any better when we outsource our own leadership to those who have never been or have left our profession?**

Charity rightly points out that so many of us that climb that ladder end up unhappy. Isn't that reason enough to try something else? We've parleyed the scarcity of engineers into free snacks and high salaries, why not direct it towards building a ladder that more of us would want to climb, one that appropriately balances the power and duty of leadership?

# Does growing as a manager mean you have to stop growing technically?

Ultimately it's a question of time and priorities. There are only so many hours in a day and they need to be carefully allocated across the varying skill domains.

The initial learning curve of becoming a manager is a tough one to climb. I think Charity gets it right in saying you need to dedicate two years to giving it a proper go. I also think the three to five year "slide" in technical skill sounds right. This asymmetry creates the opportunity to keep a foot in each camp: focus intently on the new manager skills for a couple of years, then bring your goals and focus into a more balanced program of personal development.

Ever since [Maker's Schedule, Manager's Schedule](http://www.paulgraham.com/makersschedule.html) was written the common belief has been that there's some inviolable law that you cannot fit programming and managing in the same calendar. At a surface level they're definitely incompatible tasks that carry a high cost in context switching &dash; but rather than telling me it's an unsolvable problem it highlights how important the environmental context will be to make a change.

Some may try and make an argument that technical work and people work are of fundamentally different types and cannot be mixed. I would find such an argument unconvincing; a repeat of the thing/person dichotomy that makes no sense except as a tool to gate-keep certain classes of people from different kinds of work. It may be hard or uncommon to see those types of work mixed together, but I have seen it, and that's enough of a motivation for me to try and bring more of it into the world.

# What can we hope for?

During my time leading the team at [99designs](https://99designs.com/) we came pretty close to keeping everyone at each level of the engineering organisation technical. It took a special confluence of events to get it there, and technical time certainly got thinner the higher in the tree you were, but it's given me hope that we can blend technical and people leadership roles in the future.

Right now it feels like there is rough consensus that you can still stay technical as a manager at the line level working directly with engineers, at least within the right teams. I think a sensible mid-term goal would be for more teams to nudge it up one more rung in the org chart. To have all line managers and their managers be hands on, within sensible boundaries of time and purpose, on a continuous basis.

**Rather than forcing technical leaders to swing the pendulum every two or three years when they change roles or companies we let them swing it every few weeks, or even every few days.** We've learned so much about reducing batch size when building our software we should reduce our batch size on building skills within different silos.

I don't think the movement will take root in the larger, pre-existing firms but instead in the businesses we're starting today - the same way developers running their software in production emerged.

# Ok, how can we do it?

Looking back to the DevOps emergence, we didn't get here just by exhorting developers to do more. It took a wide range of changes within the environment to make it possible. 

Off the top of my head I think server virtualisation was a big one, giving us VPS providers and then "the cloud", freeing end developers from racking and provisioning machines. Rails kicked off a general trend of accessible automation, starting with database migrations, then Capistrano, then Chef and Puppet that gave us more leverage on our time in ops work. General agile and business trends had moved a lot of thinking away from separate build and operational phases of software project, and the underlying CapEx/OpEx accounting that naturally split development and operations teams.

So what do we need going for us to change the face of the management track in tech?

## Stakeholder Support

We need the support of everyone around us in the organisation. In the worst case a big cultural change initiative is doomed to fail without leadership support. In the best case it's bad manners. It's a big bet to make such a change and to do without the buy in of those who are ultimately on the line for the change.

The rising "indie" or "fundstrapped" scene seems like a more fertile ground for experimenting than companies operating under the VC model, if only because there are fewer stakeholders to get on board with the experiment. Executive support is one thing, and board support is a whole other one.

It was tricky enough getting a "no ops team" staffing plan past a board ten years ago who weren't pattern matching through a broad portfolio. I can't imagine the rhetorical judo necessary to get your VC partner on board until someone else has done it first.

We also need buy in from the engineers on the value of technical managers. If your engineers believe a manager's primary function is to hold a shit umbrella and stay out of their way, they're not going to be the partners you need in blurring the line between tech and management.

Manager's schedules don't fill past capacity for no reason at all, and part of making space for remaining technical will be sharing some of that load down into the team. My experience has been that engineers generally prefer reporting to someone technical, and when you phrase the benefit of sharing some of the management tasks around the team as a trade-off to ensure they can continue reporting to someone technical it usually goes pretty well.

## Structural change

The big thing we learned at 99designs trying to keep our leadership technical was that we had to hold the number of direct reports per manager low. Common practice is to shoot for five to seven direct reports per manager, but I think to make this work you need to hold it around four or five. Time is a very hard constraint and the non-linear growth in communications channels between reports will eat that time.

To lead effectively you need to keep a full and rich emotional image of each team member in your head. So many parts of management require empathy and with seven humans in your team you will fill up your RAM pretty quickly, making the context switch out to technical work a full swap to disk.

## Tools and process

We spend so much time seeking for mechanical leverage on our engineering tasks, why aren't we doing the same on our management tasks? By streamlining the essential tasks of management that leaves more time for the discretionary tasks of management, and instead of filling that new hole in your calendar with another meeting - fill it with technical work.

We need to find the management equivalent of racking servers and figure out a way to abstract it away. We're not there yet, but as more and more teams push their workspaces online we've got great opportunities to make the machines do some of the work. It would be a crying shame for our collaboration tools to buy us managers more time, only to use that leverage to manage even more direct reports.

## Fin

Charity's post is right but is only advice on how to succeed within the status quo. The status quo keeps the human and technical aspects of our work unhealthily separate. Every day some new ethical shitshow is unveiled in the tech industry and we keep clamouring for a career path that takes away from the levers that can fix them. For many engineers our technical choices are also deeply personal, and the idea that technical decisions can be made on one track isolated from the personal development of those who need to implement them sounds ridiculous.

I founded Hecate to try and get some traction on this problem and enable compassionate technical leaders to get more leverage on their days. If it doesn't work out, I'll keep following Charity's advice and get back on the pendulum. In the meantime, I'm going to keep working towards making more technical management practical.

If you're looking for more leverage as a dev manager, check out [our products](https://hecate.co/#products) or get in touch via [email](mailto:hello@hecate.co) or [Twitter](https://twitter.com/HecateApp) as I always love chatting to like-minded tech leaders.

_[Cover Photo by Enrique Fernandez on Unsplash](https://unsplash.com/photos/Nmewv8jDfyE). Thanks to [Jon Williams](https://twitter.com/jonathannen) and [Erwin van der Koogh](https://twitter.com/evanderkoogh) for review.