package jodaTime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var date, _ = time.Parse("2006-01-02 15:04:05.9999999 MST", "2007-02-03 16:05:06.1234567 UTC")

func TestFormatMoreQuotes(t *testing.T) {
	tests := []struct {
		format   string
		expected string
	}{

		{"HHTmm", "16T05"},
		{"HH''mm", "16'05"},
		{"HH''mm", "16'05"},
		{"HH'H'mm", "16H05"},
		{"HH'H''mm", "16Hmm"},
		{"HH'H'''mm", "16H'05"},
		{"HH'''H'mm", "16'H05"},
		{"HH'''H'''mm", "16'H'05"},

		{"HH'H D'mm", "16H D05"},
		{"HH'H D''mm", "16H Dmm"},
		{"HH'H D'''mm", "16H D'05"},
		{"HH'''H D'mm", "16'H D05"},
		{"HH'''H D'''mm", "16'H D'05"},

		{"'With Emoji' HH😀mm", "With Emoji 16😀05"},
		{"'With Emoji' HH'😀'mm", "With Emoji 16😀05"},
		{"'With Emoji' HH'😀🕐'mm", "With Emoji 16😀🕐05"},

		{"HH 'hours and' mm 'minutes'", "16 hours and 05 minutes"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Format(test.format, date), "("+test.format+") they should be equal")
	}
}

