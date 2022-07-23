package transaction

import (
	"errors"
	"sort"
	"time"

	"github.com/fadeevdev/uddug_test_task/timestamp"
)

type Transaction struct {
	Value     int
	Timestamp time.Time
}

type Transactions []*Transaction

const (
	MONTH = "month"
	WEEK  = "week"
	DAY   = "day"
	HOUR  = "hour"
)

const ERROR_WRONG_INTERVAL = "GeneratePlot function got unexpected interval input parameter"

func GeneratePlot(ts Transactions, interval string) (Transactions, error) {
	// if input array lenght < 1 return it with no error
	if len(ts) < 1 {
		return ts, nil
	}

	// sort the array by timestamp descending
	sort.Slice(ts, func(i, j int) bool {
		return ts[i].Timestamp.After(ts[j].Timestamp)
	})

	// truncate timestamps according requested interval or return error on wrong interval
	switch interval {
	case MONTH:
		ts.truncateDays()
	case WEEK:
		ts.truncateWeek()
	case DAY:
		ts.truncateHours()
	case HOUR:
		ts.truncateMinutes()
	default:
		return nil, errors.New(ERROR_WRONG_INTERVAL)
	}

	// if array lenght = 1 return truncated value
	if len(ts) == 1 {
		return ts, nil
	}

	// initialize array for returning the results
	results := make([]*Transaction, 0, len(ts))

	lastArrayElement := len(ts) - 1

	// first element will be always added to results array,
	// as the input array was sorted and first is always latest(biggest timestamp) in it's interval
	results = append(results, ts[0])

	// starting for loop from second element of array
	for i := 1; i <= lastArrayElement; i++ {
		// comparing timestamp with previous element in array, if it is dirrent,
		// then i element was the last from it's interval (array was sorted descending before truncation),
		// adding to result array
		if ts[i].Timestamp != ts[i-1].Timestamp {
			results = append(results, ts[i])
		}
	}

	return results, nil
}

//truncateMinutes truncates min, sec, msec to 0, in each Timestamp field in array of Transaction
func (ts Transactions) truncateMinutes() {
	for i := 0; i < len(ts); i++ {
		ts[i].Timestamp = timestamp.TruncateMinutes(ts[i].Timestamp)
	}
}

//truncateHours truncates hour, min, sec, msec to 0, in each Timestamp field in array of Transaction
func (ts Transactions) truncateHours() {
	for i := 0; i < len(ts); i++ {
		ts[i].Timestamp = timestamp.TruncateHours(ts[i].Timestamp)
	}
}

//truncateWeek truncates each Timestamp field in array of Transaction to the date of Monday for current ISOWeek of Timestamp field
func (ts Transactions) truncateWeek() {
	for i := 0; i < len(ts); i++ {
		ts[i].Timestamp = timestamp.TruncateWeek(ts[i].Timestamp)
	}

}

//truncateDays truncates each Timestamp field in array of Transaction to the first day of current Month of Timestamp field
func (ts Transactions) truncateDays() {
	for i := 0; i < len(ts); i++ {
		ts[i].Timestamp = timestamp.TruncateDays(ts[i].Timestamp)
	}
}
