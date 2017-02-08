package jodaTime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var date, _ = time.Parse("2006-01-02 15:04:05 MST", "2007-02-03 16:05:06 CET")

func TestFormat(t *testing.T) {

	tests := []struct {
		format   string
		expected string
	}{

		{"Y", "2007"},
		{"YY", "07"},
		{"YYY", "2007"},
		{"YYYY", "2007"},
		{"y", "2007"},
		{"yy", "07"},
		{"yyy", "2007"},
		{"yyyy", "2007"},

		{"D", "34"},
		{"DD", "34"},

		{"M", "2"},
		{"MM", "02"},
		{"MMM", "Feb"},
		{"MMMM", "February"},

		{"d", "3"},
		{"dd", "03"},

		{"e", "6"},
		{"ee", "06"},

		{"E", "Sat"},
		{"EE", "Saturday"},

		{"h", "5"},
		{"hh", "05"},

		{"H", "16"},
		{"HH", "16"},

		{"a", "PM"},

		{"m", "5"},
		{"mm", "05"},

		{"s", "6"},
		{"ss", "06"},

		{"S", "0"},
		{"SS", "00"},
		{"SSS", "000"},

		{"z", "CET"},

		{"Z", "+0100"},
		{"ZZ", "+01:00"},
		{"ZZZ", "Europe/Paris"},

		{"G", "AD"},

		{"C", "20"},

		{"K", "4"},
		{"KK", "04"},

		{"k", "17"},
		{"kk", "17"},

		{"YYYY.MM.dd", "2007.02.03"},
		{"YYYY.MM.d", "2007.02.3"},
		{"YYYY.M.d", "2007.2.3"},
		{"YY.M.d", "07.2.3"},

		{"dd MMM YYYY", "03 Feb 2007"},
		{"dd MMMM YYYY", "03 February 2007"},

		{"E dd MMMM YYYY", "Sat 03 February 2007"},
		{"EE dd MMMM YYYY", "Saturday 03 February 2007"},

		{"HH:mm:ss", "16:05:06"},
		{"HH:mm:s", "16:05:6"},
		{"HH:m:s", "16:5:6"},
		{"H:m:s", "16:5:6"},

		{"hh:m:s", "05:5:6"},
		{"h:m:s", "5:5:6"},

		{"HH:mm:ss.SSS", "16:05:06.000"},
		{"HH:mm:ss.SS", "16:05:06.00"},
		{"HH:mm:ss.S", "16:05:06.0"},

		{"dd/MM/YYYY HH:mm:ss z", "03/02/2007 16:05:06 CET"},

		{"dd/MM/YYYY HH:mm:ss z ZZ", "03/02/2007 16:05:06 CET +01:00"},
		{"dd/MM/YYYY HH:mm:ss a 世 z Z", "03/02/2007 16:05:06 PM 世 CET +0100"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Format(test.format, date), "("+test.format+") they should be equal")
	}

}
