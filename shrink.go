package gopter

// Shrink is a stream of shrinked down values
// Once the result of a shrink is false, it is considered to be exhausted
type Shrink func() (interface{}, bool)

// Filter creates a shrink filtered by a condition
func (s Shrink) Filter(condition func(interface{}) bool) Shrink {
	if condition == nil {
		return s
	}
	return func() (interface{}, bool) {
		value, ok := s()
		for ok && !condition(value) {
			value, ok = s()
		}
		return value, ok
	}
}

// Map creates a shrink by a appling a converter to each element of a shrink
func (s Shrink) Map(f func(interface{}) interface{}) Shrink {
	return func() (interface{}, bool) {
		value, ok := s()
		if ok {
			return f(value), ok
		}
		return nil, false
	}
}

type concatedShrink struct {
	index   int
	shrinks []Shrink
}

func (c *concatedShrink) Next() (interface{}, bool) {
	for c.index < len(c.shrinks) {
		value, ok := c.shrinks[c.index]()
		if ok {
			return value, ok
		}
		c.index++
	}
	return nil, false
}

// ConcatShinks concats an array of shrinks to a single shrinks
func ConcatShrinks(shrinks []Shrink) Shrink {
	concated := &concatedShrink{
		index:   0,
		shrinks: shrinks,
	}
	return concated.Next
}

type interleaved struct {
	first          Shrink
	second         Shrink
	firstExhausted bool
	secondExhaused bool
	state          bool
}

func (i *interleaved) Next() (interface{}, bool) {
	for !i.firstExhausted && !i.secondExhaused {
		i.state = !i.state
		if i.state && !i.firstExhausted {
			value, ok := i.first()
			if ok {
				return value, true
			}
			i.firstExhausted = true
		} else if !i.state && !i.secondExhaused {
			value, ok := i.second()
			if ok {
				return value, true
			}
			i.secondExhaused = true
		}
	}
	return nil, false
}

// Interleave this shrink with another
// Both shrinks are expected to produce the same result
func (s Shrink) Interleave(other Shrink) Shrink {
	interleaved := &interleaved{
		first:  s,
		second: other,
	}
	return interleaved.Next
}

// Shrinker creates a shrink for a given value
type Shrinker func(value interface{}) Shrink

var NoShrink = Shrink(func() (interface{}, bool) {
	return nil, false
})

var NoShrinker = Shrinker(func(value interface{}) Shrink {
	return NoShrink
})
