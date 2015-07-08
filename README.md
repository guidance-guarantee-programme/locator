# Locator

[![Build Status](https://travis-ci.org/guidance-guarantee-programme/locator.svg)](https://travis-ci.org/guidance-guarantee-programme/locator)

Locate branches of the Citizens Advice service providing [Pension Wise] face to face guidance appointments.


## Prerequisites

* [Go]
* [Git]
* [A configured GOPATH](https://github.com/golang/go/wiki/GOPATH)


## Installation

Clone the repository:

```sh
$ cd $GOPATH/src
$ git clone https://github.com/guidance-guarantee-programme/locator.git
```

Install godep:

```sh
go get github.com/tools/godep
```

## Usage

To start the application:

```sh
$ godep go build .
$ ./locator
```

## Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)


## Contributing

Please see the [contributing guidelines](/CONTRIBUTING.md).

[git]: http://git-scm.com
[go]: https://golang.org
[pension wise]: https://www.pensionwise.gov.uk
