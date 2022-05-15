package cron

import (
	"testing"
	"time"
)

func TestRepeatCountTimesNext(t *testing.T) {
	tests := []struct {
		time     string
		from     string
		delay    time.Duration
		count    int
		expected string
	}{
		{"Mon Jul 9 14:45 2012", "Mon Jul 9 14:45 2012", 15*time.Minute + 50*time.Nanosecond, 3, "Mon Jul 9 15:00 2012"},
		{"Mon Jul 9 14:59 2012", "Mon Jul 9 14:59 2012", 15 * time.Minute, 3, "Mon Jul 9 15:14:00 2012"},
		{"Mon Jul 9 14:59:59 2012", "Mon Jul 9 14:59:59 2012", 15 * time.Minute, 3, "Mon Jul 9 15:14:59 2012"},

		{"Mon Jul 9 14:59 2012", "Mon Jul 9 14:00 2012", 15 * time.Minute, 3, ""},
		{"Mon Jul 9 14:32:00.104 2012", "Mon Jul 9 14:17 2012", 15 * time.Minute, 2, "Mon Jul 9 14:47:00 2012"},
		{"Mon Jul 9 14:29 2012", "Mon Jul 9 14:04 2012", 15 * time.Minute, 2, "Mon Jul 9 14:34:00 2012"},
		{"Mon Jul 9 14:59 2012", "Mon Jul 9 14:04 2012", 15 * time.Minute, 4, "Mon Jul 9 15:04:00 2012"},
	}

	for _, c := range tests {
		actual := RepeatCountTimesFrom(getTime(c.from), c.delay, c.count).Next(getTime(c.time))
		expected := getTime(c.expected)
		if actual != expected {
			t.Errorf("%s, \"%s\": (expected) %v != %v (actual)", c.time, c.delay, expected, actual)
		}
	}
}
