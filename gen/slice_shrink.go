package gen

import (
	"fmt"
	"reflect"

	"github.com/leanovate/gopter"
)

type sliceShrink struct {
	original    reflect.Value
	length      int
	offset      int
	chunkLength int
}

func (s *sliceShrink) Next() (interface{}, bool) {
	if s.chunkLength == 0 {
		return nil, false
	}
	value := reflect.AppendSlice(reflect.MakeSlice(s.original.Type(), 0, s.length-s.chunkLength), s.original.Slice(0, s.offset))
	s.offset += s.chunkLength
	if s.offset < s.length {
		value = reflect.AppendSlice(value, s.original.Slice(s.offset, s.length))
	} else {
		s.offset = 0
		s.chunkLength >>= 1
	}

	return value.Interface(), true
}

func SliceShrinker(v interface{}) gopter.Shrink {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("%#v is not a slice", v))
	}
	sliceShrink := &sliceShrink{
		original:    rv,
		offset:      0,
		length:      rv.Len(),
		chunkLength: rv.Len() >> 1,
	}

	return sliceShrink.Next
}
