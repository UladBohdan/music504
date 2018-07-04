package main

import "time"

// Whenever the method is called, it treats UTC's 'yesterday' as the last day
// of the range. I.e. to make it work naturally, the method should be called
// on Friday after 12pm.

func timeframeToTime(period timeframe) (from, to time.Time) {
	now := time.Now().In(time.UTC)
	to = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	switch period {
	case oneWeek:
		from = to.AddDate(0, 0, -7)
	case twoWeeks:
		from = to.AddDate(0, 0, -14)
	case oneMonth:
		from = to.AddDate(0, -1, 0)
	case threeMonths:
		from = to.AddDate(0, -3, 0)
	case oneYear:
		from = to.AddDate(-1, 0, 0)
	}
	return
}

func parseTimePeriod(period string) (parsed timeframe) {
	switch period {
	case "oneWeek":
		parsed = oneWeek
	case "twoWeeks":
		parsed = twoWeeks
	case "oneMonth":
		parsed = oneMonth
	case "threeMonths":
		parsed = threeMonths
	case "oneYear":
		parsed = oneYear
	}
	return
}
