package jodaTime

/*
jodaTime provides a date formatter using the yoda syntax.
http://joda-time.sourceforge.net/apidocs/org/joda/time/format/DateTimeFormat.html

Why?

because :)

Usage

	package main
	import (
		"vjeantet/jodaTime"
		"fmt"
	)

	func main() {
		date := jodaTime.Format("YYYY.MM.dd", time.Now())
		fmt.Println(date)
	}
*/

import (
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/k0kubun/pp"
)

/*
 Symbol  Meaning                      Presentation  Examples
 ------  -------                      ------------  -------
 G       era                          text          AD
 C       century of era (>=0)         number        20
 Y       year of era (>=0)            year          1996

 x       weekyear                     year          1996
 w       week of weekyear             number        27
 e       day of week                  number        2
 E       day of week                  text          Tuesday; Tue

 y       year                         year          1996
 D       day of year                  number        189
 M       month of year                month         July; Jul; 07
 d       day of month                 number        10

 a       halfday of day               text          PM
 K       hour of halfday (0~11)       number        0
 h       clockhour of halfday (1~12)  number        12

 H       hour of day (0~23)           number        0
 k       clockhour of day (1~24)      number        24
 m       minute of hour               number        30
 s       second of minute             number        55
 S       fraction of second           number        978

 z       time zone                    text          Pacific Standard Time; PST
 Z       time zone offset/id          zone          -0800; -08:00; America/Los_Angeles

 '       escape for text              delimiter
 ''      single quote                 literal       '
*/

// Format formats a date based on Microsoft Excel (TM) conventions
func Format(format string, date time.Time) string {
	lenString := utf8.RuneCountInString(format)
	formatRune := []rune(format)
	out := ""
	for i := 0; i < lenString; i++ {
		pp.Println("i,v-->", i, string(formatRune[i]))
		switch r := formatRune[i]; r {
		case 'Y': // Y YYYY YY
			out += strconv.Itoa(date.Year())
			// while next is Y, combine
			// switch combined and replace

		case 'y': // y yyyy yy
			out += strconv.Itoa(date.Year())

		case 'D': // D DD
			out += strconv.Itoa(date.YearDay())

		case 'w': // w ww
			_, w := date.ISOWeek()
			out += strconv.Itoa(w)

		case 'M': // M MM MMM MMMM
			out += date.Month().String()

		case 'd': // d dd
			out += strconv.Itoa(date.Day())

		case 'e': // e ee
			out += strconv.Itoa(int(date.Weekday()))

		case 'E': // E EE
			out += date.Weekday().String()

		case 'h': // h hh
			out += strconv.Itoa(date.Hour())

		case 'H': // H HH
			out += strconv.Itoa(date.Hour())

		case 'a': // a
			if date.Hour() > 12 {
				out += "AM"
			} else {
				out += "PM"
			}

		case 'm': // m mm
			out += strconv.Itoa(date.Minute())

		case 's': // s ss
			out += strconv.Itoa(date.Second())

		case 'S': // S SS SSS
			out += strconv.Itoa(date.Nanosecond())

		case 'z': // z
			z, _ := date.Zone()
			out += z

		case 'Z': // Z ZZ
			_, z := date.Zone()
			out += strconv.Itoa(z)

		case '\'': // ' (text delimiter)  or '' (real quote)

		case 'G':
		case 'C':
		case 'x': // x xxxx xx
		case 'K': // K KK
		case 'k': // k kk
		default:
			out += string(r)
		}
	}
	return out

}

var longDayNames = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

var shortDayNames = []string{
	"Sun",
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
}

var shortMonthNames = []string{
	"---",
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

var longMonthNames = []string{
	"---",
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}
