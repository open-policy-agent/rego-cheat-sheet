package cheat

import rego.v1

valid_staff_email if {
	regex.match(`^\S+@\S+\.\S+$`, input.email) # and
	endswith(input.email, "example.com")
}
