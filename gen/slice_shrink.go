package gen

import (
	"fmt"
	"reflect"

	"github.com/leanovate/gopter"
)

type sliceShrink struct {
	original    reflect.Value
	value       reflect.Value
	length      int
	offset      int
	chunkLength int
}

func (s *sliceShrink) Next() (interface{}, bool) {
	if s.length == 0 {
		return nil, false
	}
	reflect.Copy(s.value, s.original.Slice(0, s.offset))
	s.offset += s.chunkLength
	if s.offset < s.length {
		s.value = reflect.AppendSlice(s.value, s.original.Slice(s.offset, s.length))
	} else {
		s.offset = 0
		s.chunkLength >>= 1
	}

	return s.value.Interface(), true
}

func SliaceShrinker(v interface{}) gopter.Shrink {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("%#v is not a slice", v))
	}
	sliceShrink := &sliceShrink{
		original:    rv,
		value:       reflect.MakeSlice(rv.Type().Elem(), 0, rv.Len()),
		offset:      0,
		length:      rv.Len(),
		chunkLength: rv.Len(),
	}

	return sliceShrink.Next
}
