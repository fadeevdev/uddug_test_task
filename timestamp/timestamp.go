package timestamp

import "time"

func TruncateMinutes(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(),
		t.Day(), t.Hour(), 0, 0, 0, t.Location())
}

func TruncateHours(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(),
		t.Day(), 0, 0, 0, 0, t.Location())
}

func TruncateWeek(t time.Time) time.Time {
	year, week := t.ISOWeek()
	timeBenchmark := time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)
	weekStartBenchmark := timeBenchmark.AddDate(0, 0, -(int(timeBenchmark.Weekday())+6)%7)

	_, weekBenchmark := weekStartBenchmark.ISOWeek()

	return weekStartBenchmark.AddDate(0, 0, (week-weekBenchmark)*7)
}

func TruncateDays(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(),
		1, 0, 0, 0, 0, t.Location())
}
