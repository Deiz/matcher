package matcher

import (
	"errors"
	"reflect"
	"time"
)

type Comparator interface {
	GreaterThan(interface{}, interface{}) bool
	LessThan(interface{}, interface{}) bool
	EqualTo(interface{}, interface{}) bool
	NotEqualTo(interface{}, interface{}) bool

	Valid(interface{}) error
}

func RegisterDefaults() {
	Register(reflect.ValueOf("foo").Type(), &StringComparator{})
	Register(reflect.ValueOf(3.14159).Type(), &FloatComparator{})
	Register(reflect.ValueOf(1).Type(), &IntComparator{})
	Register(reflect.ValueOf(time.Now()).Type(), &TimeComparator{})
}

func Compare(comparator Comparator, operator string, lh, rh interface{}) (bool, error) {
	switch op := operator; op {
	case "gt":
		return comparator.GreaterThan(lh, rh), nil
	case "lt":
		return comparator.LessThan(lh, rh), nil
	case "eq":
		return comparator.EqualTo(lh, rh), nil
	case "ne":
		return comparator.NotEqualTo(lh, rh), nil
	default:
		return false, errors.New("Unknown comparison operator " + operator)
	}
}
