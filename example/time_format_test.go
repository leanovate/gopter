package example

import (
	"testing"
	"time"

	"github.com/untoldwind/gopter"
	"github.com/untoldwind/gopter/gen"
	"github.com/untoldwind/gopter/prop"
)

func TestTimeFormat(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("time can be parsed",
		prop.ForAllNoShrink(prop.Check1(func(arg interface{}) (interface{}, error) {
			actual := arg.(time.Time)
			str := actual.Format(time.RFC3339Nano)
			parsed, err := time.Parse(time.RFC3339Nano, str)
			return actual.Equal(parsed), err
		}), gen.TimeRange(time.Now(), time.Duration(100*24*365)*time.Hour)))

	properties.Run(t)
}
