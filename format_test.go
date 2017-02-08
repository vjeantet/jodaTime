package jodaTime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var date, _ = time.Parse("2006-01-02 15:04:05 MST", "2007-02-03 16:05:06 UTC")

func TestFormat(t *testing.T) {

	tests := []struct {
		format   string
		expected string
	}{
		{"YY", "07"},
		// {"YYYY", "2007"},
		// {"YYYY.MM.dd", "2007.02.03"},
		// {"YYYY.MM.d", "2007.02.3"},
		// {"YYYY.M.d", "2007.2.3"},
		// {"YY.M.d", "07.2.3"},

		// {"dd MMM YYYY", "03 Feb 2007"},
		// {"dd MMMM YYYY", "03 February 2007"},

		// {"E dd MMMM YYYY", "Sat 03 February 2007"},
		// {"EE dd MMMM YYYY", "Saturday 03 February 2007"},

		// {"HH:mm:ss", "16:05:06"},
		// {"HH:mm:s", "16:05:6"},
		// {"HH:m:s", "16:5:6"},
		// {"H:m:s", "16:5:6"},

		// {"hh:m:s", "04:5:6"},
		// {"h:m:s", "4:5:6"},

		// {"HH:mm:ssSSS", "16:05:06.000"},
		// {"HH:mm:ssSS", "16:05:06.00"},
		// {"HH:mm:ssS", "16:05:06.0"},

		// {"dd/MM/YYYY HH:mm:ss z", "03/02/2007 16:05:06 UTC"},

		// {"dd/MM/YYYY HH:mm:ss z ZZ", "03/02/2007 16:05:06 UTC +00:00"},
		// {"dd/MM/YYYY HH:mm:ss a 世 z Z", "03/02/2007 16:05:06 PM 世 UTC +0000"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Format(test.format, date), "they should be equal")
	}

}
