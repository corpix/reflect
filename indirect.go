package reflect

import (
	"reflect"
)

type Value = reflect.Value
type Type = reflect.Type

var (
	ValueOf = reflect.ValueOf
	TypeOf  = reflect.TypeOf
)

func IndirectValue(reflectValue Value) Value {
	if reflectValue.Kind() == reflect.Ptr {
		return reflectValue.Elem()
	}
	return reflectValue
}

func IndirectType(reflectType Type) Type {
	if reflectType.Kind() == reflect.Ptr || reflectType.Kind() == reflect.Slice {
		return reflectType.Elem()
	}
	return reflectType
}
