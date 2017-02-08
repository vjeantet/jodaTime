[![GoDoc](https://godoc.org/github.com/vjeantet/jodaTime?status.svg)](https://godoc.org/github.com/vjeantet/jodaTime)
[![Build Status](https://travis-ci.org/vjeantet/jodaTime.svg)](https://travis-ci.org/vjeantet/jodaTime)
[![Coverage Status](https://coveralls.io/repos/vjeantet/jodaTime/badge.png?branch=master)](https://coveralls.io/r/vjeantet/jodaTime?branch=master)
[![Go Report Card](http://goreportcard.com/badge/vjeantet/jodaTime)](http:/goreportcard.com/report/vjeantet/jodaTime)

# JodaTime
Format golang time.Time with joda layout

reference : [http://joda-time.sourceforge.net/apidocs/org/joda/time/format/DateTimeFormat.html](http://joda-time.sourceforge.net/apidocs/org/joda/time/format/DateTimeFormat.html)

# Usage
```go
package main

import (
	"fmt"
	"time"

	"github.com/vjeantet/jodaTime"
)

func main() {
	date := jodaTime.Format("YYYY.MM.dd", time.Now())
	fmt.Println(date)
}
```