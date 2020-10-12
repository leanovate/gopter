package prop

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

func deepCloneArgs(args []reflect.Value) []reflect.Value {
	dst := make([]reflect.Value, 0, len(args))

	for _, arg := range args {
		dst = append(dst, deepClone(arg))
	}

	return dst
}

func deepClone(v reflect.Value) reflect.Value {
	switch v.Kind() {
	case reflect.Array:
		return deepCloneArray(v)
	case reflect.Slice:
		return deepCloneSlice(v)
	case reflect.Map:
		return deepCloneMap(v)
	default:
		return v
	}
}

func deepCloneArray(v reflect.Value) reflect.Value {
	dst := reflect.New(v.Type())
	num := v.Len()

	if isScalar(v.Type().Elem().Kind()) {
		p := unsafe.Pointer(dst.Pointer())
		dst = dst.Elem()

		shadowCopy(v, p)
	} else {
		dst = dst.Elem()

		for i := 0; i < num; i++ {
			dst.Index(i).Set(deepClone(v.Index(i)))
		}
	}

	return dst
}

func deepCloneSlice(v reflect.Value) reflect.Value {
	if v.IsNil() {
		return reflect.Zero(v.Type())
	}

	num := v.Len()
	cap := v.Cap()
	dst := reflect.MakeSlice(v.Type(), num, cap)

	if isScalar(v.Type().Elem().Kind()) {
		src := unsafe.Pointer(v.Pointer())
		dst := unsafe.Pointer(dst.Pointer())
		sz := int(v.Type().Elem().Size())
		l := num * sz
		cc := cap * sz
		copy((*[math.MaxInt32]byte)(dst)[:l:cc], (*[math.MaxInt32]byte)(src)[:l:cc])
	} else {
		for i := 0; i < num; i++ {
			dst.Index(i).Set(deepClone(v.Index(i)))
		}
	}

	return dst
}

func deepCloneMap(v reflect.Value) reflect.Value {
	if v.IsNil() {
		return reflect.Zero(v.Type())
	}

	dst := reflect.MakeMap(v.Type())

	for iter := v.MapRange(); iter.Next(); {
		dst.SetMapIndex(deepClone(iter.Key()), deepClone(iter.Value()))
	}

	return dst
}

func isScalar(k reflect.Kind) bool {
	switch k {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.String, reflect.Func,
		reflect.UnsafePointer,
		reflect.Invalid:
		return true
	}

	return false
}

func shadowCopy(src reflect.Value, p unsafe.Pointer) {
	switch src.Kind() {
	case reflect.Bool:
		*(*bool)(p) = src.Bool()
	case reflect.Int:
		*(*int)(p) = int(src.Int())
	case reflect.Int8:
		*(*int8)(p) = int8(src.Int())
	case reflect.Int16:
		*(*int16)(p) = int16(src.Int())
	case reflect.Int32:
		*(*int32)(p) = int32(src.Int())
	case reflect.Int64:
		*(*int64)(p) = src.Int()
	case reflect.Uint:
		*(*uint)(p) = uint(src.Uint())
	case reflect.Uint8:
		*(*uint8)(p) = uint8(src.Uint())
	case reflect.Uint16:
		*(*uint16)(p) = uint16(src.Uint())
	case reflect.Uint32:
		*(*uint32)(p) = uint32(src.Uint())
	case reflect.Uint64:
		*(*uint64)(p) = src.Uint()
	case reflect.Uintptr:
		*(*uintptr)(p) = uintptr(src.Uint())
	case reflect.Float32:
		*(*float32)(p) = float32(src.Float())
	case reflect.Float64:
		*(*float64)(p) = src.Float()
	case reflect.Complex64:
		*(*complex64)(p) = complex64(src.Complex())
	case reflect.Complex128:
		*(*complex128)(p) = src.Complex()

	case reflect.Array:
		t := src.Type()

		if src.CanAddr() {
			srcPtr := unsafe.Pointer(src.UnsafeAddr())
			sz := t.Size()
			copy((*[math.MaxInt32]byte)(p)[:sz:sz], (*[math.MaxInt32]byte)(srcPtr)[:sz:sz])
			return
		}

		val := reflect.NewAt(t, p).Elem()

		if src.CanInterface() {
			val.Set(src)
			return
		}

		sz := t.Elem().Size()
		num := src.Len()

		for i := 0; i < num; i++ {
			elemPtr := unsafe.Pointer(uintptr(p) + uintptr(i)*sz)
			shadowCopy(src.Index(i), elemPtr)
		}
	case reflect.Map:
		*((*uintptr)(p)) = src.Pointer()
	case reflect.Ptr:
		*((*uintptr)(p)) = src.Pointer()
	case reflect.Slice:
		*(*reflect.SliceHeader)(p) = reflect.SliceHeader{
			Data: src.Pointer(),
			Len:  src.Len(),
			Cap:  src.Cap(),
		}
	case reflect.String:
		s := src.String()
		*(*reflect.StringHeader)(p) = *(*reflect.StringHeader)(unsafe.Pointer(&s))
	case reflect.Struct:
		t := src.Type()

		if src.CanAddr() {
			srcPtr := unsafe.Pointer(src.UnsafeAddr())
			sz := t.Size()
			copy((*[math.MaxInt32]byte)(p)[:sz:sz], (*[math.MaxInt32]byte)(srcPtr)[:sz:sz])
			return
		}

		val := reflect.NewAt(t, p).Elem()

		if src.CanInterface() {
			val.Set(src)
			return
		}

		num := t.NumField()

		for i := 0; i < num; i++ {
			field := t.Field(i)
			fieldPtr := unsafe.Pointer(uintptr(p) + field.Offset)
			shadowCopy(src.Field(i), fieldPtr)
		}

	default:
		panic(fmt.Errorf("impossible type `%v` when cloning private field", src.Type()))
	}
}
