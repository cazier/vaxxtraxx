# Python

## Dependencies
  1. [Arrow](https://arrow.readthedocs.io/en/latest/index.html) - Python *does* include a library to manage dates and tiems. However, it can be a bit of a headache to work with when manipulating date intervals.

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

All good!