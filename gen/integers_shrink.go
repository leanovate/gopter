package gen

import "github.com/leanovate/gopter"

type int64Shrink struct {
	original int64
	half     int64
}

func (s *int64Shrink) Next() (interface{}, bool) {
	if s.half == 0 {
		return nil, false
	}
	value := s.original - s.half
	s.half /= 2
	return value, true
}

// Int64Shrinker is a shrinker for int64 numbers
func Int64Shrinker(v interface{}) gopter.Shrink {
	negShrink := int64Shrink{
		original: -v.(int64),
		half:     -v.(int64),
	}
	posShrink := int64Shrink{
		original: v.(int64),
		half:     v.(int64) / 2,
	}
	return gopter.Shrink(negShrink.Next).Interleave(gopter.Shrink(posShrink.Next))
}

// Int32Shrinker is a shrinker for int32 numbers
func Int32Shrinker(v interface{}) gopter.Shrink {
	return Int64Shrinker(int64(v.(int32))).Map(int64To32)
}

// Int16Shrinker is a shrinker for int16 numbers
func Int16Shrinker(v interface{}) gopter.Shrink {
	return Int64Shrinker(int64(v.(int16))).Map(int64To16)
}

// Int8Shrinker is a shrinker for int8 numbers
func Int8Shrinker(v interface{}) gopter.Shrink {
	return Int64Shrinker(int64(v.(int8))).Map(int64To8)
}
