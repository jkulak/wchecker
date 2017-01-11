# wchecker

A command-line tool for checking your articles for writing consistency.
- use of dashes

## Usage

## Compilation

I like to use Docker 🐳 to work with Golang on my local dev machine.

Run

```
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.7 go run main.go
```

Build

```
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.7 go build -v
```

Work inside the Docker container (you should set the GOPATH env variable)

```
docker run --rm -ti -e "GOPATH=/root/go" -v "$PWD":/root/go -w /root/go golang:1.7 bash
```

## Build, test and doc status

[![Build Status](https://travis-ci.org/jkulak/wchecker.svg?branch=master)](https://travis-ci.org/jkulak/wchecker)
[![Go Report Card](https://goreportcard.com/badge/github.com/jkulak/wchecker)](https://goreportcard.com/report/github.com/jkulak/wchecker)
[![GoDoc](https://godoc.org/github.com/jkulak/wchecker?status.svg)](https://godoc.org/github.com/jkulak/wchecker)
