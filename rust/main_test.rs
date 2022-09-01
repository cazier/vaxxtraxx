use super::*;

#[test]
fn test_get() {
    let date: NaiveDate = NaiveDate::from_ymd(2020, 02, 01);
    let date_string: &str = "2020-02-01";

    let result = get(date_string);

    // assert!(matches!(result, Result::Ok { .. }));
    assert_eq!(result.ok().unwrap(), date);
}

#[test]
fn test_get_due_dates_months() {
    let date = get_safe("2020-02-26");

    let resp = get_due_dates(date, "months", 6, 3);
	let expected: [NaiveDate; 3] = [
		get_safe("2020-08-26"),
		get_safe("2021-02-26"),
		get_safe("2021-08-26"),
    ];

    for i in 0..expected.len(){
        assert_eq!(expected[i], resp[i]);
    }
}

#[test]
fn test_get_due_dates_leap_years() {
    let date = get_safe("2020-02-29");
    let resp = get_due_dates(date, "years", 1, 4);
	let expected: [NaiveDate; 4] = [
		get_safe("2021-02-28"),
		get_safe("2022-02-28"),
		get_safe("2023-02-28"),
		get_safe("2024-02-29"),
    ];

    for i in 0..expected.len(){
        assert_eq!(expected[i], resp[i]);
    }
}

#[test]
fn test_get_due_dates_leap_days() {
    let date = get_safe("2020-02-26");
    let resp = get_due_dates(date, "days", 1, 5);
	let expected: [NaiveDate; 5] = [
		get_safe("2020-02-27"),
		get_safe("2020-02-28"),
		get_safe("2020-02-29"),
		get_safe("2020-03-01"),
		get_safe("2020-03-02"),
    ];

    for i in 0..(expected.len()){
        assert_eq!(expected[i], resp[i]);

    }
    let date = get_safe("2021-02-26");
    let resp = get_due_dates(date, "days", 1, 5);
	let expected: [NaiveDate; 5] = [
		get_safe("2021-02-27"),
		get_safe("2021-02-28"),
		get_safe("2021-03-01"),
		get_safe("2021-03-02"),
		get_safe("2021-03-03"),
    ];

    for i in 0..(expected.len()){
        assert_eq!(expected[i], resp[i]);
    }
}

#[test]
fn test_overdue_date_is_overdue() {
	let due = get_safe("2020-02-05");
	let now = get_safe("2025-02-05");

	assert!(is_overdue(due, now, true));
}

#[test]
fn test_overdue_date_is_not_overdue() {
	let due = get_safe("2030-02-05");
	let now = get_safe("2025-02-05");

	assert!(!is_overdue(due, now, true));
}

#[test]
fn test_overdue_date_is_equality_overdue() {
	let date = get_safe("2030-02-05");

	assert!(is_overdue(date, date, true));
	assert!(!is_overdue(date, date, false));
}
