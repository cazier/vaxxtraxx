use chrono::{Duration, Months, NaiveDate, ParseError};
use std::vec;

fn main() {
    let out = get_due_dates(NaiveDate::from_ymd(2020, 02, 20), "months", 4, 3);
    println!("My dates are {:?}", out);
}

pub fn get(date_string: &str) -> Result<NaiveDate, ParseError> {
    return NaiveDate::parse_from_str(date_string, "%Y-%m-%d");
}

pub fn get_safe(date_string: &str) -> NaiveDate {
    return get(date_string).ok().unwrap();
}

pub fn get_due_dates(
    start: NaiveDate,
    rate_span: &str,
    rate_frequency: i64,
    nb: u8,
) -> Vec<NaiveDate> {
    let mut result: Vec<NaiveDate> = vec![];

    for i in 1..(nb + 1) {
        match rate_span {
            "days" => result.push(start + Duration::days(i as i64 * rate_frequency)),
            "weeks" => result.push(start + Duration::weeks(i as i64 * rate_frequency)),
            "months" => result.push(start + Months::new(i as u32 * rate_frequency as u32)),
            "years" => result.push(start + Months::new(i as u32 * rate_frequency as u32 * 12)),
            _ => panic!("One of 'days', 'weeks', 'months', or 'years' must be supplied."),
        }
    }
    return result;
}

pub fn is_overdue(due_date: NaiveDate, check_date: NaiveDate, strict: bool) -> bool {
    if due_date == check_date {
        return strict;
    } else if due_date < check_date {
        return true;
    } else {
        return false;
    }
}
#[cfg(test)]
mod main_test;
