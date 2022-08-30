# Python

## Dependencies
 1. [Arrow](https://arrow.readthedocs.io/en/latest/index.html) - Python *does* include a library to manage dates and tiems. However, it can be a bit of a headache to work with when manipulating date intervals.

     - The link above goes to the Arrow documentation. That should give you some of the details for how to do things like create an arrow *object* for a particular date.
     - Further, you can look into some of the other parts of the documentation/API reference to get things like ***shift***ing an Arrow object by specific date spans.

## Installation
```bash
$ git clone https://github.com/cazier/vaxxtraxx
$ cd vaxxtraxx/python

# It is recommended to use a virtual environment to keep your dependencies clean. 
$ python3 -m venv venv

# After creating the venv, you will need to remember to activate it:
$ source venv/bin/activate

# Once it's activated, you should see the venv name on your prompt. (It may not be identical to the following, though)

# Now you can drop the 3 from the python (because the venv will manage that) and install dependencies
(python) $ python -m pip install -r requirements.txt
```

All good! Once you have the virtual environment made, you can also run your code in that environment with:
```bash
(python) $ python main.py
```

## Tests
Once you've started to fill in the functions with actual code, you should be able to run the tests to verify that your code does what you expect it to do. To run the test, simply run unittest in your command prompt:

```bash
(python) $ python -m unittest tests/test_main.py
```

You can take a look at the tests in the [tests](./tests/) directory to take a look at what they are doing.
