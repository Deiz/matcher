The matcher package is a quick experiment in reflection-driven dynamic matching against arbitrary structs.

The intent is that given a query expression (JSON in the example utility, but equally composable by hand), a struct can be evaluted against one or more conditionals that check whether its properties match certain criteria.

An example JSON query might look like:
```json
[
	{
		"field": "name",
		"operator": "eq",
		"value": "Alice"
	},
	{
		"field": "age",
		"operator": "gt",
		"value": 30
	}
]
```

A struct we want to match against might be defined as such:


```go
type Person struct {
	Name string `matcher:"name"`
	Age  int    `matcher:"age"`
}
```

Usage then, looks something like:

```go
matcher.RegisterDefaults()

clauses := []*matcher.Clause{&matcher.Clause{
	Field:    "name",
	Operator: "eq",
	Value:    "Alice",
}}

p := Person{
	Name: "Alice",
	Age:  32,
}

// true, nil
matched, err := matcher.Matches(p, clauses)
```
