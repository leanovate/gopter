package gen

import (
	"time"

	"github.com/leanovate/gopter"
)

// TimeShrinker is a shrinker for time.Time structs
func TimeShrinker(v interface{}) gopter.Shrink {
	t := v.(time.Time)
	sec := t.Unix()
	nsec := int64(t.Nanosecond())
	secShrink := int64Shrink{
		original: sec,
		half:     sec,
	}
	nsecShrink := int64Shrink{
		original: nsec,
		half:     nsec,
	}
	return gopter.Shrink(secShrink.Next).Map(func(v interface{}) interface{} {
		return time.Unix(v.(int64), nsec)
	}).Interleave(gopter.Shrink(nsecShrink.Next).Map(func(v interface{}) interface{} {
		return time.Unix(sec, v.(int64))
	}))
}
