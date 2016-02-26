package arbitrary

import (
	"reflect"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
)

func (a *Arbitrary) genForKind(kind reflect.Kind) gopter.Gen {
	switch kind {
	case reflect.Bool:
		return gen.Bool()
	case reflect.Int8:
		return gen.Int8()
	case reflect.Uint8:
		return gen.UInt8()
	case reflect.Int16:
		return gen.Int16()
	case reflect.Uint16:
		return gen.UInt16()
	case reflect.Int32:
		return gen.Int32()
	case reflect.Uint32:
		return gen.UInt32()
	case reflect.Int64:
		return gen.Int64()
	case reflect.Uint64:
		return gen.UInt64()
	case reflect.Float32:
		return gen.Float32()
	case reflect.Float64:
		return gen.Float64()
	case reflect.String:
		return gen.AnyString()
	}
	return nil
}
