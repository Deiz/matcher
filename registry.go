package matcher

import (
	"errors"
	"fmt"
	"reflect"
)

var registry map[reflect.Type]Comparator

func Register(rtype reflect.Type, wrapper Comparator) {
	if registry == nil {
		registry = make(map[reflect.Type]Comparator)
	}

	registry[rtype] = wrapper
}

func GetComparator(rtype reflect.Type) (Comparator, error) {
	if registry == nil {
		return nil, errors.New("Comparator registry isn't initialized")
	}

	if cmp, ok := registry[rtype]; !ok {
		return nil, errors.New(fmt.Sprintf("A comparator for the type %+v is not registered", rtype))
	} else {
		return cmp, nil
	}
}
