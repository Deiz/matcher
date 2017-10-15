package matcher

import "errors"

type IntComparator struct{}

func (s *IntComparator) Valid(data interface{}) error {
	if _, ok := data.(int); !ok {
		return errors.New("Invalid argument")
	}

	return nil
}

func (s *IntComparator) GreaterThan(a, b interface{}) bool {
	return a.(int) > b.(int)
}

func (s *IntComparator) LessThan(a, b interface{}) bool {
	return a.(int) < b.(int)
}

func (s *IntComparator) EqualTo(a, b interface{}) bool {
	return a.(int) == b.(int)
}

func (s *IntComparator) NotEqualTo(a, b interface{}) bool {
	return a.(int) != b.(int)
}
