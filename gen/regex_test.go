package gen_test

import (
	"regexp"
	"testing"

	"github.com/leanovate/gopter/gen"
)

func TestRegexMatch(t *testing.T) {
	regexs := []string{
		"[a-z][0-9a-zA-Z]*",
		"AB[0-9]+",
		"1?(zero|one)0",
		"ABCD.+1234",
		"^[0-9]{3}[A-Z]{5,}[a-z]{10,20}$",
		"(?s)[^0-9]*ABCD.*1234",
	}
	for _, regex := range regexs {
		pattern, err := regexp.Compile(regex)
		if err != nil {
			t.Error("Invalid regex", err)
		}
		gen := gen.RegexMatch(regex)
		for i := 0; i < 100; i++ {
			value, ok := gen.Sample()

			if !ok || value == nil {
				t.Errorf("Invalid value: %#v", value)
			}
			str, ok := value.(string)
			if !ok || !pattern.MatchString(str) {
				t.Errorf("Invalid value: %#v", value)
			}
		}
	}

	gen := gen.RegexMatch("]]}})Invalid{]]]")
	value, ok := gen.Sample()
	if ok || value != nil {
		t.Errorf("Invalid value: %#v", value)
	}
}
