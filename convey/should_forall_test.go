package convey_test

import (
	"errors"
	"math"
	"reflect"
	"testing"

	"github.com/leanovate/gopter"
	. "github.com/leanovate/gopter/convey"
	"github.com/leanovate/gopter/gen"
	. "github.com/smartystreets/goconvey/convey"
)

type QudraticEquation struct {
	A, B, C float64
}

func (q *QudraticEquation) Eval(x float64) float64 {
	return q.A*x*x + q.B*x + q.C
}

func (q *QudraticEquation) Solve() (float64, float64, error) {
	if q.A == 0 {
		return 0, 0, errors.New("No solution")
	}
	v := q.B*q.B - 4*q.A*q.C
	if v < 0 {
		return 0, 0, errors.New("No solution")
	}
	v = math.Sqrt(v)
	return (-q.B + v) / 2 / q.A, (-q.B - v) / 2 / q.A, nil
}

func TestShouldSucceedForAll(t *testing.T) {
	Convey("Quadratic equations should be solvable", t, func() {
		checkSolve := func(quadratic *QudraticEquation) bool {
			x1, x2, err := quadratic.Solve()
			if err != nil {
				return true
			}

			return math.Abs(quadratic.Eval(x1)) < 1e-5 && math.Abs(quadratic.Eval(x2)) < 1e-5
		}
		anyQudraticEquation := gen.StructPtr(reflect.TypeOf(QudraticEquation{}), map[string]gopter.Gen{
			"A": gen.Float64Range(-1e8, 1e8),
			"B": gen.Float64Range(-1e8, 1e8),
			"C": gen.Float64Range(-1e8, 1e8),
		})

		So(checkSolve, ShouldSucceedForAll, anyQudraticEquation)
	})
}
