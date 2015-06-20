# Gosnow [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/kujenga/gosnow) [![Build Status](https://travis-ci.org/kujenga/gosnow.svg?branch=master)](https://travis-ci.org/kujenga/gosnow) [![Coverage Status](https://coveralls.io/repos/kujenga/gosnow/badge.svg)](https://coveralls.io/r/kujenga/gosnow)

A Go library for handling the [API Blueprint](https://apiblueprint.org) format, wrapping [Drafter](https://github.com/apiaryio/drafter) and [Snow Crash](https://github.com/apiaryio/snowcrash).

Influenced in principle by [RedSnow](https://github.com/apiaryio/redsnow), the Ruby binding for [Snow Crash](https://github.com/apiaryio/snowcrash).

## Setup

Setup the inner drafter directory with: 
`git submodule update --init --recursive`

Install the drafter dylib with `make install`

Run the pure C tests with `make test`

## Issues

linking to the `libdrafter.dylib` is currently done by linking the library to the global scope in `/usr/local/lib/`. It would be much preferred to have the dylib found locally.

This is the runtime error that occurs when the global library is not present
```
dyld: Library not loaded: /usr/local/lib/libdrafter.dylib
  Referenced from: /Users/uname/dev/gosnow/./test
  Reason: image not found
```
