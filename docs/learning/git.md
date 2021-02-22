# Git Walk-through

## Introduction

- [Programming with Mosh - What is Git?](https://youtu.be/2ReR1YJrNOM)
- [Paul Programming - What is Git?](https://youtu.be/OqmSzXDrJBk)

The first video is "Git in principle", and the second is "Git in practice".

The second video seems to imply that Git is centralized, which is in opposition
to the description given in the first video. git is not dependent on a central
repository, rather, releases are. It is common for projects to be "forked" (i.e.
a copy of the repo is created under another organization/namespace), then
modified, then have selected modifications "pulled" into the original
repository for the improvements to be released publicly.

### Installation

- [Git - Getting Started](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

Side note: It is preferable to use Git version 2.28 or higher since it permits
configuring `init.defaultBranch` which allows for using user-defined branches
as default. 

### Configuration

- [Git - Customizing Git](https://git-scm.com/book/en/v2/Customizing-Git-Git-Configuration) 

The most pertinent information is in the "Git Configuration" section only. Feel
free to read further for more details and advanced configuration. The email and
username you select will show up in publicly available git commit histories.

### Operation

- [Udacity - Git Workflow](https://youtu.be/3a2x1iJFJWc)
- [Kev The Dev - Easy Overview](https://youtu.be/7dYHRI55wxo)

In the Udacity video, four columns are shown. It may help to think
of the column on the far left as the newest changes to the code. We make changes
to code, then pass those changes through certain steps to ensure the difference 
the changes make is correct and understandable.

Kev's overview has some caveats. First, when cloning a project, if you have not
first setup SSH access to your repository, use HTTPS instead (seen at 6:54).
Second, when you see the usage of `git checkout` as it applies to switching
branches, it is preferrable to use the `git switch` command. So, `git checkout
-b new-branch` should be `git switch -c new-branch`, and `git checkout
existing-branch` should be `git switch existing-branch`. The usage of checkout
to move to a specific commit (seen at 4:50) is correct. The git command will
give notices to use `git switch` in the relevant cases and those messages can be
seen in the video. Lastly, it is currently preferred to use the default branch
name "main" rather than "master".

### Interpretation

A simple list of commands to start work on an issue.
```sh
git switch main
git pull origin main
git switch -c me/my_branch-123
# make change(s)
git add .
git commit -m"Make specific change"
git push -u origin me/my_branch-123
```

Afterward, if additional commits are made, only `git push` is necessary since
the upstream will have been set. Also, it may be useful to run `git remote prune
origin` after `git switch main` to ensure that the references to remote branches
are cleaned up.
