import arrow


def get_due_dates(start, rate_span, rate_frequency, nb=3):
    """Given a start date, a repeat frequency rate, and a number of results, return a list with the
    next due dates for each booster.

    For example, to get the booster due dates of a shot every 7 months, you would call:

    ```python
    >>> date = arrow.get('2022-01-05')
    >>> get_due_dates(date, 'month', 7)
    ... [<Arrow [2022-01-05T00:00:00+00:00]>,
    <Arrow [2022-08-05T00:00:00+00:00]>,
    <Arrow [2023-03-05T00:00:00+00:00]>]
    ```

    Args:
        start (arrow.Arrow): starting date
        rate_span (str): an exact string matching one of 'weeks', 'months', 'years'
        rate_frequency (int): interval of `rate_spans` per booster
        nb (int, optional): number of subsequent due dates to return. Defaults to 3.

    Returns:
        list[arrow.Arrow]: each of upcoming due dates
    """
    return [start.shift(**{rate_span: rate_frequency * i}) for i in range(1, nb + 1)]



def is_overdue(due_date, check_date, strict=True):
    """When supplied two dates, return whether the user is overdue. i.e., the current (`check_date`)
    is after the due date.

    ```python
    >>> due = arrow.get('2022-01-05')
    >>> check = arrow.get('2022-02-05')
    >>> is_overdue(due, check)
    ... False
    ```

    Args:
        due_date (arrow.Arrow): date that the next booster of the vaccine is due
        check_date (arrow.Arrow): date against which to check
        strict (bool, optional): if True, the exact same date means overdue. Defaults to True.

    Returns:
        bool: True if you are overdue for a vaccine, false if all is good!
    """
    if strict:
        return due_date <= check_date
    
    return due_date < check_date