func TestFormatMore(t *testing.T) {
	tests := []struct {
		format   string
		expected string
		date     time.Time
	}{
		{"a", "AM",
			time.Date(2007, time.February, 3, 11, 10, 5, 4, time.UTC)},
		{"dd/MM/YYYY HH:mm:ss ZZ", "03/02/2007 16:05:06 -01:00",
			time.Date(2007, time.February, 3, 16, 05, 6, 0, time.FixedZone("", -1*3600))},
		{"yyyy-MM-dd'T'HH:mm:ss.SSSZ", "2007-02-03T16:05:06.000-0100",
			time.Date(2007, time.February, 3, 16, 05, 6, 0, time.FixedZone("", -1*3600))},
		// ZeroFixes
		{"DD", "05", time.Date(2007, time.January, 5, 16, 05, 6, 0, time.FixedZone("", -1*3600))},
		{"ww", "14", time.Date(2007, time.April, 5, 16, 05, 6, 0, time.FixedZone("", -1*3600))},
		{"dd", "15", time.Date(2007, time.April, 15, 16, 05, 6, 0, time.FixedZone("", -1*3600))},
		{"MM", "11", time.Date(2007, time.November, 15, 16, 05, 6, 0, time.FixedZone("", -1*3600))},
		{"ee", "04", time.Date(2007, time.November, 15, 16, 05, 6, 0, time.FixedZone("", -1*3600))},
		{"hh", "11", time.Date(2007, time.November, 15, 11, 05, 6, 0, time.FixedZone("", -1*3600))},
		{"hh", "09", time.Date(2007, time.November, 15, 20, 05, 6, 0, time.FixedZone("", -1*3600))},
		{"HH", "04", time.Date(2007, time.November, 15, 04, 05, 6, 0, time.FixedZone("", -1*3600))},
		{"mm", "05", time.Date(2007, time.November, 15, 04, 05, 6, 0, time.FixedZone("", -1*3600))},
		{"mm", "55", time.Date(2007, time.November, 15, 04, 55, 6, 0, time.FixedZone("", -1*3600))},
		{"ss", "46", time.Date(2007, time.November, 15, 04, 55, 46, 0, time.FixedZone("", -1*3600))},

		{"SS", "00", time.Date(2007, time.November, 15, 04, 55, 46, 9, time.FixedZone("", -1*3600))},
		{"SSS", "000", time.Date(2007, time.November, 15, 04, 55, 46, 9, time.FixedZone("", -1*3600))},
		{"SSS", "100", time.Date(2007, time.November, 15, 04, 55, 46, 100000000, time.FixedZone("", -1*3600))},
		{"SSS", "010", time.Date(2007, time.November, 15, 04, 55, 46, 10000000, time.FixedZone("", -1*3600))},
		{"SSS", "001", time.Date(2007, time.November, 15, 04, 55, 46, 1000000, time.FixedZone("", -1*3600))},

		{"Z", "-1100", time.Date(2007, time.November, 15, 04, 55, 46, 1000000, time.FixedZone("", -11*3600))},
		{"ZZ", "-11:00", time.Date(2007, time.November, 15, 04, 55, 46, 1000000, time.FixedZone("", -11*3600))},
		{"KK", "11", time.Date(2007, time.November, 15, 23, 55, 46, 1000000, time.FixedZone("", -11*3600))},
		{"kk", "09", time.Date(2007, time.November, 15, 8, 55, 46, 1000000, time.FixedZone("", -11*3600))},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Format(test.format, test.date), "("+test.format+") they should be equal")
	}

}
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

		{"D ", "34 "},
		{"DD", "34"},

		{"M", "2"},
		{"MM", "02"},
		{"MMM", "Feb"},
		{"MMMM", "February"},

		{"d", "3"},
		{"dd", "03"},

		{"e", "6"},
		{"ee", "06"},
		{"eeL", "06L"},

		{"E", "Sat"},
		{"EE", "Sat"},
		{"EEE", "Sat"},
		{"EEEE", "Saturday"},

		{"h", "5"},
		{"hh", "05"},

		{"H", "16"},
		{"HH", "16"},

		{"a", "PM"},

		{"m", "5"},
		{"mm", "05"},

		{"s", "6"},
		{"ss", "06"},

		{"S", "1"},
		{"SS", "12"},
		{"SSSL", "123L"},

		{"zL", "UTCL"},

		{"ZL", "+0000L"},
		{"ZZ", "+00:00"},
		{"ZZZ", ""},

		{"G", "AD"},

		{"C", "20"},

		{"K", "4"},
		{"KKL", "04L"},

		{"kL", "17L"},
		{"kk k", "17 17"},

		{"w", "5"},
		{"ww", "05"},
		{"wwT", "05T"},

		{"YYYY.MM.dd", "2007.02.03"},
		{"YYYY.MM.d", "2007.02.3"},
		{"YYYY.M.d", "2007.2.3"},
		{"YY.M.d", "07.2.3"},

		{"dd MMM YYYY", "03 Feb 2007"},
		{"dd MMMM YYYY", "03 February 2007"},

		{"E dd MMMM YYYY", "Sat 03 February 2007"},
		{"EE dd MMMM YYYY", "Sat 03 February 2007"},
		{"EEE dd MMMM YYYY", "Sat 03 February 2007"},
		{"EEEE dd MMMM YYYY", "Saturday 03 February 2007"},

		{"HH:mm:ss", "16:05:06"},
		{"HH:mm:s", "16:05:6"},
		{"HH:m:s", "16:5:6"},
		{"H:m:s", "16:5:6"},

		{"hh:m:s", "05:5:6"},
		{"h:m:s", "5:5:6"},

		{"HH:mm:ss.SSS", "16:05:06.123"},
		{"HH:mm:ss.SS", "16:05:06.12"},
		{"HH:mm:ss.S", "16:05:06.1"},

		{"dd/MM/YYYY HH:mm:ss z", "03/02/2007 16:05:06 UTC"},

		{"dd/MM/YYYY HH:mm:ss z ZZ", "03/02/2007 16:05:06 UTC +00:00"},
		{"dd/MM/YYYY HH:mm:ss a 世 z Z", "03/02/2007 16:05:06 PM 世 UTC +0000"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Format(test.format, date), "("+test.format+") they should be equal")
	}

}

func BenchmarkFormat(b *testing.B) {
	// run the Parse function b.N times
	for n := 0; n < b.N; n++ {
		Format("dd/MM/YYYY HH:mm:ss z", date)
	}
}
