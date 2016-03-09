package gen

import (
	"reflect"
	"regexp"
	"regexp/syntax"
	"strings"

	"github.com/leanovate/gopter"
)

func RegexMatch(regexStr string) gopter.Gen {
	regexSyntax, err := syntax.Parse(regexStr, syntax.Perl)
	if err != nil {
		return Fail(reflect.TypeOf(""))
	}
	regex, err := regexp.Compile(regexStr)
	if err != nil {
		return Fail(reflect.TypeOf(""))
	}
	return regexMatchGen(regexSyntax.Simplify()).SuchThat(func(v interface{}) bool {
		return regex.MatchString(v.(string))
	}).WithShrinker(SliceShrinker(gopter.NoShrinker))
}

func regexMatchGen(regex *syntax.Regexp) gopter.Gen {
	switch regex.Op {
	case syntax.OpLiteral:
		return Const(string(regex.Rune))
	case syntax.OpCharClass:
		gens := make([]gopter.Gen, 0, len(regex.Rune)/2)
		for i := 0; i+1 < len(regex.Rune); i += 2 {
			gens = append(gens, RuneRange(regex.Rune[i], regex.Rune[i+1]).Map(func(v interface{}) interface{} {
				return string(v.(rune))
			}))
		}
		return OneGenOf(gens...)
	case syntax.OpAnyChar:
		return Rune().Map(func(v interface{}) interface{} {
			return string(v.(rune))
		})
	case syntax.OpAnyCharNotNL:
		return RuneNoControl().Map(func(v interface{}) interface{} {
			return string(v.(rune))
		})
	case syntax.OpCapture:
		return regexMatchGen(regex.Sub[0])
	case syntax.OpStar:
		elementGen := regexMatchGen(regex.Sub[0])
		return SliceOf(elementGen).Map(func(v interface{}) interface{} {
			return strings.Join(v.([]string), "")
		})
	case syntax.OpPlus:
		elementGen := regexMatchGen(regex.Sub[0])
		return gopter.CombineGens(elementGen, SliceOf(elementGen)).Map(func(v interface{}) interface{} {
			vs := v.([]interface{})
			return vs[0].(string) + strings.Join(vs[1].([]string), "")
		})
	case syntax.OpQuest:
		elementGen := regexMatchGen(regex.Sub[0])
		return OneGenOf(Const(""), elementGen)
	case syntax.OpConcat:
		gens := make([]gopter.Gen, len(regex.Sub))
		for i, sub := range regex.Sub {
			gens[i] = regexMatchGen(sub)
		}
		return gopter.CombineGens(gens...).Map(func(v interface{}) interface{} {
			result := ""
			for _, str := range v.([]interface{}) {
				result += str.(string)
			}
			return result
		})
	case syntax.OpAlternate:
		gens := make([]gopter.Gen, len(regex.Sub))
		for i, sub := range regex.Sub {
			gens[i] = regexMatchGen(sub)
		}
		return OneGenOf(gens...)
	}
	return Const("")
}
