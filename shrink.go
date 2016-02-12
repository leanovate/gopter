package gopter

type Shrink func() (interface{}, bool)

type Shrinker func(value interface{}) Shrink

var NoShrink = Shrink(func() (interface{}, bool) {
	return nil, false
})

var NoShrinker = Shrinker(func(value interface{}) Shrink {
	return NoShrink
})
