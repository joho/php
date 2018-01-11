As far as I can tell there are roughly three flavours of working style in an "agile" environment that all get named "iterating".

Let's call the three flavours scientific, greedy, and defensive. There's definitely some overlap between the three. I want to explore them a little in this post a lot of the interesting (BYO airquotes) moments in engineering management come up when the different varieties rub up against each other.

For the sake of this argument these are the definitions I'm going to use.

* Scientific iteration is about shipping to prioritise learning. This is where the Eric Ries Lean Startup fans hang out.
* Greedy iteration is about shipping incrementally to get the value of each increment sooner. This is where the Dan Reinertsen Cost of delay kids are.
* Defensive iteration is building incrementally parts of a whole, not for the value of any increment specifically, but so you've defended against having an incoherent whole when an arbitrary product cycle closes. This is the domain of "potentially shippable increment" scrum types.

The three flavours somewhat map onto [Ken Beck's Product development triathlon](https://www.facebook.com/notes/kent-beck/the-product-development-triathlon/1215075478525314/) of explore/expand/extract. You could also describe them as the team working styles of [Pioneers, Settlers, and Town planners](http://blog.gardeviance.org/2015/03/on-pioneers-settlers-town-planners-and.html) from Simon Wardley. They're not completely analogous because Kent and Simon's models tend to describe phases of a lifecycle that flow from one to the other. While a particular iteration style may suit one period in a lifecycle better than another, there's no hard linkage there.

I just want to quickly burn through some definitions with some observations, then dive into a couple of scenarios where I think the ideas come in to play.

## 1. Scientific Iteration

This is what in this day and age you could reasonably call "the lean startup" development methodology. Though I think the working style predate, and will live long past, the Eric Ries book. It's delivering incrementally to uncover insight. To drive sales and nudge metrics not for the purpose of driving sales or nudging metrics, but to gather proof from changing those numbers. The sales are nice if you can get them though.

It's implicit in working this way that you are gathering this data in order to act on it, and if the knowledge you gain tells you to do something as large as a pivot, then pivoting is the right thing to do. Given that making big changes is on the table, any development style should heavily lean on prototyping. Building to last when new information could mark a radical change in direction is an inefficient use of time and energy.

I'm not saying those working in the other modes don't prioritise learning, but that in this mode each increment's success is judged by what it teaches. It seems an intuitive match for Beck's explore phase for reasons so intuitive I feel dumb trying to list them out. I don't believe it's as natural a match for the Wardley's pioneer mindset. Some pioneers will definitely iterate scientifically to find something of value, but I think just as many or more strongly believe in something and work their way towards it, working more in the greedy fashion.

## 2. Greedy

I call this style greedy because I think of the Poppendiek's "Concept to cash" subtitle to one of their books. It's the mode of incremental delivery that judges the success of each increment in the value it delivers. Usually measured in cash. Or some proxy metric that you really, really believe will eventually mean cash. Sometimes measured in other units like warm fuzzies or social impact.

Slicing up a larger vision into deliverable slices is a fine art in this mode. Each slice needs to add something of value to your end user. If you're doing it well, the slice will also have accretive value to the next slice you want to build. Working so that every increment has realisable gains without undue rework before the next game is really hard. This is what Alistair Cockburn means when he says "The purpose of playing this game well is to be able to get the best position on the next game." in [Software development as a cooperative game](http://alistair.cockburn.us/Software+development+as+a+cooperative+game).

You're trading off the value of a coherent, fully realised product against the [cost of delay](http://leanmagazine.net/lean/cost-of-delay-don-reinertsen/) of some subsection.

## 3. Defensive

Of the three kinds of iterating, this is the kind I'm least sure really exists as a freestanding concept. I'd broadly describe it as incremental building where delivery of each increment is optional. It's the realm of the ["Potentially Shippable Increment"](https://www.infoq.com/news/2008/02/done-shippable-quality). Each increment is valid if it could conceivably be the last increment. That's the theory at least. In practice, the "potentially finality" of any increment is inversely correlated with the perceived length of time until a preconceived stopping point.

This is the style of agile-in-waterfall, where you iterate a lot in your development phase. Early tasks are rather haphazardly sliced and are "potentially shippable" only because they won't break anything, not because they will add any value. As you close in on your deadline or budget is running out then each task as a checkpoint of useful software becomes more important.

It's a model that contains both traps and opportunities. The trap is that because "potentially shippable" is so fuzzy, each incremental step could be generating no forward progress and you find out too late. The opportunity is because "potentially shippable" is so fuzzy, each incremental step can be carved up smartly to minimise rework or supporting scaffolding for an efficient build.

## Accidental misalignment

The motivating incident behind this post was a meeting a few months ago. I had an engineering team who were not getting along super well with their counterparts in product and UX. No one could agree if their proceeding project was a success or not and there was a lot of tension in the kick off of their upcoming project. It was then I realised although everyone had agreed they needed to "iterate" their way to a solution, in fact they had wildly different expectations of what that meant. The engineers had a scientific style in mind, prioritising delivery to verify the project was a good idea. The product manager had more of a defensive model in mind, to work towards a relatively preconceived stopping point and deliver it. Once we had that out in the open, we worked towards a compromise of working more in the greedy style.

It's with hindsight I then look back at a lot of, if not failed projects, certainly unhappy projects and recognise that this false alignment on iterating style has been a major, recurring problem. I also look at online debates between different agile/management thoughtleader types and see the same failure to come to common terms.

Save you and your teams some heartbreak - double check you've all agreed to work in the same style.

## The best and worst of both worlds

The defensive style is a thorny one. As it can be a bit of a hybrid between agile and waterfall it has the potential to be the best or worst of both worlds.

On balance, I think it leads to more failed agile projects than successes, as it's the resting place of failed enterprise agile transformations. The "potentially shippable increment" in larger conservative environments creates so much cover to stop you properly learning to carve up work into real shippable increments. Even worse is if you go through the pantomime of each part being shippable knowing full well that it won't. It all adds up to extra work without any corresponding benefit.

I think this mode is only really safe after you've comfortably learned to work in the scientific or greedy mode. Real incremental delivery often involves a two steps forward, one step back of effort as you ship a bit of extra supporting infrastructure for something that's meant to be part of a larger whole stand alone. There can be a lot of saved energy in construction if you can forgo that scaffolding for part of your project lifecycle.

## Courses for horses

All of the iteration styles have value. Each style can be a better or worse fit for a project. Each style can be a better or worse fit for an individual. Which style you work in (I believe) has less impact on successfully building a product than in making sure you're all working in the same one. Not everyone is comfortable moving between styles, so when you do flush out one of these accidental differences be careful. It's hard to compromise between these working styles (without being  wasteful) so someone will need to pick a winner. Acknowledging the differences can go a long way towards smoothing them over.

In short: iteration can mean a few different things, each thing has pros and cons but misalignment is all cons, so double check you're all aligned and be kind to each other.

