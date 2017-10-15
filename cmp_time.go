package matcher

import (
	"errors"
	"time"
)

type TimeComparator struct{}

func (s *TimeComparator) Valid(data interface{}) error {
	if _, ok := data.(time.Time); !ok {
		return errors.New("Invalid argument")
	}

	return nil
}

func (s *TimeComparator) GreaterThan(a, b interface{}) bool {
	return a.(time.Time).After(b.(time.Time))
}

func (s *TimeComparator) LessThan(a, b interface{}) bool {
	return a.(time.Time).Before(b.(time.Time))
}

func (s *TimeComparator) EqualTo(a, b interface{}) bool {
	return a.(time.Time).Equal(b.(time.Time))
}

func (s *TimeComparator) NotEqualTo(a, b interface{}) bool {
	return !s.EqualTo(a, b)
}
