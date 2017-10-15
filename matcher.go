package matcher

import (
	"errors"
	"fmt"
	"reflect"
)

type Clause struct {
	Field    string      `json:"field"`
	Operator string      `json:"op"`
	Value    interface{} `json:"value"`
}

type FieldInfo struct {
	Index int
	Value reflect.Value
	Type  reflect.Type
}

func Matches(data interface{}, clauses []*Clause) (bool, error) {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		return false, errors.New("Matching is only permitted against structs")
	}

	// Return early if matching will be a no-op
	if len(clauses) == 0 {
		return true, nil
	}

	referenced := make(map[string]bool)
	matchable := make(map[string]*FieldInfo)

	for _, clause := range clauses {
		referenced[clause.Field] = true
	}

	st := v.Type()
	n := st.NumField()

	for i := 0; i < n; i++ {
		field := st.Field(i)
		if field.PkgPath != "" && !field.Anonymous {
			continue
		}

		tag := field.Tag.Get("matcher")
		if tag == "-" || tag == "" {
			continue
		}

		matchable[tag] = &FieldInfo{
			Index: i,
			Value: v.Field(i),
			Type:  field.Type,
		}
	}

	for _, clause := range clauses {
		fieldInfo, ok := matchable[clause.Field]
		if !ok {
			return false, fmt.Errorf("\"%s\" is not a matchable field within the struct: %+v", clause.Field, st)
		}

		cmp, err := GetComparator(fieldInfo.Type)
		if err != nil {
			return false, err
		}

		data := fieldInfo.Value.Interface()
		if err := cmp.Valid(data); err != nil {
			return false, err
		} else if err := cmp.Valid(clause.Value); err != nil {
			return false, err
		}

		matched, err := Compare(cmp, clause.Operator, data, clause.Value)
		if err != nil {
			return false, err
		} else if !matched {
			return false, nil
		}
	}

	return true, nil
}
