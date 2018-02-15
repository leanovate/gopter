package arbitrary

import (
	"reflect"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
)

func (a *Arbitraries) genForKind(rt reflect.Type) gopter.Gen {
	switch rt.Kind() {
	case reflect.Bool:
		return gen.Bool().MapResult(func(result *gopter.GenResult) *gopter.GenResult {
			value := reflect.New(rt).Elem()
			if v, ok := result.Result.(bool); ok {
				value.SetBool(v)
			}
			return &gopter.GenResult{
				Labels:     result.Labels,
				ResultType: rt,
				Result:     value.Interface(),
			}
		})
	case reflect.Int:
		return gen.Int().MapResult(func(result *gopter.GenResult) *gopter.GenResult {
			value := reflect.New(rt).Elem()
			if v, ok := result.Result.(int); ok {
				value.SetInt(int64(v))
			}
			return &gopter.GenResult{
				Labels:     result.Labels,
				ResultType: rt,
				Result:     value.Interface(),
			}
		})
	case reflect.Uint:
		return gen.UInt().MapResult(func(result *gopter.GenResult) *gopter.GenResult {
			value := reflect.New(rt).Elem()
			if v, ok := result.Result.(uint); ok {
				value.SetUint(uint64(v))
			}
			return &gopter.GenResult{
				Labels:     result.Labels,
				ResultType: rt,
				Result:     value.Interface(),
			}
		})
	case reflect.Int8:
		return gen.Int8().MapResult(func(result *gopter.GenResult) *gopter.GenResult {
			value := reflect.New(rt).Elem()
			if v, ok := result.Result.(int8); ok {
				value.SetInt(int64(v))
			}
			return &gopter.GenResult{
				Labels:     result.Labels,
				ResultType: rt,
				Result:     value.Interface(),
			}
		})
	case reflect.Uint8:
		return gen.UInt8().MapResult(func(result *gopter.GenResult) *gopter.GenResult {
			value := reflect.New(rt).Elem()
			if v, ok := result.Result.(uint8); ok {
				value.SetUint(uint64(v))
			}
			return &gopter.GenResult{
				Labels:     result.Labels,
				ResultType: rt,
				Result:     value.Interface(),
			}
		})
	case reflect.Int16:
		return gen.Int16().MapResult(func(result *gopter.GenResult) *gopter.GenResult {
			value := reflect.New(rt).Elem()
			if v, ok := result.Result.(int16); ok {
				value.SetInt(int64(v))
			}
			return &gopter.GenResult{
				Labels:     result.Labels,
				ResultType: rt,
				Result:     value.Interface(),
			}
		})
	case reflect.Uint16:
		return gen.UInt16().MapResult(func(result *gopter.GenResult) *gopter.GenResult {
			value := reflect.New(rt).Elem()
			if v, ok := result.Result.(uint16); ok {
				value.SetUint(uint64(v))
			}
			return &gopter.GenResult{
				Labels:     result.Labels,
				ResultType: rt,
				Result:     value.Interface(),
			}
		})
	case reflect.Int32:
		return gen.Int32().MapResult(func(result *gopter.GenResult) *gopter.GenResult {
			value := reflect.New(rt).Elem()
			if v, ok := result.Result.(int32); ok {
				value.SetInt(int64(v))
			}
			return &gopter.GenResult{
				Labels:     result.Labels,
				ResultType: rt,
				Result:     value.Interface(),
			}
		})
	case reflect.Uint32:
		return gen.UInt32().MapResult(func(result *gopter.GenResult) *gopter.GenResult {
			value := reflect.New(rt).Elem()
			if v, ok := result.Result.(uint32); ok {
				value.SetUint(uint64(v))
			}
			return &gopter.GenResult{
				Labels:     result.Labels,
				ResultType: rt,
				Result:     value.Interface(),
			}
		})
	case reflect.Int64:
		return gen.Int64().MapResult(func(result *gopter.GenResult) *gopter.GenResult {
			value := reflect.New(rt).Elem()
			if v, ok := result.Result.(int64); ok {
				value.SetInt(v)
			}
			result.ResultType = rt
			result.Result = value.Interface()
			return result
		})
	case reflect.Uint64:
		return gen.UInt64().MapResult(func(result *gopter.GenResult) *gopter.GenResult {
			value := reflect.New(rt).Elem()
			if v, ok := result.Result.(uint64); ok {
				value.SetUint(v)
			}
			return &gopter.GenResult{
				Labels:     result.Labels,
				ResultType: rt,
				Result:     value.Interface(),
			}
		})
	case reflect.Float32:
		return gen.Float32().MapResult(func(result *gopter.GenResult) *gopter.GenResult {
			value := reflect.New(rt).Elem()
			if v, ok := result.Result.(float32); ok {
				value.SetFloat(float64(v))
			}
			return &gopter.GenResult{
				Labels:     result.Labels,
				ResultType: rt,
				Result:     value.Interface(),
			}
		})
	case reflect.Float64:
		return gen.Float64().MapResult(func(result *gopter.GenResult) *gopter.GenResult {
			value := reflect.New(rt).Elem()
			if v, ok := result.Result.(float64); ok {
				value.SetFloat(v)
			}
			return &gopter.GenResult{
				Labels:     result.Labels,
				ResultType: rt,
				Result:     value.Interface(),
			}
		})
	case reflect.Complex64:
		return gen.Complex64().MapResult(func(result *gopter.GenResult) *gopter.GenResult {
			value := reflect.New(rt).Elem()
			if v, ok := result.Result.(complex64); ok {
				value.SetComplex(complex128(v))
			}
			return &gopter.GenResult{
				Labels:     result.Labels,
				ResultType: rt,
				Result:     value.Interface(),
			}
		})
	case reflect.Complex128:
		return gen.Complex128().MapResult(func(result *gopter.GenResult) *gopter.GenResult {
			value := reflect.New(rt).Elem()
			if v, ok := result.Result.(complex128); ok {
				value.SetComplex(v)
			}
			return &gopter.GenResult{
				Labels:     result.Labels,
				ResultType: rt,
				Result:     value.Interface(),
			}
		})
	case reflect.String:
		return gen.AnyString().MapResult(func(result *gopter.GenResult) *gopter.GenResult {
			value := reflect.New(rt).Elem()
			if v, ok := result.Result.(string); ok {
				value.SetString(v)
			}
			return &gopter.GenResult{
				Labels:     result.Labels,
				ResultType: rt,
				Result:     value.Interface(),
			}
		})
	case reflect.Slice:
		if elementGen := a.GenForType(rt.Elem()); elementGen != nil {
			return gen.SliceOf(elementGen)
		}
	case reflect.Ptr:
		if rt.Elem().Kind() == reflect.Struct {
			gens := make(map[string]gopter.Gen)
			for i := 0; i < rt.Elem().NumField(); i++ {
				field := rt.Elem().Field(i)
				if gen := a.GenForType(field.Type); gen != nil {
					gens[field.Name] = gen
				}
			}
			return gen.StructPtr(rt, gens)
		}
		return gen.PtrOf(a.GenForType(rt.Elem()))
	}
	return nil
}
