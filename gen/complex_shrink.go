package gen

import "github.com/leanovate/gopter"

func Complex128Shrinker(v interface{}) gopter.Shrink {
	c := v.(complex128)
	realShrink := Float64Shrinker(real(c)).Map(func(r interface{}) interface{} {
		return complex(r.(float64), imag(c))
	})
	imagShrink := Float64Shrinker(imag(c)).Map(func(i interface{}) interface{} {
		return complex(real(c), i.(float64))
	})
	return realShrink.Interleave(imagShrink)
}

func Complex64Shrinker(v interface{}) gopter.Shrink {
	c := v.(complex64)
	realShrink := Float64Shrinker(float64(real(c))).Map(func(r interface{}) interface{} {
		return complex(float32(r.(float64)), imag(c))
	})
	imagShrink := Float64Shrinker(float64(imag(c))).Map(func(i interface{}) interface{} {
		return complex(real(c), float32(i.(float64)))
	})
	return realShrink.Interleave(imagShrink)
}
