package cheat

import rego.v1

example_string := "Build Policy as Code with OPA!"

check_contains if contains(example_string, "OPA")
check_startswith if startswith(example_string, "Build")
check_endswith if endswith(example_string, "!")
check_replace := replace(example_string, "OPA", "OPA!")
check_sprintf := sprintf("OPA is %s!", ["awesome"])
