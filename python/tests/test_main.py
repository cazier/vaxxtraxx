import unittest

import arrow

import main

class TestGettingDueDates(unittest.TestCase):
    def test_months(self):
        date = arrow.get("2020-02-26")
        assert main.get_due_dates(date, "months", 6) == [
            arrow.get("2020-08-26"),
            arrow.get("2021-02-26"),
            arrow.get("2021-08-26"),
        ]

    def test_leap_years(self):
        date = arrow.get("2020-02-29")
        assert main.get_due_dates(date, "years", 1, 4) == [
            arrow.get("2021-02-28"),
            arrow.get("2022-02-28"),
            arrow.get("2023-02-28"),
            arrow.get("2024-02-29"),
        ]

    def test_leap_days(self):
        date = arrow.get("2020-02-26")
        assert main.get_due_dates(date, "days", 1, 5) == [
            arrow.get("2020-02-27"),
            arrow.get("2020-02-28"),
            arrow.get("2020-02-29"),
            arrow.get("2020-03-01"),
            arrow.get("2020-03-02"),
        ]

        date = arrow.get("2021-02-26")
        assert main.get_due_dates(date, "days", 1, 5) == [
            arrow.get("2021-02-27"),
            arrow.get("2021-02-28"),
            arrow.get("2021-03-01"),
            arrow.get("2021-03-02"),
            arrow.get("2021-03-03"),
        ]


class TestOverdueDates(unittest.TestCase):
    def test_is_overdue(self):
        due = arrow.get("2020-02-05")
        now = arrow.get("2025-02-05")

        assert main.is_overdue(due, now) == True

    def test_is_not_overdue(self):
        due = arrow.get("2030-02-05")
        now = arrow.get("2025-02-05")

        assert main.is_overdue(due, now) == False

    def test_is_equality_overdue(self):
        date = arrow.get("2030-02-05")
        assert main.is_overdue(date, date) == True
        assert main.is_overdue(date, date, False) == False
