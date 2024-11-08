# The state of the web

As of the time of writing this post in 2024 the World Wide Web (Web) has been around publically for 33 years. In those 33 years we've seen many paradigm shifts which have lead us far from the static pages and simple web applications typical of the web 1.0 and into the rich UI/UX of the web 2.0 and the social media age.

How did we get from html served over the wire to javascript frontend frameworks? What are the pros and the cons of the modern web? Can we do better? These are the questions I aim to answer with this post.

## The Web 1.0

The term "Web 1.0" is used to refer to the first stage of the Web which roughly coincides with the 1990's. At this time in the Web's history most users were consuming content produced by a relatively smaller amount of content creators who maintained mostly static websites.

With the release of Javascript and HTML 2/3, in the latter half of the 90's, the web application starts taking form (pun intended). All kinds of organizations from governments to local businesses began porting their processes to the web hoping (and failing) to reduce friction for those looking to access their services.

Full page reloads on every action, slow loading speeds, clunky UI and UX and broken applications are only some of the issues plaging these early Web Applications which, wether born from technical limitations or the understandable ignorance that comes from working in a new platform, are still the filter through which the Web 2.0 was created.

## The Web 2.0

The Web 2.0 is the age of APIs, User Generated Content and Social Media. Starting around 2004 and evolving over the past 20 years into the modern web we all love to hate but can't live without, the Web 2.0 has had an undeniable impact on the way we see and interact with computers.

I wrote and rewrote this part trying not to sound too biased but I am biased so hang in there.

Fast forward to the 2010's. We got Javascript running on the server due to NodeJS, Google releases angular an opens the flood gates of the client side rendering javascript web frameworks. In the interest of controlling my blood pressure let's just say that, in typical Web fashion, every time we found a minor inconvenience we reinvented the wheel.

- Client side rendering
- Server side rendering
- Hydration
- Partial Hydration
- Preprocessors
- Bundlers
- Typescript
- The Edge

These are only a few of the many examples of the over-complication the web has become since the early days of the Web 2.0

## Is this really that big a deal?

Yes.

- Complicated stacks create multiple points of failure
- Huge javascript bundles and unoptimized assets increase load times for users with poorer internet connections
- The over-reliance on layers of abstractions to solve problems that are not very difficult to solve creates left-pad (if you know you know)

## What are the solutions?

The current state of the web is built with the goal of creating rich interactable applications. A propper solution can't dismiss this goal.

We need:
- Fast load times (time to interactive)
- Interactivity 