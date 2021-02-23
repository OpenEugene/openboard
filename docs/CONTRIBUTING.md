## Welcome

### Please & Thank You

First off, thank you for considering contributing to Openboard. It's people
like you that make projects like this possible!

Following this guide helps to communicate that you respect the time of the 
developers managing and developing this open source project. In return, they 
should reciprocate that respect in addressing your issue, assessing changes, 
and helping you finalize your pull requests.

## Accountability

Openboard is a project utilizing and teaching [Go](https://golang.org) and 
[Elm](https://elm-lang.org). This project is part of the 
[OpenEugene](http://openeugene.org) collective, and our meetings are part of 
the [EugeneTech](https://eugenetech.org) community. Therefore, we align 
ourselves with the codes of conduct put forth by those orgs/communities.

- OpenEugene Code of Conduct
- EugeneTech Code of Conduct
- [Go Community Code of Conduct](https://golang.org/conduct)
- [Elm Community Spaces Code of Conduct](https://github.com/elm-community/discussions/blob/master/code-of-conduct.md)

## Getting Started (Tools and Building)

- [Openboard Readme](../README.md)

## Contributing

### General

We would appreciate that most contributions to this project be made by attending
our sprint wrap & plan and/or check-ins. See the
[README](../README.md) for more
information. Contributing from forks has not yet been tested regarding our
automated tracking mechanisms.

### Reporting

- "Issues" contains requested corrections and improvements.
- "Pull Requests" contains mutations that each directly resolve "Issues"

### Tracking

#### [Issue Qualification](https://github.com/OpenEugene/openboard/projects/3)

When an Issue is:
- up for consideration, it is added to the "Proposed" column
- accepted to be resolved, it is moved to the "Accepted" column
- actively being resolved, it is moved to the "Claimed" column
- fully resolved, it is moved to the "Resolved" column

*Most of these steps are automated. The only step that is often performed
manually is the first step where an issue is added to the "Proposed" column.

#### [Change Management](https://github.com/OpenEugene/openboard/projects/1)

When a PR is:
- created, it is added to the "Initiated" column
- ready for review, it is added to the "Submitted" column
- closed, it is added to the "Finished" column

*All of these steps are automated.

#### Automation Notes

PR Creation, Issue Association, and Tracking:
A Branch that has the suffix "-{issue_number}" (e.g. "my_branch-123") and is
pushed to the project will automatically have a PR opened that is associated 
with the relevant Issue. The PR title is derived from the Issue title. The PR is
automatically added to the "Change Management" project.

WIP Convenience:
A Branch that contains a commit message that has the prefix "WIP" will have its
PR title prefixed with "WIP: " (e.g. "WIP: Fix some bug"). A Branch that
contains a commit message that has the prefix "NOWIP" and also has an existing
PR tracking it will cause the PR title "WIP: " prefix to be trimmed.
