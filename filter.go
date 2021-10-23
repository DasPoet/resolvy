package resolvy

import (
	"reflect"
	"strings"
)

// A filterFunc is a predicate for filtering struct fields.
type filterFunc func(field reflect.StructField, value reflect.Value) (keep bool)

// all reports whether all the given filters return true for the given field and value.
func all(field reflect.StructField, value reflect.Value, filters ...filterFunc) bool {
	for _, f := range filters {
		if !f(field, value) {
			return false
		}
	}
	return true
}

// isExported is a filterFunc for filtering exported struct fields.
func isExported(field reflect.StructField, _ reflect.Value) bool {
	return field.IsExported()
}

// shouldOmit returns a filterFunc that omits fields with a
// struct tag with the given name that contains ",omitempty"
// or that is equal to "-".
func shouldOmit(tagName string) filterFunc {
	return func(field reflect.StructField, value reflect.Value) (keep bool) {
		tag := field.Tag.Get(tagName)
		if tag == "-" {
			return false
		}
		parts := strings.Split(tag, ",")
		omit := len(parts) >= 2 && parts[1] == "omitempty"
		return !(omit && value.IsValid() && value.IsZero())
	}
}
