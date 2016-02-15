package gen

import "github.com/untoldwind/gopter"

type int64Shrink struct {
	original int64
	half     int64
}

func (s *int64Shrink) Next() (interface{}, bool) {
	if s.half == 0 {
		return nil, false
	}
	value := s.original - s.half
	s.half >>= 1
	return value, true
}

func Int64Shrinker(v interface{}) gopter.Shrink {
	int64Shrink := int64Shrink{
		original: v.(int64),
		half:     v.(int64),
	}
	shrink := gopter.Shrink(int64Shrink.Next)
	return shrink.Interleave(shrink.Map(func(v interface{}) interface{} {
		return -v.(int64)
	}))
}

func Int32Shrinker(v interface{}) gopter.Shrink {
	return Int64Shrinker(int64(v.(int32))).Map(int64To32)
}

func Int16Shrinker(v interface{}) gopter.Shrink {
	return Int64Shrinker(int64(v.(int16))).Map(int64To16)
}

func Int8Shrinker(v interface{}) gopter.Shrink {
	return Int64Shrinker(int64(v.(int8))).Map(int64To8)
}
