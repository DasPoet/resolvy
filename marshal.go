package resolvy

import (
	"reflect"
)

// MarshalConfig configures Marshal.
type MarshalConfig struct {
	// Tag is the struct tag that configures the name of a
	// field and whether it should be omitted if it is empty.
	Tag string

	// Marshalers contains marshalers for specific fields.
	// If a field has a struct tag defining its name, the
	// key for the field's marshaler must match that name.
	Marshalers map[string]FieldMarshaler
}

// normalise returns a copy of config containing default
// values instead of zero values for certain fields.
func (config MarshalConfig) normalise() MarshalConfig {
	if config.Tag == "" {
		config.Tag = TagName
	}
	if config.Marshalers == nil {
		config.Marshalers = make(map[string]FieldMarshaler)
	}
	return config
}

type FieldMarshaler func() (interface{}, error)

// Marshal produces a map from v.
//
// Marshal iterates over v's exported fields, adding the
// marshalled representation for each field to the returned
// map.
//
// If the given config contains a resolver for a field,
// Marshal uses the resolver to produce the field's
// marshalled representation. Otherwise, Marshal uses
// the field's value as its marshalled representation.
//
// The encoding of struct fields can be customised by the
// format string stored under the "resolvy" key in the struct
// field's tag. This key can be changed, however, by supplying
// a different key via the config parameter to Marshal.
//
// The format string gives the name of the field, possibly
// followed by the "omitempty" option.The name may be empty
// in order to specify the "omitempty"option without overriding
// the default field name.
//
// The "omitempty" option specifies that the field should
// be omitted from the encoding if the field has an empty
// value.
//
// As a special case, if the field tag is "-",the field is
// always omitted. Note that a field with name "-" can still
// be generated using the tag "-,".
//
// If v is neither a struct nor a pointer to a
// struct, Marshal returns an InvalidArgError.
//
func Marshal(v interface{}, config MarshalConfig) (map[string]interface{}, error) {
	return marshal(v, config.normalise())
}

func marshal(v interface{}, config MarshalConfig) (map[string]interface{}, error) {
	fields := make(map[string]interface{})
	return fields, forEachField(v, marshalField(fields, config), isExported, shouldOmit(config.Tag))
}

func marshalField(fields map[string]interface{}, config MarshalConfig) fieldFunc {
	return func(field reflect.StructField, value reflect.Value) error {
		data, err := marshalOrFallback(field, value, config)
		if err == nil {
			fields[fieldName(field, config.Tag)] = data
		}
		return err
	}
}

func marshalOrFallback(field reflect.StructField, value reflect.Value, config MarshalConfig) (interface{}, error) {
	marshaler, ok := config.Marshalers[fieldName(field, config.Tag)]
	if !ok {
		return value.Interface(), nil
	}
	return marshaler()
}
