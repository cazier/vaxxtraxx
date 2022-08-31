package main

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	var date time.Time = time.Date(2020, 2, 1, 0, 0, 0, 0, time.UTC)
	var date_string string = "2020-02-01"

	resp, err := get(date_string)

	if err != nil {
		t.Error(err)
		return
	}

	if resp != date {
		t.Errorf("Date was not parsed. %s became %s instead of %s", date_string, resp, date)
		return
	}
}

func TestGetDueDates_months(t *testing.T) {
	var date time.Time = get_safe("2020-02-26")
	resp, err := get_due_dates(date, "months", 6)
	expected := []time.Time{
		get_safe("2020-08-26"),
		get_safe("2021-02-26"),
		get_safe("2021-08-26"),
	}

	if err != nil {
		t.Error(err)
		return
	}
	for i := 0; i < len(expected); i++ {
		if resp[i] != expected[i] {
			t.Errorf("At index %d, resp = %s, expected = %s", i, resp[i], expected[i])
		}
	}
}

func TestGetDueDates_leap_years(t *testing.T) {
	t.Skip("To maintain equality with python, leap years are just going to be ignored.")
	var date time.Time = get_safe("2020-02-29")
	resp, err := get_due_dates(date, "years", 1, 4)
	expected := []time.Time{
		get_safe("2021-03-01"),
		get_safe("2022-03-01"),
		get_safe("2023-03-01"),
		get_safe("2024-02-29"),
	}

	if err != nil {
		t.Error(err)
		return
	}
	for i := 0; i < len(expected); i++ {
		if resp[i] != expected[i] {
			t.Errorf("At index %d, resp = %s, expected = %s", i, resp[i], expected[i])
		}
	}
}

func TestGetDueDates_leap_days(t *testing.T) {
	var date time.Time = get_safe("2020-02-26")
	resp, err := get_due_dates(date, "days", 1, 5)
	expected := []time.Time{
		get_safe("2020-02-27"),
		get_safe("2020-02-28"),
		get_safe("2020-02-29"),
		get_safe("2020-03-01"),
		get_safe("2020-03-02"),
	}

	if err != nil {
		t.Error(err)
		return
	}
	for i := 0; i < len(expected); i++ {
		if resp[i] != expected[i] {
			t.Errorf("At index %d, resp = %s, expected = %s", i, resp[i], expected[i])
		}
	}

	date = get_safe("2021-02-26")
	resp, err = get_due_dates(date, "days", 1, 5)
	expected = []time.Time{
		get_safe("2021-02-27"),
		get_safe("2021-02-28"),
		get_safe("2021-03-01"),
		get_safe("2021-03-02"),
		get_safe("2021-03-03"),
	}

	if err != nil {
		t.Error(err)
		return
	}
	for i := 0; i < len(expected); i++ {
		if resp[i] != expected[i] {
			t.Errorf("At index %d, resp = %s, expected = %s", i, resp[i], expected[i])
		}
	}
}

func TestOverdueDate_is_overdue(t *testing.T) {
	var due time.Time = get_safe("2020-02-05")
	var now time.Time = get_safe("2025-02-05")

	if is_overdue(due, now) != true {
		t.Errorf("This is overdue, but was reported not to be")
		return
	}
}

func TestOverdueDate_is_not_overdue(t *testing.T) {
	var due time.Time = get_safe("2030-02-05")
	var now time.Time = get_safe("2025-02-05")

	if is_overdue(due, now) == true {
		t.Errorf("This is NOT overdue, but was reported TO be")
		return
	}
}

func TestOverdueDate_is_equality_overdue(t *testing.T) {
	var date time.Time = get_safe("2030-02-05")

	if is_overdue(date, date) != true {
		t.Errorf("This is equal, but was set to be strict so should be overdue")
		return
	}

	if is_overdue(date, date, false) == true {
		t.Errorf("This is equal, but was NOT set to be strict so should NOT be overdue")
		return
	}
}
