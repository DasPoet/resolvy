package resolvy

import (
	"reflect"
	"strings"
)

// A fieldFunc is used for iterating over struct fields.
type fieldFunc func(field reflect.StructField, value reflect.Value) error

// forEachField calls f on every field of v for which predicate
// returns true. If v is neither a struct type nor a pointer to
// a struct type, forEachField returns an InvalidArgError.
// forEachField returns the first error returned by
// a call to f, or nil if f never returned an error.
func forEachField(v interface{}, f fieldFunc, predicates ...filterFunc) error {
	typ, val, err := toStructTypeAndValue(v)
	if err != nil {
		return err
	}

	for i := 0; i < typ.NumField(); i++ {
		field, fieldVal := typ.Field(i), val.Field(i)
		if all(field, fieldVal, predicates...) {
			if err = f(field, fieldVal); err != nil {
				return err
			}
		}
	}
	return nil
}

// toStructTypeAndValue returns the reflect.Type and reflect.Value
// of v. If v is neither a struct type nor a pointer to a struct type,
// toStructTypeAndValue returns an InvalidArgError.
func toStructTypeAndValue(v interface{}) (reflect.Type, reflect.Value, error) {
	switch {
	case isStruct(v):
		return reflect.TypeOf(v), reflect.ValueOf(v), nil
	case isStructPointer(v):
		return reflect.TypeOf(v).Elem(), reflect.ValueOf(v).Elem(), nil
	}
	return nil, reflect.Value{}, InvalidArgError{gotType: reflect.TypeOf(v)}
}

// isStruct reports whether v is a struct type.
func isStruct(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Struct
}

// isStructPointer reports whether v is a pointer to a struct type.
func isStructPointer(v interface{}) bool {
	typ := reflect.TypeOf(v)
	return typ.Kind() == reflect.Ptr && typ.Elem().Kind() == reflect.Struct
}

// fieldName returns the given field's name, or the name
// specified in the field's tag with the given tag name,
// if such a tag exists.
func fieldName(field reflect.StructField, tagName string) string {
	if tag := field.Tag.Get(tagName); tag != "" {
		if tag == "-," {
			return "-"
		}
		return strings.Split(tag, ",")[0]
	}
	return field.Name
}
