package gen

import "github.com/leanovate/gopter"

// Complex128Box generate complex128 numbers within a rectangle/box in the complex plane
func Complex128Box(min, max complex128) gopter.Gen {
	return gopter.CombineGens(
		Float64Range(real(min), real(max)),
		Float64Range(imag(min), imag(max)),
	).Map(func(v interface{}) interface{} {
		values := v.([]interface{})
		return complex(values[0].(float64), values[1].(float64))
	}).SuchThat(func(v interface{}) bool {
		c := v.(complex128)
		return real(c) >= real(min) && real(c) <= real(max) &&
			imag(c) >= imag(min) && imag(c) <= imag(max)
	}).WithShrinker(Complex128Shrinker)
}

// Complex128 generate arbitrary complex128 numbers
func Complex128() gopter.Gen {
	return gopter.CombineGens(
		Float64(),
		Float64(),
	).Map(func(v interface{}) interface{} {
		values := v.([]interface{})
		return complex(values[0].(float64), values[1].(float64))
	}).WithShrinker(Complex128Shrinker)
}

// Complex64Box generate complex64 numbers within a rectangle/box in the complex plane
func Complex64Box(min, max complex64) gopter.Gen {
	return gopter.CombineGens(
		Float32Range(real(min), real(max)),
		Float32Range(imag(min), imag(max)),
	).Map(func(v interface{}) interface{} {
		values := v.([]interface{})
		return complex(values[0].(float32), values[1].(float32))
	}).SuchThat(func(v interface{}) bool {
		c := v.(complex64)
		return real(c) >= real(min) && real(c) <= real(max) &&
			imag(c) >= imag(min) && imag(c) <= imag(max)
	}).WithShrinker(Complex64Shrinker)
}

// Complex64 generate arbitrary complex64 numbers
func Complex64() gopter.Gen {
	return gopter.CombineGens(
		Float32(),
		Float32(),
	).Map(func(v interface{}) interface{} {
		values := v.([]interface{})
		return complex(values[0].(float32), values[1].(float32))
	}).WithShrinker(Complex128Shrinker)
}
