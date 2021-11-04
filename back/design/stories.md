# User Story Workflows

## Stories Slated for OpenBoard MVP

These stories were originally written at https://docs.google.com/spreadsheets/d/1lmktJ9qTNuslGLlE-HCsuh9S4nUrxaftqcn6jr9FNgs/edit?usp=sharing.

Sage is the sample user for the following stories.
Sage is the typical user of the OpenBoard website.
The following stories are based on Sage's first time using the site as a walk-up customer.

Sage's goal is to easily access the information ze seeks.

The stories below are split into parts.
Parts are ordered and indicate the stages of interaction Sage goes through when navigating the OpenBoard site.

### Community Statistics: Ask, offer, and Member Counts

Story Part 1: Initial Perusal of Posts

When Sage goes to the home page, ze sees community stats: 

- active ask count
- active offer count

Blocking Issues 

- Add "Kind" message enums [#163](https://github.com/OpenEugene/openboard/issues/163)
- Expand FndPostsReq [#167](https://github.com/OpenEugene/openboard/issues/167)
- Add "Settled" and "Expired" timestamps to Post messages [#173](https://github.com/OpenEugene/openboard/issues/173)

### Paginated Cards with Post Details and Content

Story Part: 1: Initial Perusal of Posts

Sage sees paginated cards containing the following post details:

- type (ask or offer) 
- status (e.g. closed)
- author
- title (summary)
- description
- creation date
- category (e.g. "Jobs")

Blocking Issues

- Add "Kind" message enums [#163](https://github.com/OpenEugene/openboard/issues/163)
- Add Categories message to posts.proto [#165](https://github.com/OpenEugene/openboard/issues/165)
- Add "Categories" message to other messages [#166](https://github.com/OpenEugene/openboard/issues/166)
- Expand FndPostsReq [#167](https://github.com/OpenEugene/openboard/issues/167)
- Add "search by user id" to FndUserReq [#169](https://github.com/OpenEugene/openboard/issues/169)

### Extended Post Details and Full Description

Story Part: 1: Initial Perusal of Posts

When Sage is able to investigate a particular post, ze is able to see extended details (full description in the post).

Blocking Issues

- Add "Kind" message enums [#163](https://github.com/OpenEugene/openboard/issues/163)
- Add Categories message to posts.proto [#165](https://github.com/OpenEugene/openboard/issues/165)
- Add "Categories" message to other messages [#166](https://github.com/OpenEugene/openboard/issues/166)
- Expand FndPostsReq [#167](https://github.com/OpenEugene/openboard/issues/167)
- Add "search by user id" to FndUserReq [#169](https://github.com/OpenEugene/openboard/issues/169)

### Expanding on Categories

Story Part: 1: Initial Perusal of Posts

When Sage looks into the categories, ze sees all the available categories, for example,

- jobs
- collaboration
- events
- services
- spaces
- projects

Blocking Issues

- Add Categories message to posts.proto [#165](https://github.com/OpenEugene/openboard/issues/165)
- Add "Categories" message to other messages [#166](https://github.com/OpenEugene/openboard/issues/166)

### Filter Posts By Kind

Story Part 2: Filtering Posts

Sage is able to further reduce data by post kind (ask vs. offer).

Blocking Issues

- Add "Kind" message enums [#163](https://github.com/OpenEugene/openboard/issues/163)

Frontend Issue

- Be able to find posts by kind (frontend task for MVP).

## Stories not Slated for MVP

### User Authentication

Story Part: 1: Initial Perusal of Posts

Sage does not need to provide any credentials at this initial point.

### Message and Comment Counts

Story Part: 1: Initial Perusal of Posts

Count of messages, count of comments, and count of members.

### Location, Likes, Sum of Messagesa nd Comments

Story Part: 1: Initial Perusal of Posts

Display city and state, count of likes (heart), and sum of messages and comments (or nothing if sum is zero).

### View Count, Tags, Comments, Private Message Count

Story Part: 1: Initial Perusal of Posts

There is now a view count, tags, comments, and number of private messages.

### Tags

Story Part 2: Filtering Posts

When Sage views homepage post data, ze is able to see tags.

### Filter Posts by Tag(s)

Story Part 2: Filtering Posts

Sage is able to look up posts related to tag values.

### Create an Ask Post

Story Part 3: Creating Posts

When creating an ASK (a post kind), Sage inputs as a brief summary (title), "What computer science meetups are happening in this area?", a full description of the request, categories.

### Post Expiration Choice, Tags, Request Location

Story Part 3: Creating Posts

Sage is given an option as to when a created post expires (e.g. urgent, standard), tags, Location of request.

### Automatic Post Expiration

Story Part 3: Creating Posts

Sage automatically lets the post expire on schedule.

## Backend Tests

All new issues and features are expected to come with e2e and/or unit tests.
However, for those features that do not have such tests, the following issues have been created.

### Test

Issue:

### Test

Issue:
