package main

import (
	"fmt"
	"time"
)

type DateTimeError struct {
	Why string
}

func (err *DateTimeError) Error() string {
	return fmt.Sprintf("An error occurred: %s", err.Why)
}

func get(date_string string) (result time.Time, err error) {
	// Note that Go matches against THIS SPECIFIC DATE to determine the parser expectations
	return time.Parse("2006-01-02", date_string)
}

func get_safe(date_string string) time.Time {
	val, _ := get(date_string)
	return val
}

func get_due_dates(start time.Time, rate_span string, rate_frequency int, nb ...int) (result []time.Time, err error) {
	var length int = 3

	if len(nb) != 0 {
		length = nb[0]
	}

	result = make([]time.Time, length)

	for i := 0; i < length; i++ {
		switch rate_span {
		case "days":
			result[i] = start.AddDate(0, 0, rate_frequency*(i+1))
		case "weeks":
			result[i] = start.AddDate(0, 0, 7*rate_frequency*(i+1))
		case "months":
			result[i] = start.AddDate(0, rate_frequency*(i+1), 0)
		case "years":
			result[i] = start.AddDate(rate_frequency*(i+1), 0, 0)
		default:
			return nil, &DateTimeError{"One of 'days', 'weeks', 'months', or 'years' must be supplied."}
		}
	}
	return result, err
}

func is_overdue(due_date time.Time, check_date time.Time, strict ...bool) bool {
	var _strict bool = true

	if len(strict) != 0 {
		_strict = strict[0]
	}

	if due_date.Equal(check_date) {
		return _strict
	} else if due_date.Before(check_date) {
		return true
	} else {
		return false
	}
}
