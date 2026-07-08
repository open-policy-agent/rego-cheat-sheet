package cheat

valid_staff_email if {
	regex.match(`^\S+@\S+\.\S+$`, input.email) # and
	endswith(input.email, "example.com")
}
