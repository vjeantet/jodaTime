# JodaTime
Format golang time.Time with joda layout

reference : [http://joda-time.sourceforge.net/apidocs/org/joda/time/format/DateTimeFormat.html](http://joda-time.sourceforge.net/apidocs/org/joda/time/format/DateTimeFormat.html)

# Usage
```go
	package main
	import (
		"vjeantet/jodaTime"
		"fmt"
	)

	func main() {
		date := jodaTime.Format("YYYY.MM.dd", time.Now())
		fmt.Println(date)
	}
```