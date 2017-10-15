package main

import (
	"encoding/json"
	"fmt"

	"github.com/Deiz/matcher"
)

type Person struct {
	Name string `matcher:"name"`
	Age  int    `matcher:"age"`
}

func main() {
	matcher.RegisterDefaults()

	clausejson := `[{ "field": "name", "op": "eq", "value": "Alice" }]`
	clauses := []*matcher.Clause{}
	json.Unmarshal([]byte(clausejson), &clauses)

	p := &Person{"Alice", 32}

	matched, err := matcher.Matches(*p, clauses)
	fmt.Println(matched, err)
}
