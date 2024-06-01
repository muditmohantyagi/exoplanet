package helper

import (
	"reflect"
	"strings"
)

/*
This function will trim white space from
structure
*/
func Trimmer(structure interface{}) {
	msValuePtr := reflect.ValueOf(structure).Elem()

	for i := 0; i < msValuePtr.NumField(); i++ {
		field := msValuePtr.Field(i)

		if field.Kind() == reflect.Struct {
			// Recursively call trimmer for nested structs.
			Trimmer(field.Addr().Interface())
		}

		// Ignore fields that don't have the same type as a string
		if field.Type() != reflect.TypeOf("") {
			continue
		}

		str := field.Interface().(string)
		str = strings.TrimSpace(str)
		field.SetString(str)
	}
}
