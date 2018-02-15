package reflect

import (
	"reflect"
)

var (
	TypeOf = reflect.TypeOf
)

type Type = reflect.Type

var (
	Types = []Type{
		TypeBool,
		TypeInt,
		TypeInt8,
		TypeInt16,
		TypeInt32,
		TypeInt64,
		TypeUint,
		TypeUint8,
		TypeUint16,
		TypeUint32,
		TypeUint64,
		TypeFloat32,
		TypeFloat64,
		TypeComplex64,
		TypeComplex128,
		TypeUintptr,
		TypeString,
	}

	TypeInvalid    = Type(nil)
	TypeBool       = reflect.TypeOf(false)
	TypeInt        = reflect.TypeOf(int(0))
	TypeInt8       = reflect.TypeOf(int8(0))
	TypeInt16      = reflect.TypeOf(int16(0))
	TypeInt32      = reflect.TypeOf(int32(0))
	TypeInt64      = reflect.TypeOf(int64(0))
	TypeUint       = reflect.TypeOf(uint(0))
	TypeUint8      = reflect.TypeOf(uint8(0))
	TypeUint16     = reflect.TypeOf(uint16(0))
	TypeUint32     = reflect.TypeOf(uint32(0))
	TypeUint64     = reflect.TypeOf(uint64(0))
	TypeFloat32    = reflect.TypeOf(float32(0))
	TypeFloat64    = reflect.TypeOf(float64(0))
	TypeComplex64  = reflect.TypeOf(complex64(0))
	TypeComplex128 = reflect.TypeOf(complex128(0))
	TypeUintptr    = reflect.TypeOf(uintptr(0))
	TypeString     = reflect.TypeOf(string(""))
)

func ParseType(p string) Type {
	var (
		buf string
	)

	for k, c := range p {
		buf += string(c)

		switch buf {
		case "map":
			first := k + 2
			last := first
			for kk, vv := range p[first:] {
				switch vv {
				case ']':
					last = first + kk
				default:
					continue
				}
			}

			return reflect.MapOf(
				ParseType(string(p[first])+p[first+1:last]),
				ParseType(p[last+1:]),
			)
		case "[]":
			return reflect.SliceOf(ParseType(p[2:]))
		case "*":
			return reflect.PtrTo(ParseType(p[1:]))
		default:
			continue
		}
	}

	switch p {
	case "bool":
		return TypeBool
	case "int":
		return TypeInt
	case "int8":
		return TypeInt8
	case "int16":
		return TypeInt16
	case "int32":
		return TypeInt32
	case "int64":
		return TypeInt64
	case "uint":
		return TypeUint
	case "uint8":
		return TypeUint8
	case "uint16":
		return TypeUint16
	case "uint32":
		return TypeUint32
	case "uint64":
		return TypeUint64
	case "float32":
		return TypeFloat32
	case "float64":
		return TypeFloat64
	case "complex64":
		return TypeComplex64
	case "complex128":
		return TypeComplex128
	case "uintptr":
		return TypeUintptr
	case "string":
		return TypeString
	default:
		return TypeInvalid
	}
}
