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

## Registering custom types

Custom types must satisfy the `Comparator` interface:
```go
type Comparator interface {
	GreaterThan(interface{}, interface{}) bool
	LessThan(interface{}, interface{}) bool
	EqualTo(interface{}, interface{}) bool
	NotEqualTo(interface{}, interface{}) bool

	Valid(interface{}) error
}
```

A concrete example of this is [cmp_string.go](cmp_string.go)

Once a comparator wrapper has been implemented for a given type, it needs to be associated with the original type via the registry, e.g.:
```go
// Pass a reflect.Type corresponding to string, plus your comparator interface implementation.
matcher.Register(reflect.ValueOf("foo").Type(), &StringComparator{})
```
