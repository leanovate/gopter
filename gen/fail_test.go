package gen_test

import (
	"testing"

	"github.com/untoldwind/gopter/gen"
)

func TestFail(t *testing.T) {
	fail := gen.Fail

	value, ok := fail.Sample()

	if value != nil || ok {
		t.Fail()
	}
}
