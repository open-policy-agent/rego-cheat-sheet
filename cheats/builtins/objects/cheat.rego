package cheat

obj := {"userid": "18472", "roles": [{"name": "admin"}]}

# paths can contain array indexes too
val := object.get(obj, ["roles", 0, "name"], "missing")

defaulted_val := object.get(
	obj,
	["roles", 0, "permissions"], # path
	"unknown", # default if path is missing
)

keys := object.keys(obj)
