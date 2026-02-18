package internal

import "reflect"

// Check isNil to Struct, Map, Slice, Chan, Pointer.
func IsInterfaceNil(v any) bool {
	if v == nil {
		return true
	}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Chan, reflect.Pointer:
		return reflect.ValueOf(v).IsNil()
	}

	return false
}
