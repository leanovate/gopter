package arbitrary

import (
	"reflect"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
)

func (a *Arbitrary) genForKind(rt reflect.Type) gopter.Gen {
	switch rt.Kind() {
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
	case reflect.Slice:
		if elementGen := a.Gen(rt.Elem()); elementGen != nil {
			return gen.SliceOf(elementGen)
		}
	case reflect.Ptr:
		if rt.Elem().Kind() == reflect.Struct {
			gens := make(map[string]gopter.Gen)
			for i := 0; i < rt.Elem().NumField(); i++ {
				field := rt.Elem().Field(i)
				if gen := a.Gen(field.Type); gen != nil {
					gens[field.Name] = gen
				}
			}
			return gen.StructPtr(rt, gens)
		}
	}
	return nil
}
