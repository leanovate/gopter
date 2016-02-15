package example

import (
	"testing"
	"time"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestTimeFormat(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("time can be parsed",
		prop.ForAllNoShrink1(
			gen.TimeRange(time.Now(), time.Duration(100*24*365)*time.Hour),
			func(arg interface{}) (interface{}, error) {
				actual := arg.(time.Time)
				str := actual.Format(time.RFC3339Nano)
				parsed, err := time.Parse(time.RFC3339Nano, str)
				return actual.Equal(parsed), err
			},
		))

	properties.Run(t)
}
