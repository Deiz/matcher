package matcher

import "errors"

type FloatComparator struct{}

func (s *FloatComparator) Valid(data interface{}) error {
	if _, ok := data.(float64); !ok {
		return errors.New("Invalid argument")
	}

	return nil
}

func (s *FloatComparator) GreaterThan(a, b interface{}) bool {
	return a.(float64) > b.(float64)
}

func (s *FloatComparator) LessThan(a, b interface{}) bool {
	return a.(float64) < b.(float64)
}

func (s *FloatComparator) EqualTo(a, b interface{}) bool {
	return a.(float64) == b.(float64)
}

func (s *FloatComparator) NotEqualTo(a, b interface{}) bool {
	return a.(float64) != b.(float64)
}
