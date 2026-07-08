package cheat

unioned := object.union({"foo": true}, {"bar": false})

subset := object.subset(
	{"foo": true, "bar": false},
	{"foo": true}, # subset object
)

removed := object.remove(
	{"foo": true, "bar": false},
	{"bar"}, # remove keys
)
