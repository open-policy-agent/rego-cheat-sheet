package cheat

import rego.v1

# using multiple rules
valid_email if endswith(input.email, "@example.com")
valid_email if endswith(input.email, "@example.org")
valid_email if endswith(input.email, "@example.net")

# using functions
allowed_firstname(name) if {
	startswith(name, "a")
	count(name) > 2
}

allowed_firstname("joe") # if name == 'joe'

valid_name if allowed_firstname(input.name)

valid_request if {
	input.method in {"GET", "POST"} # using `in`
}
