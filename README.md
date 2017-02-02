# go-swetest

go-swetest is the simple wrapper of swisseph library's swetest.

[![GoDoc](https://godoc.org/github.com/DestinyLab/go-swetest?status.svg)](https://godoc.org/github.com/DestinyLab/go-swetest) [![Go Report Card](https://goreportcard.com/badge/github.com/DestinyLab/go-swetest)](https://goreportcard.com/report/github.com/DestinyLab/go-swetest) [![Build Status](https://travis-ci.org/DestinyLab/go-swetest.svg?branch=master)](https://travis-ci.org/DestinyLab/go-swetest) [![Coverage Status](https://coveralls.io/repos/github/DestinyLab/go-swetest/badge.svg?branch=master)](https://coveralls.io/github/DestinyLab/go-swetest?branch=master)

## Installation

```
go get -u github.com/DestinyLab/go-swetest
```

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/DestinyLab/go-swetest"
)

func main() {
	opt := []string{
		"-b11.11.2017",
		"-ut00:00:00",
		"-p0",
		"-fZ",
		"-eswe",
		"-head",
	}
	s := swetest.New()
	res, err := s.Query(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", res)
}
```
