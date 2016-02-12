package prop

import "github.com/untoldwind/gopter"

func noShrinkArgs(samples []interface{}) []gopter.PropArg {
	result := make([]gopter.PropArg, len(samples))
	for i, sample := range samples {
		result[i] = gopter.PropArg{
			Arg:     sample,
			OrigArg: sample,
			Shrinks: 0,
		}
	}
	return result
}

func ForAllNoShrink(sampleCheck SampleCheck, gens ...gopter.Gen) gopter.Prop {
	return func(genParams *gopter.GenParameters) gopter.PropResult {
		samples := make([]interface{}, len(gens))
		var ok bool
		for i, gen := range gens {
			samples[i], ok = gen.DoApply(genParams).Retrieve()
			if !ok {
				return gopter.PropResult{
					Status: gopter.PropUndecided,
				}
			}
		}
		return convertResult(sampleCheck(samples...))
	}
}
