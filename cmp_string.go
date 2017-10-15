package matcher

import "errors"

type StringComparator struct{}

func (s *StringComparator) Valid(data interface{}) error {
	if _, ok := data.(string); !ok {
		return errors.New("Invalid argument")
	}

	return nil
}

func (s *StringComparator) GreaterThan(a, b interface{}) bool {
	return a.(string) > b.(string)
}

func (s *StringComparator) LessThan(a, b interface{}) bool {
	return a.(string) < b.(string)
}

func (s *StringComparator) EqualTo(a, b interface{}) bool {
	return a.(string) == b.(string)
}

func (s *StringComparator) NotEqualTo(a, b interface{}) bool {
	return a.(string) != b.(string)
}
