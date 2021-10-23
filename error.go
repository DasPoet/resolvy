package resolvy

import (
	"fmt"
	"reflect"
)

// InvalidArgError is the error that occurs when Marshal
// or Unmarshal is called with an invalid argument (i.e.
// neither a struct, nor a pointer to a struct).
type InvalidArgError struct {
	gotType reflect.Type
}

// Error returns the error representation of err.
func (err InvalidArgError) Error() string {
	return formatError("invalid argument: wanted `struct` or `*struct`, but got `%s`", err.gotType.Kind())
}

// formatError formats the given error message.
func formatError(format string, args ...interface{}) string {
	return "resolvy: " + fmt.Sprintf(format, args...)
}
