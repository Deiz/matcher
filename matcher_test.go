package matcher

import (
	"fmt"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	RegisterDefaults()
	code := m.Run()
	os.Exit(code)
}

func TestErrorHandling(t *testing.T) {
	Convey("Given a non-struct", t, func() {
		v := "invalid"

		Convey("When a match is attempted, an error is returned", func() {
			clauses := []*Clause{&Clause{Field: "foo", Operator: "eq", Value: "bar"}}
			_, err := Matches(v, clauses)

			So(err, ShouldBeError)
		})
	})

	Convey("Given a struct", t, func() {
		v := struct {
			Name     string `matcher:"name"`
			Age      int    `matcher:"age"`
			Embedded struct {
				Foo int
				Bar float64
			} `matcher:"embedded"`
		}{"Alice", 32, struct {
			Foo int
			Bar float64
		}{32, 3.151459}}

		Convey("When a non-existent field is matched against, an error is returned", func() {
			clauses := []*Clause{&Clause{Field: "invalid", Operator: "eq", Value: "foo"}}
			_, err := Matches(v, clauses)

			So(err, ShouldBeError)
		})

		Convey("When attempting to match with an invalid operator, an error is returned", func() {
			clauses := []*Clause{&Clause{Field: "name", Operator: "foo", Value: "foo"}}
			_, err := Matches(v, clauses)

			So(err, ShouldBeError)
		})

		Convey("When a clause's value type doesn't match the struct, an error is returned", func() {
			clauses := []*Clause{&Clause{Field: "name", Operator: "eq", Value: 3.14159}}
			_, err := Matches(v, clauses)

			So(err, ShouldBeError)
		})

		Convey("When a comparison is attempted against an unknown type, an error is returned", func() {
			clauses := []*Clause{&Clause{Field: "embedded", Operator: "eq", Value: 3.14159}}
			_, err := Matches(v, clauses)

			So(err, ShouldBeError)
		})
	})
}

func TestMatching(t *testing.T) {
	Convey("Given a struct", t, func() {
		v := struct {
			Name    string `matcher:"name"`
			Age     int    `matcher:"age"`
			Unused  bool
			Skipped bool `matcher:"-"`
		}{"Alice", 32, false, false}

		Convey("When no clauses are provided, the struct should match", func() {
			clauses := []*Clause{}
			matched, err := Matches(v, clauses)

			So(err, ShouldBeNil)
			So(matched, ShouldBeTrue)
		})

		Convey("When all clauses match, the struct should match", func() {
			clauses := []*Clause{&Clause{Field: "name", Operator: "eq", Value: "Alice"},
				&Clause{Field: "age", Operator: "eq", Value: 32}}
			matched, err := Matches(v, clauses)

			So(err, ShouldBeNil)
			So(matched, ShouldBeTrue)
		})

		Convey("If one clause doesn't match, the struct shouldn't match", func() {
			clauses := []*Clause{&Clause{Field: "name", Operator: "eq", Value: "Alice"},
				&Clause{Field: "age", Operator: "eq", Value: 40}}
			matched, err := Matches(v, clauses)

			So(err, ShouldBeNil)
			So(matched, ShouldBeFalse)
		})
	})
}

func TestTypeMatching(t *testing.T) {
	type Person struct {
		Name string `matcher:"name"`
		Age  int    `matcher:"age"`
	}

	type TypeWrapper struct {
		Name        string
		Field       string
		StructField string
		Base        interface{}
		Data        map[string]interface{}
	}

	p := Person{
		Name: "foo",
		Age:  1,
	}

	operators := []string{"gt", "lt", "eq", "ne"}

	types := []*TypeWrapper{
		{
			Name:        "string",
			Field:       "name",
			StructField: "Name",
			Base:        "foo",
			Data: map[string]interface{}{
				"gt": "e",
				"lt": "g",
				"eq": "foo",
				"ne": "bar",
			},
		},
		{
			Name:        "int",
			Field:       "age",
			StructField: "Age",
			Base:        1,
			Data: map[string]interface{}{
				"gt": 0,
				"lt": 2,
				"eq": 1,
				"ne": 3,
			},
		},
	}

	Convey("Given a struct", t, func() {
		for _, typeTest := range types {
			for _, op := range operators {
				msg := fmt.Sprintf("When comparing %ss with the %s method, %v and %v match",
					typeTest.Name, op, typeTest.Base, typeTest.Data[op])
				Convey(msg, func() {
					clauses := []*Clause{&Clause{Field: typeTest.Field, Operator: op, Value: typeTest.Data[op]}}
					matched, err := Matches(p, clauses)

					So(err, ShouldBeNil)
					So(matched, ShouldBeTrue)
				})
			}
		}
	})
}
