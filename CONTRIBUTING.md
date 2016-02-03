# Contributing

## Prereqs

### Service Prereqs

We use a number of standard services to support our development. Ensure you have access to each of these services before proceeding. If you don't have access, contact [Nick Bartlett](http://faces.tap.ibm.com/bluepages/profile.html?email=ntbartle@us.ibm.com):

- [Github](https://github.ibm.com/digital-marketplace/nautilus) for our remote Git repo and integrated code reviews
- [Travis](https://travis.innovate.ibm.com/digital-marketplace/nautilus) for build and deployment automation
- [Team Concert](https://rtp-rtc9.tivlab.raleigh.ibm.com:9443/jazz/web/projects/Marketplace#action=com.ibm.team.dashboard.viewDashboard&tab=_3) for feature and defect tracking

### Software Prereqs

To improve productivity, we want our developers' dev environments to be as standard as possible.
N.B. If you are using Windows YOU MUST READ AND FOLLOW THE WINDOWS REQUIREMENTS SECTION BELOW. For Mac users see the Mac Tips section below on how to install prereqs. (When using git-review a username is required, do not use the full IBM intranet id instead use the id prefix. i.e. jbskeen not jbskeen@us.ibm.com)

- [Git](http://git-scm.com): Distributed version control system
- [Node](http://nodejs.org/): Runtime environment for nautilus
- [Ruby](http://ruby-lang.org): Programming language for nautilus that is only needed for developers

#### Mac Tips

We recommend using the standard Mac installers for [Git](http://git-scm.com/download/mac).

Install git-review via pip:

- `sudo easy_install pip`
- `sudo pip install git-review`

We also recommend installing [Homebrew](http://brew.sh/) and using it to install node and bash-completion.  If you installed Git using the standard Mac installer for xCode, you can re-install git using Homebrew to make bash-completion work:

- `brew install git`
- `brew install bash-completion`

Node v4.2.3 is the latest version that the project supports right now. You can install Node using Homebrew:

- `brew install homebrew/versions/node4-lts`

Or Downloading Node from [nodejs.org](https://nodejs.org/en/download/)

You also need to edit your `~/.bash_profile` file to use [bash-completion](https://github.com/bobthecow/git-flow-completion/wiki/Install-Bash-git-completion):

```
if [ -f `brew --prefix`/etc/bash_completion ]; then
    . `brew --prefix`/etc/bash_completion
fi
```

#### Linux Tips

Setup on IBM Open Client for Linux (RHEL 6.5) is painful, not recommended, and as of 2/9/15 impossible without installing new c libraries and gcc from scratch. Just don't do it!

If you are using Ubuntu's standard software repositories, be aware that installing Node via apt-get will likely get you an old and unworkable version. More details [here](https://github.com/joyent/node/wiki/Installing-Node.js-via-package-manager).  The commands below install Node from Chris Lea's PPA instead to get a later version.

On Ubuntu 14, the following commands are usually sufficient:
- `sudo apt-get install curl git git-review`
- `curl -sL https://deb.nodesource.com/setup | sudo bash -`
- `sudo apt-get update`
- `sudo apt-get install -y --force-yes nodejs`
- `sudo apt-get install ruby`
- `sudo apt-get install build-essential`
- `sudo npm install -g grunt`.
- `sudo npm install -g phantomjs`.
- `npm install karma-phantomjs-launcher`.

## Running Nautilus on Unix OS's (i.e. Mac OS X or any Linux distribution)

- After the installation of all prereqs, go to the repository root (the nautilus folder from Git)
and run `npm install` to set up the app.
- Check that you have Ruby installed (`ruby -v`). Then run `sudo gem install bundler` to install the ruby gem for the command `bundle install`.
- Again in the repository root, run `npm start`
- Browse to [localhost:20000/](localhost:20000/)
- To stop the server hit `ctrl-c` on your keyboard
- To run tests, run `npm test`.  This will validate html, javascript and SASS files.

## Review Guidelines

All changes, including changes in the staging folder, must meet the following rules:

- All changes must be in English
- No blank documents (example. an empty article)
- No grammatical errors like:
    - Incorrect punctuation and capitalization
    - Using nouns as adjectives
- No spelling mistakes
- No poor quality images
    - Pixelation
    - Resizing artifacts
    - Overlapping text
    - Too small
    - Image backgrounds should be transparent and readable/viewable on any color page background
- All images that have text require globalized images (Pre-GA staging excluded)
- Images should not have padding clear or white space. This means that the content of the image should extend to the borders of the image, and not have any extra space surrounding it. If extra space is necessary, it will be added in the page (not in the image provided) - inconsistent space around images makes page layout and visual flow much more difficult.

## How to Contribute

Our project follows the test-driven development philosophy. This means that when you code a new feature or fix, you start with the test case. All submitted code must have 100% test coverage. Your fix will not get committed to the main stream until you have both a green build and a passed code review.

If you're not already familiar with Git, take a look at the first few chapters of [the online Git book](http://git-scm.com/book). The goal for your first change is to:
1. Improve any of our doc files.
2. Add your name and email to our [authors list](../package.json).
3. Commit the change.
Before beginning, follow the instructions [here](https://github.ibm.com/digital-marketplace/documentation/blob/master/git-process.md) to setup SSH keys, fork the code and get the code pulled down to your local machine.

Your changes must pass all automated tests and a code review before they’re eligible for deployment. To submit your changes for review:
- Ensure that you can access http://github.ibm.com/

In order to commit changes please follow the [Workflow for first code changes](https://github.ibm.com/digital-marketplace/documentation/blob/master/git-process.md#workflow-for-first-code-changes)

You will need to tell git to ignore certain local files on your computer. To see which files that git is currently reading but should actually ignore, check  `git status`. Add files that you want git to ignore to ~/.gitignore_global and then enter:

- `git config --global core.excludesfile ~/.gitignore_global`

At this point, you've cloned the Git repo to your local file system. Grab the latest and create a new branch for your change. By using branches, you can have multiple fixes in flight. Steps to grab the latest changes can be found [here](https://github.ibm.com/digital-marketplace/documentation/blob/master/git-process.md#picking-up-later-changes-from-the-main-project)

- `git checkout -b mybranch`

Go to the root repository directory and open up README.md in your favorite editor. Make a change to it and add it to your staging area. Make frequent use of the status command to understand what stage your files are in. The status command also gives you the commands to move your files from one stage to another.

- `git add README.md`
- `git status`

Once you have your changes ready commit, create a pull request and a test build will kickoff in Travis CI. Note that this only commits them to `mybranch`. There is no way to break the main build so don't worry if you mess up.

- `git commit`
- `git push origin mybranch` <-- This will push your changes to your fork.

Log in to https://github.ibm.com and go to your forked repository to see your pushed branch. You will see a green button, "Compare & pull request", press it. This will create a pull request on the nautilus project.

Let's assume there's another problem in the README. Make another edit and check the status of the file. When you don't get it right the first time, amend your existing changeset and kick off another test build.

- `git commit -a --amend`
- `git push -f origin mybranch`

Take a look at the diff between the current and previous version of the file. Adjust the number of carrots to be equal to the number of revisions you want to go back.

- `git diff HEAD^ README.md`

Iterate using the amend-and-review process until you're satisfied. Go back to Github and refresh. Wait for your build to finish. Once one of the +2 developers approves your change they will merge your code through Github, you don't have to merge it yourself.

Back to our command prompt, let's look at the commit history:

- `git log`

Your fix should be at the top. Now, let's assume that other folks have pushed changes and you'd like to catch up your local development branch. Just make sure you follow these [steps](https://github.ibm.com/digital-marketplace/documentation/blob/master/git-process.md#picking-up-later-changes-from-the-main-project)

This example is unusual in that there is no test case associated with the README. Under normal circumstances, you'll want to ensure your change set includes both the fix and tests. There are two types of tests: standard and integration. Integration tests actually deploy the application to Cloud Foundry and use a production DB. Check with your reviewer for the appropriate place to put your test case if you're not sure. Finally, check the Travis build log to make sure your test case ran and passed. Navigate to your build from GHE or login to the Travis server with your build job ID. [Travis builds](https://travis.innovate.ibm.com/digital-marketplace/nautilus/builds/).

## Starting the development server

    $ git clone git@github.ibm.com:digital-marketplace/nautilus.git
    $ cd nautilus
    $ npm install
    $ npm start

## Generating PDP's from database

The following section illustrates the steps that you will need to follow to generate PDP's. Before running your app, you can do this in console to start up in production mode,

```
$ export NODE_ENV=production
```

The [link](https://github.ibm.com/digital-marketplace/pitcher#prerequisites) describes the prerequisites that are necessary to run pitcher on a cloudant backend.

Whenever you want to generate the static files, follow these steps:

1. Setup the appropriate environment variables

```
$ export GENERATE_SITE=true
```

Next, setup the nautilus project source files:

```
$ npm install
```

2. The PDPs are now being generated in `postinstall`. Thereafter you can start up the server:

```
$ npm start
```

Generating the site by reducing production dependencies helps improve the performance of our project. This was a change incorporated in the build process to speed up the application.

## Automated test suite

We recommend that you run the automated test suite locally prior to submitting your code for review. To do so, from the nautilus directory:

    $ npm test

## Errors and Messages:
- Exception: Could not connect to Github at git@github.ibm.com:digital-marketplace/nautilus.git

Ensure you have generated SSH keys and added them to https://github.ibm.com. [Generating SSH keys](https://help.github.com/enterprise/2.4/user/articles/generating-ssh-keys/)
