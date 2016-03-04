package gopter_test

import (
	"reflect"
	"testing"

	"github.com/leanovate/gopter"
)

type counterShrink struct {
	n int
}

func (c *counterShrink) Next() (interface{}, bool) {
	if c.n > 0 {
		v := c.n
		c.n--
		return v, true
	}
	return 0, false
}

func TestShinkAll(t *testing.T) {
	counter := &counterShrink{n: 10}
	shrink := gopter.Shrink(counter.Next)

	all := shrink.All()
	if !reflect.DeepEqual(all, []interface{}{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}) {
		t.Errorf("Invalid all: %#v", all)
	}
}

func TestShrinkFilter(t *testing.T) {
	counter := &counterShrink{n: 20}
	shrink := gopter.Shrink(counter.Next)

	all := shrink.Filter(func(v interface{}) bool {
		return v.(int)%2 == 0
	}).All()
	if !reflect.DeepEqual(all, []interface{}{20, 18, 16, 14, 12, 10, 8, 6, 4, 2}) {
		t.Errorf("Invalid all: %#v", all)
	}

	counter = &counterShrink{n: 5}
	shrink = gopter.Shrink(counter.Next)

	all = shrink.Filter(nil).All()
	if !reflect.DeepEqual(all, []interface{}{5, 4, 3, 2, 1}) {
		t.Errorf("Invalid all: %#v", all)
	}
}

func TestShrinkConcat(t *testing.T) {
	counterShrink1 := &counterShrink{n: 5}
	counterShrink2 := &counterShrink{n: 4}
	shrink1 := gopter.Shrink(counterShrink1.Next)
	shrink2 := gopter.Shrink(counterShrink2.Next)

	all := gopter.ConcatShrinks(shrink1, shrink2).All()
	if !reflect.DeepEqual(all, []interface{}{5, 4, 3, 2, 1, 4, 3, 2, 1}) {
		t.Errorf("Invalid all: %#v", all)
	}
}

func TestShrinkInterleave(t *testing.T) {
	counterShrink1 := &counterShrink{n: 5}
	counterShrink2 := &counterShrink{n: 7}

	shrink1 := gopter.Shrink(counterShrink1.Next)
	shrink2 := gopter.Shrink(counterShrink2.Next)

	all := shrink1.Interleave(shrink2).All()
	if !reflect.DeepEqual(all, []interface{}{5, 7, 4, 6, 3, 5, 2, 4, 1, 3, 2, 1}) {
		t.Errorf("Invalid all: %#v", all)
	}
}
