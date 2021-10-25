Form3 [![CircleCI](https://circleci.com/gh/ahmedkamals/form3.svg?style=svg)](https://circleci.com/gh/ahmedkamals/form3 "Build Status")
======

[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](LICENSE.md "License")
[![release](https://img.shields.io/github/v/release/ahmedkamals/form3.svg)](https://github.com/ahmedkamals/form3/releases/latest "Release")
[![codecov](https://codecov.io/gh/ahmedkamals/form3/branch/main/graph/badge.svg?token=XPINFB5JYV)](https://codecov.io/gh/ahmedkamals/form3 "Code Coverage")
[![GolangCI](https://golangci.com/badges/github.com/ahmedkamals/form3.svg?style=flat-square)](https://golangci.com/r/github.com/ahmedkamals/form3 "Code Coverage")
[![Go Report Card](https://goreportcard.com/badge/github.com/ahmedkamals/form3)](https://goreportcard.com/report/github.com/ahmedkamals/form3 "Go Report Card")
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/65feb277726f4a10895f028d460f9196)](https://www.codacy.com/manual/ahmedkamals/form3?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=ahmedkamals/form3&amp;utm_campaign=Badge_Grade "Code Quality")
[![GoDoc](https://godoc.org/github.com/ahmedkamals/form3?status.svg)](https://godoc.org/github.com/ahmedkamals/form3 "Documentation")
[![DepShield Badge](https://depshield.sonatype.org/badges/ahmedkamals/form3/depshield.svg)](https://depshield.github.io "DepShield")
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fahmedkamals%2Fform3.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fahmedkamals%2Fform3?ref=badge_shield "Dependencies")

```bash
FFFFFFFFFFFFFFFFFFFFFF                                                           333333333333333
F::::::::::::::::::::F                                                          3:::::::::::::::33
F::::::::::::::::::::F                                                          3::::::33333::::::3
FF::::::FFFFFFFFF::::F                                                          3333333     3:::::3
  F:::::F       FFFFFFooooooooooo   rrrrr   rrrrrrrrr      mmmmmmm    mmmmmmm               3:::::3
  F:::::F           oo:::::::::::oo r::::rrr:::::::::r   mm:::::::m  m:::::::mm             3:::::3
  F::::::FFFFFFFFFFo:::::::::::::::or:::::::::::::::::r m::::::::::mm::::::::::m    33333333:::::3
  F:::::::::::::::Fo:::::ooooo:::::orr::::::rrrrr::::::rm::::::::::::::::::::::m    3:::::::::::3
  F:::::::::::::::Fo::::o     o::::o r:::::r     r:::::rm:::::mmm::::::mmm:::::m    33333333:::::3
  F::::::FFFFFFFFFFo::::o     o::::o r:::::r     rrrrrrrm::::m   m::::m   m::::m            3:::::3
  F:::::F          o::::o     o::::o r:::::r            m::::m   m::::m   m::::m            3:::::3
  F:::::F          o::::o     o::::o r:::::r            m::::m   m::::m   m::::m            3:::::3
FF:::::::FF        o:::::ooooo:::::o r:::::r            m::::m   m::::m   m::::m3333333     3:::::3
F::::::::FF        o:::::::::::::::o r:::::r            m::::m   m::::m   m::::m3::::::33333::::::3
F::::::::FF         oo:::::::::::oo  r:::::r            m::::m   m::::m   m::::m3:::::::::::::::33
FFFFFFFFFFF           ooooooooooo    rrrrrrr            mmmmmm   mmmmmm   mmmmmm 333333333333333
```

Form3 Take Home Exercise - by [Ahmed Kamal][2]

Table of Contents
-----------------

* [🏎️ Getting Started](#-getting-started)

    * [Prerequisites](#prerequisites)
    * [Installation](#installation)
    * [Examples](#examples)

* [🕸️ Tests](#-tests)

    * [⚓ Git Hooks](#-git-hooks)

* [👨‍💻 Credits](#-credits)

* [🆓 License](#-license)

🏎️ Getting Started
------------------

### Prerequisites

* [Golang 1.17 or later][1].

### Installation

```bash
go get -u github.com/ahmedkamals/form3
cp .env.sample .env
```

### Examples

```go
config := Config{
endpoint: os.Getenv("API_ENDPOINT"),
}

apiClient = NewClient(config, &http.Client{})
ctx := context.Background()

account, err := apiClient.CreateAccount(ctx, form3.Account{})

account, err := apiClient.FetchAccount(ctx, uuid.MustParse("ab6d05a8-8d41-4957-833e-fcc42126351b"), 0)

err := apiClient.DeleteAccount(ctx, uuid.MustParse("ab6d05a8-8d41-4957-833e-fcc42126351b"), 0)
```

🕸️ Tests
--------

```bash
docker-compose up --abort-on-container-exit --build --remove-orphans api-client
```

### ⚓ Git Hooks

In order to set up tests running on each commit do the following steps:

```bash
git config --local core.hooksPath .githooks
```

👨‍💻 Credits
----------

* [ahmedkamals][2]

🆓 LICENSE
----------

Form3 is released under MIT license, please refer to
the [`LICENSE.md`](https://github.com/ahmedkamals/form3/blob/main/LICENSE.md "License") file.

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fahmedkamals%2Fform3.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fahmedkamals%2Fform3?ref=badge_large "Dependencies")

Happy Coding 🙂

[![Analytics](http://www.google-analytics.com/__utm.gif?utmwv=4&utmn=869876874&utmac=UA-136526477-1&utmcs=ISO-8859-1&utmhn=github.com&utmdt=form3&utmcn=1&utmr=0&utmp=/ahmedkamals/form3?utm_source=www.github.com&utm_campaign=form3&utm_term=form3&utm_content=form3&utm_medium=repository&utmac=UA-136526477-1)]()

[1]: https://golang.org/dl/ "Download Golang"

[2]: https://github.com/ahmedkamals "Author"
